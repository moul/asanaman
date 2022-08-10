package asana

type Client struct{}

func New() (*Client, error) {
	return &Client{}, nil
}
