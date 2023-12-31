package openseaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/xTransact/errx/v3"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

// ListNftsByAccount gets NFTs owned by a given account address.
// @param ch (chain.Chain): required: The blockchain on which to filter the results.
// @param address: required: The unique public blockchain identifier for the wallet.
// @param collection: Unique string to identify a collection on OpenSea.
// - This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
// @param limit: The number of NFTs to return. Must be between 1 and 200. Default: 50
// @param next: The cursor for the next page of results. This is returned from a previous request.
// DOC: https://docs.opensea.io/reference/list_nfts_by_account
func (c *client) ListNftsByAccount(ctx context.Context, ch chain.Chain,
	payload *openseamodels.GetNftsByAccountPayload) (resp *openseamodels.NftsResponse, err error) {

	if err = payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	// GET /api/v2/chain/{chain}/account/{address}/nfts
	url := fmt.Sprintf("%s/api/v2/chain/%s/account/%s/nfts",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), payload.Address.String())

	return c.getNfts(ctx, url, payload.ToQuery(), ch.IsTestNet())
}

// ListNftsByContract gets multiple NFTs for a smart contract.
// @param ch (chain.Chain): required: The blockchain on which to filter the results.
// @param address: required: The unique public blockchain identifier for the contract.
// @param limit: The number of NFTs to return. Must be between 1 and 200. Default: 50
// @param next: The cursor for the next page of results. This is returned from a previous request.
// DOC: https://docs.opensea.io/reference/list_nfts_by_contract
func (c *client) ListNftsByContract(ctx context.Context, ch chain.Chain,
	payload *openseamodels.GetNftsByContractPayload) (resp *openseamodels.NftsResponse, err error) {

	if err = payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	// GET /api/v2/chain/{chain}/contract/{address}/nfts
	url := fmt.Sprintf("%s/api/v2/chain/%s/contract/%s/nfts",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), payload.Address.String())

	return c.getNfts(ctx, url, payload.ToQuery(), ch.IsTestNet())
}

// ListNftsByCollection gets multiple NFTs for a collection.
// @param collectionSlug: required: Unique string to identify a collection on OpenSea.
// - This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
// @param limit: The number of NFTs to return. Must be between 1 and 200. Default: 50
// @param next: The cursor for the next page of results. This is returned from a previous request.
// DOC: https://docs.opensea.io/reference/list_nfts_by_collection
func (c *client) ListNftsByCollection(ctx context.Context, payload *openseamodels.CollectionPayload,
	opts ...RequestOptionFn) (resp *openseamodels.NftsResponse, err error) {

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	if err = payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	// GET /api/v2/collection/{collection_slug}/nfts
	url := fmt.Sprintf("%s/api/v2/collection/%s/nfts",
		openseaapiutils.GetBaseURL(o.testnets), payload.CollectionSlug)

	return c.getNfts(ctx, url, payload.ToQuery(), o.testnets)
}

// GetNft gets metadata, traits, ownership information, and rarity for a single NFT.
// @param ch (chain.Chain): required: The blockchain on which to filter the results.
// @param address: required: The unique public blockchain identifier for the contract.
// @param identifier: required: The NFT token id.
// DOC: https://docs.opensea.io/reference/get_nft
func (c *client) GetNft(ctx context.Context, ch chain.Chain,
	payload *openseamodels.GetNftPayload) (resp *openseamodels.NftResponse, err error) {

	if err = payload.Validate(); err != nil {
		return nil, errx.Wrap(err, "invalid payload")
	}

	// GET /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}
	url := fmt.Sprintf("%s/api/v2/chain/%s/contract/%s/nfts/%s",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), payload.Address.String(), payload.Identifier)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	c.acceptJson(req)
	body, err := c.doRequest(req, ch.IsTestNet())
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.NftResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}

// RefreshNftMetadata refreshes metadata for a single NFT.
// @param ch (chain.Chain): required: The blockchain on which to filter the results.
// @param address: required: The unique public blockchain identifier for the contract.
// @param identifier: required: The NFT token id.
// DOC: https://docs.opensea.io/reference/refresh_nft
func (c *client) RefreshNftMetadata(ctx context.Context,
	ch chain.Chain, address common.Address, identifier string) error {

	if !common.IsHexAddress(address.String()) {
		return errx.New("invalid address")
	}
	if identifier == "" {
		return errx.New("identifier must not be empty")
	}

	// POST /api/v2/chain/{chain}/contract/{address}/nfts/{identifier}/refresh
	url := fmt.Sprintf("%s/api/v2/chain/%s/contract/%s/nfts/%s/refresh",
		openseaapiutils.GetBaseURLByChain(ch), ch.Value(), address.String(), identifier)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return errx.WithStack(err)
	}

	_, err = c.doRequest(req, ch.IsTestNet())
	return errx.WithStack(err)
}

func (c *client) getNfts(ctx context.Context, url string,
	query neturl.Values, testnets bool) (resp *openseamodels.NftsResponse, err error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	if len(query) > 0 {
		req.URL.RawQuery = query.Encode()
	}

	c.acceptJson(req)
	body, err := c.doRequest(req, testnets)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.NftsResponse)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}
