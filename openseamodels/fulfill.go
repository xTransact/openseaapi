package openseamodels

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

type MatchAdvancedOrdersInputData struct {
	Orders            []AdvancedOrder    `json:"orders"`
	CriteriaResolvers []CriteriaResolver `json:"criteriaResolvers"`
	Fulfillments      []Fulfillment      `json:"fulfillments"`
	Recipient         common.Address     `json:"recipient"`
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
