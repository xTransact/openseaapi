package openseaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

// GetOrder gets a single order, offer or listing, by its order hash.
// Protocol and Chain are required to prevent hash collisions.
// DOC: https://docs.opensea.io/reference/get_order
func (c *client) GetOrder(ctx context.Context, payload *openseamodels.GetOrderPayload) (
	resp *openseamodels.GetOrderResponse, err error) {

	if err = payload.Validate(); err != nil {
		return nil, err
	}

	ch := chain.RequireFromString(payload.Chain)

	// GET /api/v2/orders/chain/{chain}/protocol/{protocol_address}/{order_hash}

	url := fmt.Sprintf("%s/api/v2/orders/chain/%s/protocol/%s/%s",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), payload.ProtocolAddress, payload.OrderHash)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	c.acceptJson(req)
	if !ch.IsTestNet() {
		c.challenge(req)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp = new(openseamodels.GetOrderResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}