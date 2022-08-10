package asana

type Client struct {
	token  string
	domain string
}

func New(token, domain string) (*Client, error) {
	client := Client{
		token:  token,
		domain: domain,
	}
	return &client, nil
}
