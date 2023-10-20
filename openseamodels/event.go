package openseamodels

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaenums"
)

type GetEventsByAccountPayload struct {
	*GetEventsQueryParams

	// Address : required: The unique public blockchain identifier for the contract or wallet.
	Address string `json:"address"`
}

func (p *GetEventsByAccountPayload) Validate() error {
	if p.Address == "" {
		return errors.New("address must not be empty")
	}
	return nil
}

type GetEventsByNftPayload struct {
	*GetEventsQueryParams

	// The unique public blockchain identifier for the contract or wallet.
	Address string `json:"address"`
	// The blockchain on which to filter the results.
	Chain chain.Chain `json:"chain"`
	// The NFT token id.
	Identifier string `json:"identifier"`
}

func (p *GetEventsByNftPayload) Validate() error {
	if p.Address == "" {
		return errors.New("address must not be empty")
	}
	if p.Identifier == "" {
		return errors.New("identifier must not be empty")
	}
	return nil
}

type GetEventsByCollectionPayload struct {
	*GetEventsQueryParams

	// CollectionSlug : required: Unique string to identify a collection on OpenSea.
	// This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	CollectionSlug string `json:"collection_slug"`
}

func (p *GetEventsByCollectionPayload) Validate() error {
	if p.CollectionSlug == "" {
		return errors.New("collection_slug must not be empty")
	}
	return nil
}

type GetEventsQueryParams struct {
	// Filter to only include events that occurred at or after the given timestamp.
	// The Unix epoch timstamp must be in seconds
	After *int64 `json:"after"`
	// Filter to only include events that occurred before the given timestamp.
	// The Unix epoch timstamp must be in seconds.
	Before *int64 `json:"before"`
	// The type of event to filter by. If not provided, only sales will be returned.
	EventType []openseaenums.EventType `json:"event_type"`
	// The cursor for the next page of results. This is returned from a previous request.
	Next *string `json:"next"`
}

func (p *GetEventsQueryParams) ToQuery() url.Values {
	q := make(url.Values)

	if p.After != nil {
		q.Set("after", strconv.FormatInt(*p.After, 10))
	}
	if p.Before != nil {
		q.Set("before", strconv.FormatInt(*p.Before, 10))
	}
	for _, eventType := range p.EventType {
		q.Add("event_type", string(eventType))
	}
	if p.Next != nil {
		q.Set("next", *p.Next)
	}

	return q
}

type AssetEventResponse struct {
	AssetEvent []AssetEvent `json:"asset_event"`
	Next       string       `json:"next"`
}

type AssetEvent struct {
	EventType       openseaenums.EventType `json:"event_type"`
	OrderHash       string                 `json:"order_hash,omitempty"`
	OrderType       any                    `json:"order_type,omitempty"`
	Chain           string                 `json:"chain,omitempty"`
	Transaction     string                 `json:"transaction"`
	ProtocolAddress string                 `json:"protocol_address,omitempty"`
	FromAddress     string                 `json:"from_address,omitempty"`
	ToAddress       string                 `json:"to_address,omitempty"`
	StartDate       int64                  `json:"start_date,omitempty"`
	ClosingDate     int64                  `json:"closing_date,omitempty"`
	ExpirationDate  int64                  `json:"expiration_date,omitempty"`
	Asset           *Nft                   `json:"asset,omitempty"`
	Nft             *Nft                   `json:"nft,omitempty"`
	Quantity        int                    `json:"quantity,omitempty"`
	Maker           string                 `json:"maker,omitempty"`
	Taker           string                 `json:"taker,omitempty"`
	Payment         *Payment               `json:"payment,omitempty"`
	Criteria        *Criteria              `json:"criteria,omitempty"`
}
