package asana

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"moul.io/hcfilters"
)

const (
	defaultUserAgent = "asanaman/1.0"
	defaultBaseURL   = "https://app.asana.com/api/1.0/"
)

type Client struct {
	// args
	token     string
	domain    string
	cachePath string
	logger    *zap.Logger

	// internal
	http      *http.Client
	baseURL   *url.URL
	userAgent string
}

func New(token, domain, cachePath string, logger *zap.Logger) (*Client, error) {
	client := Client{
		token:     token,
		domain:    domain,
		cachePath: cachePath,
		userAgent: defaultUserAgent,
		logger:    logger,
	}
	client.baseURL, _ = url.Parse(defaultBaseURL)

	client.http = &http.Client{
		Transport: httpcache.NewTransport(
			hcfilters.MaxSize( // skip caching results > 10Mb
				diskcache.New(cachePath),
				10*1024*1024,
			),
		),
	}

	return &client, nil
}

type ReqOpts struct {
	Method string
	Path   string
	Data   interface{}
	Form   url.Values
	// opt *Filter
}

type request struct {
	Data interface{} `json:"data,omitempty"`
}

type response struct {
	Data   interface{}  `json:"data,omitempty"`
	Errors []asanaError `json:"errors,omitempty"`
}

type asanaError struct {
	Phrase  string `json:"phrase,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e asanaError) Error() string {
	return fmt.Sprintf("%v - %v", e.Message, e.Phrase)
}

func (c *Client) Request(ctx context.Context, opts ReqOpts, ret interface{}) error {
	logger := c.logger
	defer func() { logger.Debug("req") }()
	// compute path
	rel, err := url.Parse(opts.Path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}
	u := c.baseURL.ResolveReference(rel)
	logger = logger.With(zap.String("u", u.String()))

	var body io.Reader

	// init request
	req, err := http.NewRequest(opts.Method, u.String(), body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Add("Authorization", "Bearer "+c.token)

	// compute input body
	switch {
	case opts.Data != nil:
		req.Header.Set("Content-Type", "application/json")

		b, err := json.Marshal(request{Data: opts.Data})
		if err != nil {
			return fmt.Errorf("encode body: %w", err)
		}
		c.logger.Debug("")
		body = bytes.NewReader(b)
	case opts.Form != nil:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		body = strings.NewReader(opts.Form.Encode())
	}

	// perform request
	resp, err := c.http.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("HTTP request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	// compute response
	res := &response{Data: ret}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return fmt.Errorf("cannot parse response: %w", err)
	}

	// error handling
	var errs error
	for _, err := range res.Errors {
		errs = multierr.Append(errs, err)
	}
	return err
}
