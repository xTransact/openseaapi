package openseamodels

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/xTransact/errx/v2"
)

type Nfts struct {
	Nfts []*Nft `json:"nfts"`
}

type Nft struct {
	// The NFT's unique identifier within the smart contract (also referred to as token_id)
	Identifier string `json:"identifier"`
	// Collection slug. A unique string to identify a collection on OpenSea
	Collection string `json:"collection"`
	// The unique public blockchain identifier for the contract
	Contract string `json:"contract"`
	// ERC standard of the token (erc721, erc1155)
	TokenStandard string `json:"token_standard"`
	// Name of the NFT
	Name string `json:"name"`
	// Description of the NFT
	Description string `json:"description"`
	// Link to the image associated with the NFT
	ImageUrl string `json:"image_url"`
	// Link to the offchain metadata store
	MetadataUrl string `json:"metadata_url"`
	// Deprecated Field
	CreatedAt string `json:"created_at"`
	// Last time that the NFT's metadata was updated by OpenSea
	UpdatedAt string `json:"updated_at"`
	// If the item is currently able to be bought or sold using OpenSea
	IsDisabled bool `json:"is_disabled"`
	// If the item is currently classified as 'Not Safe for Work' by OpenSea
	IsNsfw bool `json:"is_nsfw"`

	// Link to the NFT's original animation.
	AnimationUrl string `json:"animation_url,omitempty"`
	// If the item has been reported for suspicious activity by OpenSea
	IsSuspicious bool `json:"is_suspicious,omitempty"`
	// The unique public blockchain identifier, wallet address, for the creator
	Creator string `json:"creator,omitempty"`
	// List of Trait objects. The field will be null if the NFT has more than 50 traits
	Traits []*NftTrait `json:"traits,omitempty"`
	// List of Owners. The field will be null if the NFT has more than 50 owners
	Owners []*Owner `json:"owners,omitempty"`
	// Rarity data for the NFT
	Rarity *Rarity `json:"rarity,omitempty"`
}

type GetNftsBasePayload struct {
	*BaseQueryParams

	// The unique public blockchain identifier for the contract or wallet.
	Address common.Address `json:"address"` // required
}

func (p *GetNftsBasePayload) Validate() error {
	if !common.IsHexAddress(p.Address.String()) {
		return errx.New("invalid address")
	}
	return p.BaseQueryParams.Validate()
}

func (p *GetNftsBasePayload) ToQuery() url.Values {
	return p.BaseQueryParams.ToQuery()
}

type GetNftsByAccountPayload struct {
	*GetNftsBasePayload
	// Unique string to identify a collection on OpenSea.
	// This can be found by visiting the collection on the OpenSea website and noting the last path parameter.
	Collection string `json:"collection"`
}

func (p *GetNftsByAccountPayload) Validate() error {
	return p.GetNftsBasePayload.Validate()
}

func (p *GetNftsByAccountPayload) ToQuery() url.Values {
	q := p.GetNftsBasePayload.ToQuery()
	if q == nil {
		q = make(url.Values)
	}

	if p.Collection != "" {
		q.Set("collection", p.Collection)
	}

	return q
}

type GetNftsByContractPayload struct {
	*GetNftsBasePayload
}

type NftsResponse struct {
	*Nfts
	Next string `json:"next"`
}

type GetNftPayload struct {
	// The unique public blockchain identifier for the contract or wallet.
	Address common.Address `json:"address"`
	// The NFT token id.
	Identifier string `json:"identifier"`
}

func (p *GetNftPayload) Validate() error {
	if !common.IsHexAddress(p.Address.String()) {
		return errx.New("invalid address")
	}
	if p.Identifier == "" {
		return errx.New("identifier must not be empty")
	}
	return nil
}

type NftResponse struct {
	Nft *Nft `json:"nft"`
}
