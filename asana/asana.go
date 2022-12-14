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
	"time"

	"github.com/google/go-querystring/query"
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
				10*1024*1024, // nolint:gomnd
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
	Opts   interface{}
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
	start := time.Now()
	defer func() { logger.Debug("req", zap.Duration("elapsed", time.Since(start))) }()
	// compute path
	u, err := url.Parse(opts.Path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}
	if opts.Opts != nil {
		qs, err := query.Values(opts.Opts)
		if err != nil {
			return fmt.Errorf("invaid filter opts: %w", err)
		}
		u.RawQuery = qs.Encode()
	}
	rel, err := url.Parse(u.String())
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}
	u = c.baseURL.ResolveReference(rel)
	logger = logger.With(zap.String("url", u.String()))

	// compute input body
	var body io.Reader
	switch {
	case opts.Data != nil:
		b, err := json.Marshal(request{Data: opts.Data})
		if err != nil {
			return fmt.Errorf("encode body: %w", err)
		}
		logger = logger.With(zap.Int("input-data-len", len(b)))
		body = bytes.NewReader(b)
	case opts.Form != nil:
		encoded := opts.Form.Encode()
		body = strings.NewReader(encoded)
		logger = logger.With(zap.Int("input-form-len", len(encoded)))
	}

	// init request
	if opts.Method == "" {
		opts.Method = "GET"
	}
	logger = logger.With(zap.String("method", opts.Method))
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
	case opts.Form != nil:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
