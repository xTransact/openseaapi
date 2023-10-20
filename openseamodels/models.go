package openseamodels

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseaenums"
)

// Transaction is the name of the fulfillment method and associated call data.
type Transaction struct {
	// Function: required: Seaport protocol contract method to use to fulfill the order
	Function string `json:"function"`
	// Chain: required: Numeric Chain Identifier
	Chain int `json:"chain"`
	// To: required: Protocol contract address to use fto fulfill the order.
	To string `json:"to"`
	// Value: required: Wei value of the transaction
	Value *big.Int `json:"value"`
	// InputData: required: Decoded Call Data.
	InputData any `json:"input_data"`
}

type Current struct {
	Currency string `json:"currency"`
	Decimals int64  `json:"decimals"`
	Value    string `json:"value"`
}

type Owner struct {
	// The unique public blockchain identifier for the owner wallet
	Address common.Address `json:"address"`
	// The number of tokens owned
	Quantity int `json:"quantity"`
}

// Contract defines the Contract's Addresses and Chain
type Contract struct {
	// The unique public blockchain identifier for the contract
	Address common.Address `json:"address"`
	// The chain on which the contract exists
	Chain string `json:"chain"`
	// A unique string (collection slug) to identify a collection on OpenSea
	Collection string `json:"collection,omitempty"`
	// The standard of the contract (e.g., ERC721, ERC1155)
	ContractStandaard string `json:"contract_standaard,omitempty"`
	// The name of the contract
	Name string `json:"name,omitempty"`
	// The total supply of the contract
	Supply int `json:"supply,omitempty"`
}

type ContractAddress struct {
	Address string `json:"address"`
}

type RankingFeature struct {
	// Deprecated Field.
	UniqueAttributeCount int `json:"unique_attribute_count"`
}

// Rarity data for the NFT
type Rarity struct {
	// Rarity algorithm used. Currently, always 'openrarity' (but in fact, the value may be null)
	StrategyId any `json:"strategy_id"`
	// Deprecated Field
	StrategyVersion *string `json:"strategy_version"`
	// Rarity : required: Rank of the NFT in the collection
	Rank int `json:"rank"`
	// Deprecated Field: should be a number
	Score any `json:"score"`
	// Deprecated Field
	CalculatedAt string `json:"calculated_at"`
	// Deprecated Field: should be an integer
	MaxRank      any `json:"max_rank"`
	TokensScored int `json:"tokens_scored"`
	// Deprecated Field
	RankingFeatures *RankingFeature `json:"ranking_features"`
}

type BaseQueryParams struct {
	// The number of NFTs to return. Must be between 1 and 200. Default: 50
	Limit int `json:"limit"`
	// The cursor for the next page of results. This is returned from a previous request.
	Next string `json:"next"`
}

func (b *BaseQueryParams) Validate() error {
	if b.Limit != 0 && (b.Limit < 1 || b.Limit > 200) {
		return errors.New("limit must be between 1 and 200")
	}
	return nil
}

func (b *BaseQueryParams) ToQuery() url.Values {
	q := make(url.Values)

	if b.Limit != 0 {
		q.Set("limit", strconv.Itoa(b.Limit))
	}
	if b.Next != "" {
		q.Set("next", b.Next)
	}

	return q
}

type Order struct {
	Parameters *Parameters `json:"parameters"`
	// Signature: required: The order maker's signature used to validate the order.
	Signature string `json:"signature"`
}

type AdvancedOrder struct {
	*Order

	Numerator   uint64 `json:"numerator,omitempty"`
	Denominator uint64 `json:"denominator,omitempty"`
	ExtraData   string `json:"extraData,omitempty"`
}

type BaseOfferAndConsideration struct {
	ItemType openseaenums.ItemType `json:"itemType"`
	// Token: required: The item's token contract (with the null address used for native tokens)
	Token common.Address `json:"token"`
	// IdentifierOrCriteria: required: string or int
	// The ERC721 or ERC1155 token identifier or, in the case of a criteria-based item type,
	// a merkle root composed of the valid set of token identifiers for the item.
	// This value will be ignored for Ether and ERC20 item types,
	// and can optionally be zero for criteria-based item types to allow for any identifier.
	IdentifierOrCriteria json.Number `json:"identifierOrCriteria"`
	// StartAmount: required: string or int: The amount of the token in question that will be required should the order be fulfilled.
	StartAmount json.Number `json:"startAmount"`
	// EndAmount: required: string or int: When endAmount differs from startAmount,
	// the realized amount is calculated linearly based on the time elapsed since the order became active.
	EndAmount json.Number `json:"endAmount"`
}

func (b *BaseOfferAndConsideration) Validate() error {
	if !openseaenums.ValidateItemType(int(b.ItemType)) {
		return errors.New("invalid offer.itemType")
	}
	if b.Token.String() == "" {
		return errors.New("offer.token must not be empty")
	}

	if b.IdentifierOrCriteria == "" {
		return errors.New("offer.identifierOrCriteria must not be nil")
	}
	if _, err := b.IdentifierOrCriteria.Int64(); err != nil {
		return errors.New("invalid offer.identifierOrCriteria")
	}

	if err := openseaapiutils.ValidateJsonNumber("startAmount", b.StartAmount); err != nil {
		return err
	}

	if err := openseaapiutils.ValidateJsonNumber("endAmount", b.EndAmount); err != nil {
		return err
	}

	return nil
}

type Offer struct {
	*BaseOfferAndConsideration
}

type InputDataBasicOrderParameters struct {
	Parameters *BasicOrderParameters `json:"parameters"`
}

type AdditionalRecipient struct {
	Amount    string `json:"amount"`
	Recipient string `json:"recipient"`
}

// BasicOrderParameters is an auto generated low-level Go binding around an user-defined struct.
type BasicOrderParameters struct {
	ConsiderationToken                string                 `json:"considerationToken"`
	ConsiderationIdentifier           string                 `json:"considerationIdentifier"`
	ConsiderationAmount               string                 `json:"considerationAmount"`
	Offerer                           string                 `json:"offerer"`
	Zone                              string                 `json:"zone"`
	OfferToken                        string                 `json:"offerToken"`
	OfferIdentifier                   string                 `json:"offerIdentifier"`
	OfferAmount                       string                 `json:"offerAmount"`
	BasicOrderType                    uint8                  `json:"basicOrderType"`
	StartTime                         string                 `json:"startTime"`
	EndTime                           string                 `json:"endTime"`
	ZoneHash                          string                 `json:"zoneHash"`
	Salt                              string                 `json:"salt"`
	OffererConduitKey                 string                 `json:"offererConduitKey"`
	FulfillerConduitKey               string                 `json:"fulfillerConduitKey"`
	TotalOriginalAdditionalRecipients string                 `json:"totalOriginalAdditionalRecipients"`
	AdditionalRecipients              []*AdditionalRecipient `json:"additionalRecipients"`
	Signature                         string                 `json:"signature"`
}

type Consideration struct {
	*BaseOfferAndConsideration
	Recipient common.Address `json:"recipient"`
}

func (c *Consideration) Validate() error {
	if c.BaseOfferAndConsideration == nil {
		return errors.New("nil parameters")
	}
	if err := c.BaseOfferAndConsideration.Validate(); err != nil {
		return err
	}
	if !openseaapiutils.ValidateAddress(c.Recipient) {
		return errors.New("invalid recipient")
	}
	return nil
}

type Parameters struct {
	// Offerer: required: The address which supplies all the items in the offer.
	Offerer       string           `json:"offerer"`
	Offer         []*Offer         `json:"offer"`
	Consideration []*Consideration `json:"consideration"`
	StartTime     json.Number      `json:"startTime"` // string or int
	EndTime       json.Number      `json:"endTime"`   // string or int
	// OrderType: required: The type of order, which determines how it can be executed.
	OrderType openseaenums.OrderType `json:"orderType"`
	// Zone: required: Optional secondary account attached the order which can cancel orders.
	// Additionally, when the OrderType is Restricted, the zone or the offerer are the only entities which can execute the order.
	// For open orders, use the zero address.
	// For restricted orders, use the signed zone address <SIGNED_ZONE_ADDRESS>
	Zone string `json:"zone"`
	// ZoneHash: required: A value that will be supplied to the zone when fulfilling restricted orders
	// that the zone can utilize when making a determination on whether to authorize the order.
	// Most often this value will be the zero hash 0x0000000000000000000000000000000000000000000000000000000000000000
	ZoneHash string `json:"zoneHash"`
	// Salt: required: an arbitrary source of entropy for the order
	Salt string `json:"salt"`
	// ConduitKey: required: Indicates what conduit, if any, should be utilized as a source for
	// token approvals when performing transfers.
	// By default (i.e. when conduitKey is set to the zero hash), the offerer will grant transfer approvals to Seaport directly.
	// To utilize OpenSea's conduit, use 0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000
	ConduitKey string `json:"conduitKey"`
	// TotalOriginalConsiderationItems: required: Size of the consideration array.
	TotalOriginalConsiderationItems int         `json:"totalOriginalConsiderationItems"`
	Counter                         json.Number `json:"counter,omitempty"` // any could be an integer or string
}

func (p *Parameters) Validate() error {
	if p.Offerer == "" {
		return errors.New("offerer must not be empty")
	}
	for _, offer := range p.Offer {
		if offer == nil {
			continue
		}
		if err := offer.Validate(); err != nil {
			return err
		}
	}
	for _, consideration := range p.Consideration {
		if consideration == nil {
			continue
		}
		if err := consideration.Validate(); err != nil {
			return err
		}
	}

	if err := openseaapiutils.ValidateJsonNumber("startTime", p.StartTime); err != nil {
		return err
	}

	if err := openseaapiutils.ValidateJsonNumber("endTime", p.EndTime); err != nil {
		return err
	}

	if !openseaenums.ValidateOrderType(int(p.OrderType)) {
		return errors.New("invalid orderType")
	}

	if p.Zone == "" {
		return errors.New("zone must not be empty")
	}
	if p.ZoneHash == "" {
		return errors.New("zoneHash must not be empty")
	}
	if p.Salt == "" {
		return errors.New("salt must not be empty")
	}
	if p.ConduitKey == "" {
		return errors.New("conduitKey must not be empty")
	}

	if p.Counter == "" {
		return errors.New("counter must not be empty")
	}
	c, err := p.Counter.Int64()
	if err != nil {
		return fmt.Errorf("invalid counter: %w", err)
	}
	if c < 0 {
		return errors.New("invalid counter")
	}

	return nil
}

// PartialParameters : Partial set of Seaport Order Parameters
type PartialParameters struct {
	Considerations []Consideration `json:"considerations"`
	// Optional secondary account attached the order which can cancel orders.
	// Additionally, when the OrderType is Restricted, the zone or the offerer are the only entities which can execute the order.
	// For open orders, use the zero address.
	// For restricted orders, use the signed zone address 0x000000e7ec00e7b300774b00001314b8610022b8
	Zone string `json:"zone"`
	// A value that will be supplied to the zone when fulfilling restricted orders that the zone can utilize
	// when making a determination on whether to authorize the order.
	// Most often this value will be the zero hash 0x0000000000000000000000000000000000000000000000000000000000000000
	ZoneHash string `json:"zoneHash"`
}

type Identifier struct {
	User int `json:"user"`
	// A placeholder image. For the actual profile image, call the Get Account endpoint.
	ProfileImgUrl string `json:"profile_img_url"`
	// The unique blockchain identifier, address, of the account.
	Address string `json:"address"`
	// affiliate - affiliate
	// affiliate_partner - affiliate_partner
	// affiliate_requested - affiliate_requested
	// affiliate_blacklisted - affiliate_blacklisted
	// verified - verified
	// moderator - moderator
	// staff - staff
	// employee - employee
	Config string `json:"config"`
}

type Fee struct {
	Account     *Identifier `json:"account"`
	BasisPoints string      `json:"basis_points"`
}

// ReceivedItem is an auto generated low-level Go binding around an user-defined struct.
type ReceivedItem struct {
	ItemType   uint8          `json:"itemType"`
	Token      common.Address `json:"token"`
	Identifier json.Number    `json:"identifier"`
	Amount     json.Number    `json:"amount"`
	Recipient  common.Address `json:"recipient"`
}

// CriteriaResolver is an auto generated low-level Go binding around an user-defined struct.
type CriteriaResolver struct {
	OrderIndex    json.Number            `json:"orderIndex"`
	Side          openseaenums.OrderSide `json:"side"`
	Index         json.Number            `json:"index"`
	Identifier    json.Number            `json:"identifier"`
	CriteriaProof []string               `json:"criteriaProof"`
}

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Item       ReceivedItem   `json:"item"`
	Offerer    common.Address `json:"offerer"`
	ConduitKey string         `json:"conduitKey"`
}

// Fulfillment is an auto generated low-level Go binding around an user-defined struct.
type Fulfillment struct {
	ConsiderationComponents []FulfillmentComponent `json:"considerationComponents"`
	OfferComponents         []FulfillmentComponent `json:"offerComponents"`
}

// FulfillmentComponent is an auto generated low-level Go binding around an user-defined struct.
type FulfillmentComponent struct {
	OrderIndex json.Number `json:"orderIndex"`
	ItemIndex  json.Number `json:"itemIndex"`
}

type CollectionSlug struct {
	// Unique string to identify a collection on OpenSea.
	// This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	Slug string `json:"slug"`
}

type CriteriaTrait struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Criteria struct {
	// The collection in which the criteria offer is being made for.
	Collection *CollectionSlug `json:"collection"`
	// The unique public blockchain identifier, address, for the NFT contract
	Contract *ContractAddress `json:"contract"`
	// The trait that the criteria offer is being made for.
	Trait *CriteriaTrait `json:"trait"`
	// Represents a list of token ids which can be used to fulfill the criteria offer.
	// When decoded using the provided SDK function,
	// developers can now see a list of all tokens that could be used to fulfill the offer.
	EncodedTokenIDs string `json:"encoded_token_ids"`
}

type Payment struct {
	// Amount of tokens in the order
	Quantity int `json:"quantity"`
	// The contract address for the ERC20 token
	TokenAddress string `json:"token_address"`
	// Returns the number of decimals the token uses -
	// e.g. 8, means to divide the token amount by 100000000 to get its user representation.
	Decimals *big.Int `json:"decimals"`
	// Returns the symbol of the token, e.g. ETH, WETH, USDC, etc
	Symbol string `json:"symbol"`
}
