package openseaapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/http2"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

type Servicer interface {
	GetListings(ctx context.Context, ch chain.Chain, payload *openseamodels.GetListingsPayload) (resp *openseamodels.GetListingsResponse, err error)
}

type client struct {
	config     *options
	httpClient *http.Client
}

func NewClient(opts ...OptionFn) Servicer {
	o := new(options)
	for _, apply := range opts {
		apply(o)
	}

	timeout := o.timeout
	if timeout == 0 {
		timeout = time.Second * 30
	}

	httpClient := &http.Client{
		Timeout: timeout,
		Transport: &http2.Transport{
			AllowHTTP: true,
		},
	}

	return &client{
		config:     o,
		httpClient: httpClient,
	}
}

func (c *client) challenge(r *http.Request) {
	if c.config.apiKey != "" {
		r.Header.Set("x-api-key", c.config.apiKey)
	}
}

func (c *client) acceptJson(r *http.Request) {
	r.Header.Set("Accept", "application/json")
}

func (c *client) contentTypeJson(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
}

func (c *client) doRequest(r *http.Request) ([]byte, error) {
	res, err := c.httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w, status: %s", err, res.Status)
	}

	if res.StatusCode != http.StatusOK {
		return nil, openseaapiutils.ParseFailureResponse(res, body)
	}

	return body, nil
}
