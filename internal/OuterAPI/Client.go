package OuterAPI

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

/*
https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search
*/
const (
	host   = "p2p.binance.com"
	method = "bapi/c2c/v2/friendly/c2c/adv/search"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "" + token
}

func (c *Client) PostRequest(postBody io.Reader, method string) []byte {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), postBody)
	if err != nil {
		log.Println("request error", err)
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		log.Println("Response error", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("read body error", err)
	}
	return body
}
