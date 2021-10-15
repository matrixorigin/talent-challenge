package client

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/matrixorigin/talent-challenge/matrixbase/distributed/pkg/model"
	"github.com/valyala/fasthttp"
)

// Client client
type Client interface {
	// Set set key-value to store
	Set(key []byte, value []byte) error
	// Get returns the value from store
	Get(key []byte) ([]byte, error)
	// Delete remove the key from store
	Delete(key []byte) error
}

type httpClient struct {
	url    string
	addrs  []string
	client *fasthttp.HostClient
}

// New create a client
func New(addrs ...string) Client {
	return &httpClient{
		url:   "http://" + addrs[0] + "/kv",
		addrs: addrs,
		client: &fasthttp.HostClient{
			MaxConns:            2048,
			MaxIdleConnDuration: 60 * time.Second,
			MaxConnWaitTimeout:  time.Second,
			Addr:                strings.Join(addrs, ","),
		},
	}
}

func (c *httpClient) Set(key, value []byte) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(c.url)
	req.Header.SetMethod("POST")

	r := &model.Request{Key: string(key), Value: string(value)}
	req.SetBody(r.MustMarshal())

	_, err := c.doHTTP(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *httpClient) Get(key []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(fmt.Sprintf("%s?key=%s", c.url, string(key)))
	req.Header.SetMethod("GET")

	res, err := c.doHTTP(req)
	if err != nil {
		return nil, err
	}

	if res.Data == nil {
		return nil, nil
	}

	v, err := base64.StdEncoding.DecodeString(res.Data.(string))
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (c *httpClient) Delete(key []byte) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(fmt.Sprintf("%s?key=%s", c.url, string(key)))
	req.Header.SetMethod("DELETE")

	_, err := c.doHTTP(req)
	if err != nil {
		return nil
	}
	return nil
}

func (c *httpClient) doHTTP(req *fasthttp.Request) (*model.JSONResult, error) {
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := c.client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("error code %d", resp.StatusCode())
	}

	result := model.MustUnmarshalResult(resp.Body())
	if result.Code != 0 {
		return nil, errors.New(result.Err)
	}

	return result, nil
}
