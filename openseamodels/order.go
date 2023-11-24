package openseamodels

import (
	"github.com/xTransact/errx/v2"

	"github.com/xTransact/openseaapi/chain"
)

type GetOrderPayload struct {
	// The blockchain on which to filter the results.
	Chain string `json:"chain"`
	// The hash of the order to retrieve.
	OrderHash string `json:"order_hash"`
	// The contract address of the protocol to use in the request.
	ProtocolAddress string `json:"protocol_address"`
}

func (p *GetOrderPayload) Validate() error {
	if p.Chain == "" {
		return errx.New("chain must not be empty")
	}
	_, err := chain.NewFromString(p.Chain)
	if err != nil {
		return err
	}
	if p.OrderHash == "" {
		return errx.New("order_hash must not be empty")
	}
	if p.ProtocolAddress == "" {
		return errx.New("protocol_address must not be empty")
	}
	return nil
}

type GetOrderResponse struct {
	OrderHash string `json:"order_hash"`
	// An enumeration.
	// "listing" "auction" "item_offer" "collection_offer" "trait_offer"
	Type            any           `json:"type"`
	Price           *Price        `json:"price"`
	ProtocolData    *ProtocolData `json:"protocol_data"`
	ProtocolAddress string        `json:"protocol_address"`
}
