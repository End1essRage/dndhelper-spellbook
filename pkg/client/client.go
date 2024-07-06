package client

import (
	"encoding/json"
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
		basePath: "/api/spells",
		client:   http.Client{},
	}
}

func (c *Client) GetSpellInfo(spellName string) (Spell, error) {
	logrus.Info("client - GetSpellInfo")
	resp, err := c.sendRequest(spellName)

	var spell Spell
	err = json.Unmarshal(resp, &spell)
	logrus.Info("spell damage is: ", spell.Damage)

	return spell, err
}

func (c *Client) sendRequest(spellname string) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, spellname),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}

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
