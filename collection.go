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

// ListCollections gets a list of OpenSea collections.
// @param chain_identifier: The blockchain on which to filter the results. e.g., ethereum
// @param include_hidden: If true, will return hidden collections. Default: false
// @param limit: The number of collections to return. Must be between 1 and 100. Default: 100
// @param next: The cursor for the next page of results. This is returned from a previous request.
// DOC: https://docs.opensea.io/reference/list_collections
func (c *client) ListCollections(ctx context.Context, payload *openseamodels.ListCollectionsPayload) (
	resp *openseamodels.CollectionsResponse, err error) {

	if err = payload.Validate(); err != nil {
		return nil, err
	}

	ch := chain.RequireFromString(payload.ChainIdentifier)

	// GET /api/v2/collections
	url := fmt.Sprintf("%s/api/v2/collections", openseaapiutils.GetBaseURLByChain(ch))

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

	resp = new(openseamodels.CollectionsResponse)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}

// GetCollection gets a single collection including details such as fees, traits, and links.
// @param collection_slug: required: Unique string to identify a collection on OpenSea.
// This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
// DOC: https://docs.opensea.io/reference/get_collection
func (c *client) GetCollection(ctx context.Context, collectionSlug string, opts ...RequestOptionFn) (
	resp *openseamodels.SingleCollection, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	// GET /api/v2/collections/{collection_slug}
	url := fmt.Sprintf("%s/api/v2/collections/%s", openseaapiutils.GetBaseURL(o.testnets), collectionSlug)

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

	resp = new(openseamodels.SingleCollection)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}

// GetCollectionStats gets stats for a single collection.
// DOC: https://docs.opensea.io/reference/get_collection_stats
func (c *client) GetCollectionStats(ctx context.Context, collectionSlug string, opts ...RequestOptionFn) (
	resp *openseamodels.CollectionStats, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	// GET /api/v2/collections/{collection_slug}/stats
	url := fmt.Sprintf("%s/api/v2/collections/%s/stats", openseaapiutils.GetBaseURL(o.testnets), collectionSlug)

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

	resp = new(openseamodels.CollectionStats)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}
