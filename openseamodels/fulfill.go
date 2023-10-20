package openseamodels

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/xTransact/openseaapi/openseaapiutils"
)

type FulfillmentDataResponse struct {
	Protocol        string           `json:"protocol"`
	FulfillmentData *FulfillmentData `json:"fulfillment_data"`
}

type FulfillmentData struct {
	//  Transaction: required: The name of the fulfillment method and associated call data.
	Transaction *FulfillmentTransaction `json:"transaction"`
	// Orders: required: Array of Seaport Orders.
	Orders []*Order `json:"orders"`
}

type FulfillmentTransaction struct {
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

type BasicOrderFulfillmentDataResponse struct {
	Protocol        string                     `json:"protocol"`
	FulfillmentData *BasicOrderFulfillmentData `json:"fulfillment_data"`
}

type BasicOrderFulfillmentData struct {
	//  Transaction: required: The name of the fulfillment method and associated call data.
	Transaction *BasicOrderFulfillmentTransaction `json:"transaction"`
	// Orders: required: Array of Seaport Orders.
	Orders []*Order `json:"orders"`
}

type BasicOrderFulfillmentTransaction struct {
	// Function: required: Seaport protocol contract method to use to fulfill the order
	Function string `json:"function"`
	// Chain: required: Numeric Chain Identifier
	Chain int `json:"chain"`
	// To: required: Protocol contract address to use fto fulfill the order.
	To string `json:"to"`
	// Value: required: Wei value of the transaction
	Value *big.Int `json:"value"`
	// InputData: required: Decoded Call Data.
	InputData *InputDataBasicOrderParameters `json:"input_data"`
}

type AdvancedOrderFulfillmentDataResponse struct {
	Protocol        string                        `json:"protocol"`
	FulfillmentData *AdvancedOrderFulfillmentData `json:"fulfillment_data"`
}

type AdvancedOrderFulfillmentData struct {
	Transaction *AdvancedOrderFulfillmentTransaction `json:"transaction"`
	Orders      []*AdvancedOrder                     `json:"orders"`
}

type AdvancedOrderFulfillmentTransaction struct {
	// Function: required: Seaport protocol contract method to use to fulfill the order
	Function string `json:"function"`
	// Chain: required: Numeric Chain Identifier
	Chain int `json:"chain"`
	// To: required: Protocol contract address to use fto fulfill the order.
	To string `json:"to"`
	// Value: required: Wei value of the transaction
	Value *big.Int `json:"value"`
	// InputData: required: Decoded Call Data.
	InputData *AdvancedOrdersInputData `json:"input_data"`
}

type OrderInputData struct {
	Order               *Order `json:"order"`
	FulfillerConduitKey string `json:"fulfillerConduitKey"`
}

type AdvancedOrdersInputData struct {
	AdvancedOrder       *AdvancedOrder     `json:"advancedOrder"`
	CriteriaResolvers   []CriteriaResolver `json:"criteriaResolvers"`
	FulfillerConduitKey string             `json:"fulfillerConduitKey"`
	Recipient           common.Address     `json:"recipient"`
}

type AvailableOrdersInputData struct {
	Orders                    []Order                  `json:"orders"`
	OfferFulfillments         [][]FulfillmentComponent `json:"offerFulfillments"`
	ConsiderationFulfillments [][]FulfillmentComponent `json:"considerationFulfillments"`
	FulfillerConduitKey       string                   `json:"fulfillerConduitKey"`
	MaximumFulfilled          string                   `json:"maximumFulfilled"`
}

type AvailableAdvancedOrdersInputData struct {
	AdvancedOrders            []AdvancedOrder          `json:"advancedOrders"`
	CriteriaResolvers         []CriteriaResolver       `json:"criteriaResolvers"`
	OfferFulfillments         [][]FulfillmentComponent `json:"offerFulfillments"`
	ConsiderationFulfillments [][]FulfillmentComponent `json:"considerationFulfillments"`
	FulfillerConduitKey       string                   `json:"fulfillerConduitKey"`
	Recipient                 common.Address           `json:"recipient"`
	MaximumFulfilled          string                   `json:"maximumFulfilled"`
}

type MatchOrdersInputData struct {
	Fulfillments []Fulfillment `json:"fulfillments"`
	Orders       []Order       `json:"orders"`
}

type MatchAdvancedOrdersInputData struct {
	Orders            []AdvancedOrder    `json:"orders"`
	CriteriaResolvers []CriteriaResolver `json:"criteriaResolvers"`
	Fulfillments      []Fulfillment      `json:"fulfillments"`
	Recipient         common.Address     `json:"recipient"`
}

func (t *FulfillmentTransaction) FunctionName() string {
	return openseaapiutils.GetMethod(t.Function)
}

func (t *FulfillmentTransaction) ParseInputDataToBasicOrder() (p *InputDataBasicOrderParameters, err error) {
	if t.InputData == nil {
		return &InputDataBasicOrderParameters{}, nil
	}

	data, err := json.Marshal(t.InputData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input data: %w", err)
	}

	p = new(InputDataBasicOrderParameters)
	if err = json.Unmarshal(data, p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input data: %w", err)
	}

	return p, nil
}

func (t *FulfillmentTransaction) ParseInputDataToAdvancedOrders() (p *AdvancedOrdersInputData, err error) {
	if t.InputData == nil {
		return &AdvancedOrdersInputData{}, nil
	}

	data, err := json.Marshal(t.InputData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input data: %w", err)
	}

	p = new(AdvancedOrdersInputData)
	if err = json.Unmarshal(data, p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input data: %w", err)
	}

	return p, nil
}

func (t *FulfillmentTransaction) ParseInputDataToAvailableOrders() (p *AvailableOrdersInputData, err error) {
	if t.InputData == nil {
		return &AvailableOrdersInputData{}, nil
	}

	data, err := json.Marshal(t.InputData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input data: %w", err)
	}

	p = new(AvailableOrdersInputData)
	if err = json.Unmarshal(data, p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input data: %w", err)
	}

	return p, nil
}

func (t *FulfillmentTransaction) ParseInputDataToAvailableAdvancedOrders() (p *AvailableAdvancedOrdersInputData, err error) {
	if t.InputData == nil {
		return &AvailableAdvancedOrdersInputData{}, nil
	}

	data, err := json.Marshal(t.InputData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input data: %w", err)
	}

	p = new(AvailableAdvancedOrdersInputData)
	if err = json.Unmarshal(data, p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input data: %w", err)
	}

	return p, nil
}

func (t *FulfillmentTransaction) ParseInputDataToMatchAdvancedOrders() (p *MatchAdvancedOrdersInputData, err error) {
	if t.InputData == nil {
		return &MatchAdvancedOrdersInputData{}, nil
	}

	data, err := json.Marshal(t.InputData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input data: %w", err)
	}

	p = new(MatchAdvancedOrdersInputData)
	if err = json.Unmarshal(data, p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input data: %w", err)
	}

	return p, nil
}

type Fulfiller struct {
	// Address: required: Fulfiller address.
	Address string `json:"address"`
}

type FulfillOrder struct {
	// Hash: required: Hash of the order to fulfill
	Hash string `json:"hash"`
	// Chain: required
	Chain           string `json:"chain"`
	ProtocolAddress string `json:"protocol_address"`
}

type FulfillConsideration struct {
	AssetContractAddress common.Address `json:"asset_contract_address"`
	TokenID              string         `json:"token_id"`
}

type FulfillOfferPayload struct {
	Offer         *FulfillOrder         `json:"offer"`
	Fulfiller     *Fulfiller            `json:"fulfiller"`
	Consideration *FulfillConsideration `json:"consideration"`
}

func (p *FulfillOfferPayload) Validate() error {
	if p.Offer == nil || p.Fulfiller == nil || p.Consideration == nil {
		return errors.New("illegal arguments: nil")
	}

	if p.Offer.Hash == "" {
		return errors.New("hash must not be empty")
	}
	if p.Offer.Chain == "" {
		return errors.New("chain must not be empty")
	}
	if p.Offer.ProtocolAddress == "" {
		return errors.New("protocol_address must not be empty")
	}

	if p.Fulfiller.Address == "" {
		return errors.New("fulfiller address must not be empty")
	}

	if p.Consideration.AssetContractAddress.String() == "" {
		return errors.New("consideration asset_contract_address must not be empty")
	}
	if p.Consideration.TokenID == "" {
		return errors.New("consideration token_id must not be empty")
	}

	return nil
}
