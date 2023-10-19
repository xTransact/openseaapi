package openseaapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/net/http2"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

type Servicer interface {
	/* NFT Endpoints */

	// GetAccount gets an OpenSea Account Profile including details such as bio, social media usernames, and profile image.
	GetAccount(ctx context.Context, address common.Address,
		opts ...RequestOptionFn) (resp *openseamodels.Account, err error)
	ListNftsByAccount(ctx context.Context, ch chain.Chain,
		payload *openseamodels.GetNftsByAccountPayload) (resp *openseamodels.NftsResponse, err error)
	GetContract(ctx context.Context, ch chain.Chain, address common.Address) (
		resp *openseamodels.Contract, err error)
	ListNftsByContract(ctx context.Context, ch chain.Chain,
		payload *openseamodels.GetNftsByContractPayload) (resp *openseamodels.NftsResponse, err error)
	GetNft(ctx context.Context, ch chain.Chain,
		payload *openseamodels.GetNftPayload) (resp *openseamodels.NftResponse, err error)
	RefreshNftMetadata(ctx context.Context,
		ch chain.Chain, address common.Address, identifier string) error
	ListNftsByCollection(ctx context.Context, payload *openseamodels.CollectionPayload,
		opts ...RequestOptionFn) (resp *openseamodels.NftsResponse, err error)
	ListCollections(ctx context.Context, payload *openseamodels.ListCollectionsPayload) (
		resp *openseamodels.CollectionsResponse, err error)
	GetCollection(ctx context.Context, collectionSlug string, opts ...RequestOptionFn) (
		resp *openseamodels.SingleCollection, err error)
	GetTraits(ctx context.Context, collectionSlug string, opts ...RequestOptionFn) (
		resp *openseamodels.Trait, err error)

	/*  Analytics Endpoints */

	GetCollectionStats(ctx context.Context, collectionSlug string, opts ...RequestOptionFn) (
		resp *openseamodels.CollectionStats, err error)
	// ListEventsByAccount
	// ListEventsByNft
	// ListEventsByCollection

	/* OpenSea Marketplace Endpoints */

	BuildOffer(ctx context.Context, payload *openseamodels.BuildOfferPayload,
		opts ...RequestOptionFn) (resp *openseamodels.BuildOfferResponse, err error)
	GetCollectionOffers(ctx context.Context, collectionSlug string,
		opts ...RequestOptionFn) (resp *openseamodels.Offers, err error)
	CreateCriteriaOffer(ctx context.Context, payload *openseamodels.CreateCriteriaOfferPayload,
		opts ...RequestOptionFn) (resp *openseamodels.OfferResponse, err error)
	CreateIndividualOffer(ctx context.Context, ch chain.Chain,
		payload *openseamodels.CreateOrderPayload) (resp *openseamodels.OrderResponse, err error)
	// CreateListing lists a single NFT (ERC721 or ERC1155) for sale on the OpenSea marketplace.
	CreateListing(ctx context.Context, ch chain.Chain,
		payload *openseamodels.CreateOrderPayload) (resp *openseamodels.CreateListingResponse, err error)
	// FulfillListing retrieves all the information, including signatures, needed to fulfill a listing directly onchain.
	FulfillListing(ctx context.Context, ch chain.Chain,
		orderHash, fulfiller string) (resp *openseamodels.FulfillmentDataResponse, err error)
	FulfillOffer(ctx context.Context, payload *openseamodels.FulfillOfferPayload,
		opts ...RequestOptionFn) (resp *openseamodels.FulfillmentDataResponse, err error)
	// GetAllListingsByCollection gets all active, valid listings for a single collection.
	GetAllListingsByCollection(ctx context.Context, payload *openseamodels.GetAllListingsByCollectionPayload,
		opts ...RequestOptionFn) (resp *openseamodels.ListingsByCollectionResponse, err error)
	GetAllCollectionOffers(ctx context.Context, payload *openseamodels.CollectionPayload,
		opts ...RequestOptionFn) (resp *openseamodels.PageableOffers, err error)
	GetIndividualOffers(ctx context.Context, ch chain.Chain,
		payload *openseamodels.OrderPayload) (resp *openseamodels.OrdersResponse, err error)
	// GetListings gets the complete set of active, valid listings.
	GetListings(ctx context.Context, ch chain.Chain,
		payload *openseamodels.OrderPayload) (resp *openseamodels.OrdersResponse, err error)
	GetOrder(ctx context.Context, payload *openseamodels.GetOrderPayload) (
		resp *openseamodels.GetOrderResponse, err error)
	GetTraitOffers(ctx context.Context, payload *openseamodels.GetTraitOffersPayload,
		opts ...RequestOptionFn) (resp *openseamodels.Offers, err error)
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
