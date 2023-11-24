package openseaapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xTransact/errx/v2"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseaconsts"
	"github.com/xTransact/openseaapi/openseamodels"
)

// GetListings gets the complete set of active, valid listings.
// @Param ch (chain.Chain): The blockchain on which to filter the results.
// DOC: https://docs.opensea.io/reference/get_listings
func (c *client) GetListings(ctx context.Context, ch chain.Chain, payload *openseamodels.OrderPayload) (
	resp *openseamodels.OrdersResponse, err error) {

	// POST /api/v2/orders/{chain}/{protocol}/listings
	url := fmt.Sprintf("%s/api/v2/orders/%s/%s/listings",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), openseaconsts.ProtocolName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	if payload != nil {
		if err = payload.Validate(); err != nil {
			return nil, errx.Wrap(err, "invalid payload")
		}

		qs := payload.ToQuery()
		if len(qs) > 0 {
			req.URL.RawQuery = qs.Encode()
		}
	}

	c.acceptJson(req)
	body, err := c.doRequest(req, ch.IsTestNet())
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.OrdersResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
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
		return nil, errx.WithStack(err)
	}

	if err = payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	qs := payload.ToQuery()
	if len(qs) > 0 {
		req.URL.RawQuery = qs.Encode()
	}

	c.acceptJson(req)
	body, err := c.doRequest(req, o.testnets)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.ListingsByCollectionResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}

// CreateListing lists a single NFT (ERC721 or ERC1155) for sale on the OpenSea marketplace.
// DOC: https://docs.opensea.io/reference/post_listing
func (c *client) CreateListing(ctx context.Context, ch chain.Chain,
	payload *openseamodels.CreateOrderPayload) (resp *openseamodels.CreateListingResponse, err error) {

	if err = payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	payloadData, err := json.Marshal(payload)
	if err != nil {
		return nil, errx.Wrap(err, "marshal payload")
	}

	// POST /api/v2/orders/{chain}/{protocol}/listings
	url := fmt.Sprintf("%s/api/v2/orders/%s/%s/listings", openseaapiutils.GetBaseURLByChain(ch), ch.Value(), openseaconsts.ProtocolName)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payloadData))
	if err != nil {
		return nil, errx.WithStack(err)
	}

	c.acceptJson(req)
	c.contentTypeJson(req)
	body, err := c.doRequest(req, ch.IsTestNet())
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.CreateListingResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}
