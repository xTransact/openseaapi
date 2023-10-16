package openseaapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseaconsts"
	"github.com/xTransact/openseaapi/openseamodels"
)

// GetListings gets the complete set of active, valid listings.
// @Param ch (chain.Chain): The blockchain on which to filter the results.
// DOC: https://docs.opensea.io/reference/get_listings
func (c *client) GetListings(ctx context.Context, ch chain.Chain, payload *openseamodels.GetListingsPayload) (
	resp *openseamodels.GetListingsResponse, err error) {

	// POST /api/v2/orders/{chain}/{protocol}/listings
	url := fmt.Sprintf("%s/api/v2/orders/%s/%s/listings", openseaapiutils.GetBaseURLByChain(ch), ch.Value(), openseaconsts.ProtocolName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	if payload != nil {
		if err = payload.Validate(); err != nil {
			return nil, fmt.Errorf("invalid payload: %w", err)
		}

		qs := payload.ToQuery()
		if len(qs) > 0 {
			req.URL.RawQuery = qs.Encode()
		}
	}
	c.acceptJson(req)
	if !ch.IsTestNet() {
		c.challenge(req)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp = new(openseamodels.GetListingsResponse)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}

// GetAllListingsByCollection gets all active, valid listings for a single collection.
// @Param collection_slug: required: Unique string to identify a collection on OpenSea. This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
// @Param limit: The number of listings to return. Must be between 1 and 100. Default: 100
// @Param next: The cursor for the next page of results. This is returned from a previous request.
// DOC: https://docs.opensea.io/reference/get_all_listings_on_collection_v2
func (c *client) GetAllListingsByCollection(ctx context.Context, payload *openseamodels.GetAllListingsByCollectionPayload,
	opts ...RequestOptionFn) (resp *openseamodels.ListingsByCollectionResponse, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	// GET /api/v2/listings/collection/{collection_slug}/all
	url := fmt.Sprintf("%s/api/v2/listings/collection/%s/all", openseaapiutils.GetBaseURL(o.testnets), payload.CollectionSlug)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	if err = payload.Validate(); err != nil {
		return nil, fmt.Errorf("invalid payload: %w", err)
	}

	qs := payload.ToQuery()
	if len(qs) > 0 {
		req.URL.RawQuery = qs.Encode()
	}

	c.acceptJson(req)
	if !o.testnets {
		c.challenge(req)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp = new(openseamodels.ListingsByCollectionResponse)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}

// CreateListing lists a single NFT (ERC721 or ERC1155) for sale on the OpenSea marketplace.
// DOC: https://docs.opensea.io/reference/post_listing
func (c *client) CreateListing(ctx context.Context, ch chain.Chain,
	payload *openseamodels.CreateListingPayload) (resp *openseamodels.CreateListingResponse, err error) {

	if err = payload.Validate(); err != nil {
		return nil, fmt.Errorf("invalid payload: %w", err)
	}

	payloadData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// POST /api/v2/orders/{chain}/{protocol}/listings
	url := fmt.Sprintf("%s/api/v2/orders/%s/%s/listings", openseaapiutils.GetBaseURLByChain(ch), ch.Value(), openseaconsts.ProtocolName)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payloadData))
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	c.acceptJson(req)
	c.contentTypeJson(req)
	if !ch.IsTestNet() {
		c.challenge(req)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp = new(openseamodels.CreateListingResponse)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}

// FulfillListing retrieves all the information, including signatures, needed to fulfill a listing directly onchain.
// @Param orderHash: required: hash of the order to fulfill.
// @Param ch (chain.Chain): required
// @Param: fulfiller: required: fulfiller address
func (c *client) FulfillListing(ctx context.Context, ch chain.Chain,
	orderHash, fulfiller string) (resp *openseamodels.ListingsFulfillmentData, err error) {

	if orderHash == "" || fulfiller == "" || ch < 0 {
		return nil, errors.New("illegal arguments: nil")
	}

	payload := &openseamodels.FulfillListingPayload{
		Listing: &openseamodels.FulfillListingPayloadListing{
			Hash:            orderHash,
			Chain:           ch.Value(),
			ProtocolAddress: openseaconsts.SeaportV15Address.String(),
		},
		FulFiller: &openseamodels.FulfillListingPayloadFulfiller{
			Address: fulfiller,
		},
	}

	payloadData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// POST /api/v2/listings/fulfillment_data
	url := fmt.Sprintf("%s/api/v2/listings/fulfillment_data", openseaapiutils.GetBaseURLByChain(ch))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payloadData))
	if err != nil {
		return nil, fmt.Errorf("failed to new request: %w", err)
	}

	c.acceptJson(req)
	c.contentTypeJson(req)
	if !ch.IsTestNet() {
		c.challenge(req)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp = new(openseamodels.ListingsFulfillmentData)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return resp, nil
}
