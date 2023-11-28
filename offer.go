package openseaapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xTransact/errx/v3"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseaconsts"
	"github.com/xTransact/openseaapi/openseamodels"
)

// BuildOffer builds a portion of a criteria offer including the merkle tree needed to post an offer.
// DOC: https://docs.opensea.io/reference/build_offer_v2
func (c *client) BuildOffer(ctx context.Context, payload *openseamodels.BuildOfferPayload,
	opts ...RequestOptionFn) (resp *openseamodels.BuildOfferResponse, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	if err = payload.Validate(); err != nil {
		return nil, err
	}

	payloadData, err := json.Marshal(payload)
	if err != nil {
		return nil, errx.Wrap(err, "marshal payload")
	}

	// POST /api/v2/offers/build
	url := fmt.Sprintf("%s/api/v2/offers/build", openseaapiutils.GetBaseURL(o.testnets))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payloadData))
	if err != nil {
		return nil, errx.WithStack(err)
	}

	c.acceptJson(req)
	c.contentTypeJson(req)
	body, err := c.doRequest(req, o.testnets)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.BuildOfferResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}

// CreateCriteriaOffer creates a criteria offer to purchase any NFT in a collection or which matches the specified trait.
// DOC: https://docs.opensea.io/reference/post_criteria_offer_v2
func (c *client) CreateCriteriaOffer(ctx context.Context, payload *openseamodels.CreateCriteriaOfferPayload,
	opts ...RequestOptionFn) (resp *openseamodels.OfferResponse, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	if err = payload.Validate(); err != nil {
		return nil, err
	}

	payloadData, err := json.Marshal(payload)
	if err != nil {
		return nil, errx.Wrap(err, "marshal payload")
	}

	// POST /api/v2/offers
	url := fmt.Sprintf("%s/api/v2/offers", openseaapiutils.GetBaseURL(o.testnets))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payloadData))
	if err != nil {
		return nil, errx.WithStack(err)
	}

	c.acceptJson(req)
	c.contentTypeJson(req)
	body, err := c.doRequest(req, o.testnets)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.OfferResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}

// CreateIndividualOffer creates an offer to purchase a single NFT (ERC721 or ERC1155).
// DOC: https://docs.opensea.io/reference/post_offer
func (c *client) CreateIndividualOffer(ctx context.Context, ch chain.Chain,
	payload *openseamodels.CreateOrderPayload) (resp *openseamodels.OrderResponse, err error) {

	if err = payload.Validate(); err != nil {
		return nil, err
	}

	payloadData, err := json.Marshal(payload)
	if err != nil {
		return nil, errx.Wrap(err, "marshal payload")
	}

	// POST /api/v2/orders/{chain}/{protocol}/offers
	url := fmt.Sprintf("%s/api/v2/orders/%s/%s/offers",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), openseaconsts.ProtocolName)

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

	resp = new(openseamodels.OrderResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}

// GetCollectionOffers gets the active, valid collection offers for the specified collection.
// DOC: https://docs.opensea.io/reference/get_collection_offers_v2
func (c *client) GetCollectionOffers(ctx context.Context, collectionSlug string,
	opts ...RequestOptionFn) (resp *openseamodels.Offers, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	// GET /api/v2/offers/collection/{collection_slug}
	url := fmt.Sprintf("%s/api/v2/offers/collection/%s",
		openseaapiutils.GetBaseURL(o.testnets), collectionSlug)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	c.acceptJson(req)
	body, err := c.doRequest(req, o.testnets)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.Offers)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}

// GetAllCollectionOffers gets all active, valid offers for the specified collection. This includes individual and criteria offers.
// DOC: https://docs.opensea.io/reference/get_all_offers_on_collection_v2
func (c *client) GetAllCollectionOffers(ctx context.Context, payload *openseamodels.CollectionPayload,
	opts ...RequestOptionFn) (resp *openseamodels.PageableOffers, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	if err = payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	// GET /api/v2/offers/collection/{collection_slug}/all
	url := fmt.Sprintf("%s/api/v2/offers/collection/%s/all",
		openseaapiutils.GetBaseURL(o.testnets), payload.CollectionSlug)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
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

	resp = new(openseamodels.PageableOffers)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}

// GetIndividualOffers gets the active, valid individual offers. This does not include criteria offers.
// DOC: https://docs.opensea.io/reference/get_offers
func (c *client) GetIndividualOffers(ctx context.Context, ch chain.Chain,
	payload *openseamodels.OrderPayload) (resp *openseamodels.OrdersResponse, err error) {

	if err = payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	// GET /api/v2/orders/{chain}/{protocol}/offers
	url := fmt.Sprintf("%s/api/v2/orders/%s/%s/offers",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), openseaconsts.ProtocolName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	qs := payload.ToQuery()
	if len(qs) > 0 {
		req.URL.RawQuery = qs.Encode()
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

// GetTraitOffers gets the active, valid trait offers for the specified collection.
// DOC: https://docs.opensea.io/reference/get_trait_offers_v2
func (c *client) GetTraitOffers(ctx context.Context, payload *openseamodels.GetTraitOffersPayload,
	opts ...RequestOptionFn) (resp *openseamodels.Offers, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	if err = payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	// GET /api/v2/offers/collection/{collection_slug}/traits
	url := fmt.Sprintf("%s/api/v2/offers/collection/%s/traits",
		openseaapiutils.GetBaseURL(o.testnets), payload.CollectionSlug)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
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

	resp = new(openseamodels.Offers)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}
