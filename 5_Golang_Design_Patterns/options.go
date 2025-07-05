package patterns

import "time"

type Client struct {
	timeout time.Duration
}

type Option func(*Client)

func WithTimeout(time time.Duration) Option {
	return func(c *Client) {
		c.timeout = time
	}
}

func NewClient(options ...Option) *Client {
	client := &Client{}

	for _, opt := range options {
		opt(client)
	}
	return client
}

func main() {
	c := NewClient(WithTimeout(5 * time.Second))
}
