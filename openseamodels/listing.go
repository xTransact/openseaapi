package openseamodels

import (
	"encoding/json"
	"errors"
	"net/url"
	"slices"
	"strconv"
	"time"

	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseaenums"
)

const dateTimeFormat = "2006-01-02T15:04:05.000Z"

const (
	OrderByCreatedDate = "created_date"
	OrderByEthPrice    = "eth_price"

	OrderDirectionAsc  = "asc"
	OrderDirectionDesc = "desc"
)

type OrderPayload struct {
	// Filter results by the contract address for NFT(s).
	// NOTE: If used, token_ids or token_id is required.
	AssetContractAddress *string `query:"asset_contract_address"`
	// Restricts results to only include orders that are bundles of NFTs. Default: false
	Bundled *bool `query:"bundled"`
	// The cursor for the next page of results. This is returned from a previous request.
	Cursor *string `query:"cursor"`
	// The number of orders to return. Must be between 1 and 50. Default: 20
	Limit *int `query:"limit"`
	// Filter to only include orders that were listed after the given timestamp. This is a Unix epoch timestamp in seconds.
	ListedAfter *time.Time `query:"listed_after"`
	// Filter to only include orders that were listed before the given timestamp. This is a Unix epoch timestamp in seconds.
	ListedBefore *time.Time `query:"listed_before"`
	// Filter results by the order maker's wallet address.
	Maker *string `query:"maker"`
	// The order in which to sort the results. Default: created_date
	// NOTE: If eth_price is used, asset_contract_address and token_id are required.
	OrderBy *string `query:"order_by"`
	// The direction in which to sort the results. Default: desc
	OrderDirection *string `query:"order_direction"`
	// Payment Token Address to filter results. This ensures all returned orders are listed in a single currency.
	PaymentTokenAddress *string `query:"payment_token_address"`
	// Filter results by the order taker's wallet address.
	Taker *string `query:"taker"`
	// An array of token IDs to search for (e.g. ?token_ids=1&token_ids=209).
	// This endpoint will return a list of orders with token_id matching any of the IDs in this array.
	// NOTE: If used, asset_contract_address is required.
	TokenIDs []json.Number `query:"token_ids"`
}

func (p *OrderPayload) Validate() error {
	if p.AssetContractAddress != nil && len(p.TokenIDs) == 0 {
		return errors.New("token_ids must be used if asset_contract_address is not nil")
	}
	if len(p.TokenIDs) > 0 && p.AssetContractAddress == nil {
		return errors.New("asset_contract_address must not be nil if token_ids is used")
	}

	if p.Limit != nil && (*p.Limit < 1 || *p.Limit > 50) {
		return errors.New("limit must be between 1 and 50")
	}

	if p.OrderBy != nil {
		if !slices.Contains([]string{OrderByCreatedDate, OrderByEthPrice}, *p.OrderBy) {
			return errors.New("invalid order_by")
		}

		if *p.OrderBy == OrderByEthPrice && (p.AssetContractAddress == nil || len(p.TokenIDs) == 0) {
			return errors.New("asset_contract_address and token_ids are required if order_by is 'eth_price'")
		}
	}

	if p.OrderDirection != nil && !slices.Contains([]string{OrderDirectionAsc, OrderDirectionDesc}, *p.OrderDirection) {
		return errors.New("invalid order_direction")
	}

	return nil
}

// ToQuery -
func (p *OrderPayload) ToQuery() url.Values {
	q := make(url.Values)

	if p.AssetContractAddress != nil {
		q.Set("asset_contract_address", openseaapiutils.PtrToString(p.AssetContractAddress))
	}
	if p.Bundled != nil {
		q.Set("bundled", strconv.FormatBool(openseaapiutils.PtrToBool(p.Bundled)))
	}
	if p.Cursor != nil {
		q.Set("cursor", openseaapiutils.PtrToString(p.Cursor))
	}
	if p.Limit != nil {
		q.Set("limit", strconv.Itoa(openseaapiutils.PtrToInt(p.Limit)))
	}
	if p.ListedAfter != nil {
		q.Set("listed_after", p.ListedAfter.Format(dateTimeFormat))
	}
	if p.ListedBefore != nil {
		q.Set("listed_before", p.ListedBefore.Format(dateTimeFormat))
	}
	if p.Maker != nil {
		q.Set("maker", openseaapiutils.PtrToString(p.Maker))
	}
	if p.OrderBy != nil {
		q.Set("order_by", openseaapiutils.PtrToString(p.OrderBy))
	}
	if p.OrderDirection != nil {
		q.Set("order_direction", openseaapiutils.PtrToString(p.OrderDirection))
	}
	if p.PaymentTokenAddress != nil {
		q.Set("payment_token_address", openseaapiutils.PtrToString(p.PaymentTokenAddress))
	}
	if p.Taker != nil {
		q.Set("taker", openseaapiutils.PtrToString(p.Taker))
	}
	for _, id := range p.TokenIDs {
		q.Add("token_ids", id.String())
	}

	return q
}

type OrdersResponse struct {
	// The cursor for the next page of results.
	Next string `json:"next"`
	// The cursor for the previous page of results.
	Previous string `json:"previous"`
	// Models OrderV2 objects to serialize to a 'similar' schema to what we have with OrderV1s
	Orders []*OrderResponse `json:"orders"`
}

func (r *OrdersResponse) ToJson() []byte {
	data, _ := json.Marshal(r)
	return data
}

type OrderResponse struct {
	// Date the order was created
	CreatedDate string `json:"created_date"`
	// Date the order was closed
	ClosingDate string `json:"closing_date"`
	// Timestamp representation of created_date
	ListingTime int64 `json:"listing_time"`
	// Timestamp representation of closing_date
	ExpirationTime int64 `json:"expiration_time"`
	// An identifier for the order
	OrderHash    string        `json:"order_hash"`
	ProtocolData *ProtocolData `json:"protocol_data"`
	// Exchange Contract Address. Typically the address of the Seaport contract.
	ProtocolAddress string `json:"protocol_address"`
	// Current price of the order
	CurrentPrice string `json:"current_price"`
	// The unique blockchain identifier, address, of the wallet which is the order maker.
	Maker *Identifier `json:"maker"`
	// The unique blockchain identifier, address, of the wallet which is the order taker.
	Taker     *Identifier `json:"taker"`
	MakerFees []*Fee      `json:"maker_fees"`
	TakerFees []*Fee      `json:"taker_fees"`
	// The side of the order, either bid (offer) or ask(listing).
	Side string `json:"side"`
	// basic - basic
	// dutch - dutch
	// english - english
	// criteria - criteria
	OrderType string `json:"order_type"`
	// If true, the order maker has canceled the order which means it can no longer be filled.
	Cancelled bool `json:"cancelled"`
	// If true, the order has already been filled.
	Finalized bool `json:"finalized"`
	// If true, the order is currently invalid and can not be filled.
	MarkedInvalid bool `json:"marked_invalid"`
	// The remaining quantity of the order that has not been filled. This is useful for erc1155 orders.
	RemainingQuantity int `json:"remaining_quantity"`
	// Deprecated Field
	RelayId string `json:"relay_id"`
	// A merkle root composed of the valid set of token identifiers for the order
	CriteriaProof []string `json:"criteria_proof"`
	// Deprecated
	MakerAssetBundle any `json:"maker_asset_bundle,omitempty"`
	// Deprecated
	TakerAssetBundle any `json:"taker_asset_bundle,omitempty"`
}

type CreateOrderPayload struct {
	Parameters *Parameters `json:"parameters"`
	// Signature of the signed type data represented by the parameters field.
	Signature string `json:"signature"`
	// Exchange contract address. Must be one of ['0x00000000000000adc04c56bf30ac9d3c0aaf14dc']
	ProtocolAddress string `json:"protocol_address"`
}

func (p *CreateOrderPayload) Validate() error {
	if p.Parameters == nil {
		return errors.New("invalid parameters")
	}
	if p.Parameters.Offerer == "" {
		return errors.New("offerer must not be empty")
	}
	if len(p.Parameters.Offer) == 0 {
		return errors.New("offer must not be empty")
	}
	for _, o := range p.Parameters.Offer {
		if o == nil {
			return errors.New("offer must not be nil")
		}
		if err := o.Validate(); err != nil {
			return err
		}
	}
	for _, c := range p.Parameters.Consideration {
		if c == nil {
			return errors.New("consideration must not be nil")
		}
		if err := c.Validate(); err != nil {
			return err
		}
	}

	if err := openseaapiutils.ValidateJsonNumber("startTime", p.Parameters.StartTime); err != nil {
		return err
	}
	if err := openseaapiutils.ValidateJsonNumber("endTime", p.Parameters.EndTime); err != nil {
		return err
	}

	if p.Parameters.OrderType < 0 {
		return errors.New("invalid orderType")
	}
	if p.Parameters.Zone == "" {
		return errors.New("zone must not be empty")
	}
	if p.Parameters.ZoneHash == "" {
		return errors.New("zoneHash must not be empty")
	}
	if p.Parameters.Salt == "" {
		return errors.New("salt must not be empty")
	}
	if p.Parameters.ConduitKey == "" {
		return errors.New("conduitKey must not be empty")
	}
	if p.Parameters.Counter == "" {
		return errors.New("counter must not be empty")
	}

	if p.Signature == "" {
		return errors.New("signature must not be empty")
	}
	if p.ProtocolAddress == "" {
		return errors.New("protocol_address must not be empty")
	}

	return nil
}

type CreateListingResponse struct {
	Order *OrderResponse `json:"order"`
}

type ProtocolData struct {
	Parameters *Parameters `json:"parameters"`
	Signature  string      `json:"signature"`
}

func (r *OrdersResponse) IsNotFound() bool {
	return len(r.Orders) == 0
}

type GetAllListingsByCollectionPayload struct {
	// Unique string to identify a collection on OpenSea.
	// This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	CollectionSlug string `json:"collection_slug"` // required
	// The number of listings to return. Must be between 1 and 100. Default: 100
	Limit int `json:"limit"`
	// The cursor for the next page of results. This is returned from a previous request.
	Next string `json:"next"`
}

func (p *GetAllListingsByCollectionPayload) Validate() error {
	if p.CollectionSlug == "" {
		return errors.New("illegal arguments: collection slug must not be empty")
	}
	if p.Limit < 0 || p.Limit > 100 {
		return errors.New("illegal arguments: limit must be between 0 and 100")
	}

	return nil
}

func (p *GetAllListingsByCollectionPayload) ToQuery() url.Values {
	q := make(url.Values)

	if p.Limit > 0 {
		q.Set("limit", strconv.Itoa(p.Limit))
	}
	if p.Next != "" {
		q.Set("next", p.Next)
	}

	return q
}

type FulfillListingPayload struct {
	// Listing: required
	Listing *FulfillOrder `json:"listing"`
	// FulFiller: required
	FulFiller *Fulfiller `json:"fulfiller"`
}

type ListingsByCollectionResponse struct {
	Listings []*CollectionListing `json:"listings"`
	Next     string               `json:"next"`
}

type CollectionListing struct {
	OrderHash    string            `json:"order_hash"`
	Type         openseaenums.Type `json:"type"`
	Price        *Price            `json:"price"`
	ProtocolData *ProtocolData     `json:"protocol_data"`
}

type Price struct {
	Current *Current `json:"currenty"`
}
