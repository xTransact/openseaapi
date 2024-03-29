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

// FulfillListing retrieves all the information, including signatures, needed to fulfill a listing directly onchain.
// @Param orderHash: required: hash of the order to fulfill.
// @Param ch (chain.Chain): required
// @Param: fulfiller: required: fulfiller address
func (c *client) FulfillListing(ctx context.Context, ch chain.Chain,
	orderHash, fulfiller string) (resp *openseamodels.FulfillmentDataResponse, err error) {

	if orderHash == "" || fulfiller == "" || ch < 0 {
		return nil, errx.New("illegal arguments: nil")
	}

	payload := &openseamodels.FulfillListingPayload{
		Listing: &openseamodels.FulfillOrder{
			Hash:            orderHash,
			Chain:           ch.Value(),
			ProtocolAddress: openseaconsts.SeaportV16Address.String(),
		},
		FulFiller: &openseamodels.Fulfiller{
			Address: fulfiller,
		},
	}

	payloadData, err := json.Marshal(payload)
	if err != nil {
		return nil, errx.Wrap(err, "marshal payload")
	}

	// POST /api/v2/listings/fulfillment_data
	url := fmt.Sprintf("%s/api/v2/listings/fulfillment_data", openseaapiutils.GetBaseURLByChain(ch))

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

	resp = new(openseamodels.FulfillmentDataResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}

// FulfillOffer retrieves all the information, including signatures, needed to fulfill an offer directly onchain.
// DOC: https://docs.opensea.io/reference/generate_offer_fulfillment_data_v2
func (c *client) FulfillOffer(ctx context.Context, payload *openseamodels.FulfillOfferPayload,
	opts ...RequestOptionFn) (resp *openseamodels.FulfillmentDataResponse, err error) {

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

	// POST /api/v2/offers/fulfillment_data
	url := fmt.Sprintf("%s/api/v2/offers/fulfillment_data", openseaapiutils.GetBaseURL(o.testnets))

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

	resp = new(openseamodels.FulfillmentDataResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}
