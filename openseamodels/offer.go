package openseamodels

import (
	"net/url"
	"strconv"

	"github.com/xTransact/errx/v3"
)

type BuildOfferPayload struct {
	// The address which supplies all the items in the offer.
	Offerer string `json:"offerer"`
	// The number of offers to place.
	Quantity int `json:"quantity"`
	// Criteria for the collection or trait offer
	Criteria *Criteria `json:"criteria"`
	// Exchange contract address. Must be one of ['0x00000000000000adc04c56bf30ac9d3c0aaf14dc']
	ProtocolAddress string `json:"protocol_address"`
	// Builds the offer on OpenSea's signed zone to provide offer protections from receiving an item which is disabled from trading.
	OfferProtectionEnabled bool `json:"offer_protection_enabled"`
}

type BuildOfferResponse struct {
	// Partial set of Seaport Order Parameters
	PartialParameters *PartialParameters `json:"partialParameters"`
	// Represents a list of token ids which can be used to fulfill the criteria offer.
	// When decoded using the provided SDK function,
	// developers can now see a list of all tokens that could be used to fulfill the offer.
	EncodedTokenIDs string `json:"encoded_token_ids"`
}

func (p *BuildOfferPayload) Validate() error {
	if p.Offerer == "" {
		return errx.New("offerer must not be empty")
	}
	if p.Criteria == nil || p.Criteria.Collection == nil || p.Criteria.Contract == nil {
		return errx.New("illegal arguments: nil")
	}
	if p.Criteria.Collection.Slug == "" {
		return errx.New("collection slug must not be empty")
	}
	if p.Criteria.Contract.Address == "" {
		return errx.New("contract address must not be empty")
	}
	if p.Criteria.Trait != nil {
		if p.Criteria.Trait.Type == "" {
			return errx.New("trait type must not be empty")
		}
		if p.Criteria.Trait.Value == "" {
			return errx.New("trait value must not be empty")
		}
	}
	if p.ProtocolAddress == "" {
		return errx.New("protocol address must not be empty")
	}

	return nil
}

type CreateCriteriaOfferPayload struct {
	ProtocolData    *ProtocolData `json:"protocol_data"`
	Criteria        *Criteria     `json:"criteria"`
	ProtocolAddress string        `json:"protocol_address"`
}

func (p *CreateCriteriaOfferPayload) Validate() error {
	if p.ProtocolData == nil || p.Criteria == nil {
		return errx.New("illegal arguments: nil")
	}

	if p.ProtocolAddress == "" {
		return errx.New("protocol_address must not be empty")
	}

	return nil
}

type OfferResponse struct {
	// Order hash
	OrderHash string `json:"order_hash"`
	// OpenSea supported chains.
	Chain any `json:"chain"`
	// Criteria for collection or trait offers
	Criteria *Criteria `json:"criteria"`
	// The onchain order data.
	ProtocolData *ProtocolData `json:"protocol_data"`
	// Exchange contract address
	ProtocolAddress string `json:"protocol_address"`
}

type Offers struct {
	Offers []OfferResponse `json:"offers"`
}

type PageableOffers struct {
	Offers []OfferResponse `json:"offers"`
	Next   string          `json:"next"`
}

type GetTraitOffersPayload struct {
	// CollectionSlug: required: Unique string to identify a collection on OpenSea.
	// This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	CollectionSlug string `json:"collection_slug"`
	// The value of the trait (e.g. 0.5).
	// This is only used for decimal-based numeric traits to ensure it is parsed correctly.
	FloatValue *float64 `json:"float_value"`
	// The value of the trait (e.g. 10).
	// This is only used for numeric traits to ensure it is parsed correctly.
	IntValue *int `json:"int_value"`
	// The name of the trait (e.g. 'Background')
	Type *string `json:"type"`
	// The value of the trait (e.g. 'Red')
	Value *string `json:"value"`
}

func (p *GetTraitOffersPayload) Validate() error {
	if p.CollectionSlug == "" {
		return errx.New("collection_slug must not be empty")
	}
	return nil
}

func (p *GetTraitOffersPayload) ToQuery() url.Values {
	q := make(url.Values)

	if p.FloatValue != nil {
		q.Set("float_value", strconv.FormatFloat(*p.FloatValue, 'f', -1, 64))
	}
	if p.IntValue != nil {
		q.Set("int_value", strconv.Itoa(*p.IntValue))
	}
	if p.Type != nil {
		q.Set("type", *p.Type)
	}
	if p.Value != nil {
		q.Set("value", *p.Value)
	}

	return q
}
