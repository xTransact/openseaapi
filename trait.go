package openseaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

// GetTraits gets the traits in a collection.
// @param: collection_slug: required: Unique string to identify a collection on OpenSea.
// This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
// DOC: https://docs.opensea.io/reference/get_traits
func (c *client) GetTraits(ctx context.Context, collectionSlug string, opts ...RequestOptionFn) (
	resp *openseamodels.Trait, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	// GET /api/v2/traits/{collection_slug}
	url := fmt.Sprintf("%s/api/v2/traits/%s", openseaapiutils.GetBaseURL(o.testnets), collectionSlug)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	c.acceptJson(req)
	if !o.testnets {
		c.challenge(req)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp = new(openseamodels.Trait)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}
