package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/sirupsen/logrus"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string) *Client {
	return &Client{
		host:     host,
		basePath: "",
		client:   http.Client{},
	}
}

func (c *Client) GetSpellInfo(spellName string) (string, error) {
	logrus.Info("client - GetSpellInfo")
	return "", nil
}

func (c *Client) sendRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response: %w", err)
	}

	return body, nil
}
