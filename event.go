package openseaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xTransact/errx/v2"

	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

// ListEventsByAccount gets a list of events for an account. The list will be paginated and include up to 100 events per page.
// DOC: https://docs.opensea.io/reference/list_events_by_account
func (c *client) ListEventsByAccount(ctx context.Context, payload *openseamodels.GetEventsByAccountPayload,
	opts ...RequestOptionFn) (resp *openseamodels.AssetEventResponse, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	// GET /api/v2/events/accounts/{address}
	url := fmt.Sprintf("%s/api/v2/events/accounts/%s",
		openseaapiutils.GetBaseURL(o.testnets), payload.Address)

	return c.getEvents(ctx, payload, o.testnets, url)
}

// ListEventsByNft gets a list of events for a single NFT. The list will be paginated and include up to 100 events per page.
// DOC: https://docs.opensea.io/reference/list_events_by_nft
func (c *client) ListEventsByNft(ctx context.Context,
	payload *openseamodels.GetEventsByNftPayload) (resp *openseamodels.AssetEventResponse, err error) {

	// GET /api/v2/events/chain/{chain}/contract/{address}/nfts/{identifier}
	url := fmt.Sprintf("%s/api/v2/events/chain/%s/contract/%s/nfts/%s",
		openseaapiutils.GetBaseURLByChain(payload.Chain), payload.Chain.Value(), payload.Address, payload.Identifier)

	return c.getEvents(ctx, payload, payload.Chain.IsTestNet(), url)
}

// ListEventsByCollection gets a list of events for a collection.
// The list will be paginated and include up to 100 events per page.
// DOC: https://docs.opensea.io/reference/list_events_by_collection
func (c *client) ListEventsByCollection(ctx context.Context, payload *openseamodels.GetEventsByCollectionPayload,
	opts ...RequestOptionFn) (resp *openseamodels.AssetEventResponse, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	// GET /api/v2/events/collection/{collection_slug}
	url := fmt.Sprintf("%s/api/v2/events/collection/%s",
		openseaapiutils.GetBaseURL(o.testnets), payload.CollectionSlug)

	return c.getEvents(ctx, payload, o.testnets, url)
}

func (c *client) getEvents(ctx context.Context, payload Payloader, testnets bool, url string) (
	*openseamodels.AssetEventResponse, error) {

	if err := payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	qs := payload.ToQuery()
	if len(qs) > 0 {
		req.URL.RawQuery = qs.Encode()
	}

	c.acceptJson(req)
	body, err := c.doRequest(req, testnets)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp := new(openseamodels.AssetEventResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}
