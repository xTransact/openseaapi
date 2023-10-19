package openseamodels

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaenums"
)

type Collection struct {
	// Collection slug. A unique string to identify a collection on OpenSea
	Collection string `json:"collection"`
	// Name of the collection
	Name string `json:"name"`
	// Description of the collection
	Description string `json:"description"`
	// Square image used to represent the collection
	ImageUrl string `json:"image_url"`
	// Banner image used to represent the collection
	BannerImageUrl string `json:"banner_image_url"`
	// The unique public blockchain identifier, address, for the owner wallet.
	Owner string `json:"owner"`
	// Status of the collection verification requests.
	// not_requested requested approved verified disabled_top_trending
	SafelistStatus openseaenums.SafelistStatus `json:"safelist_status"`
	// Category of the collection (e.g. PFPs, Memberships, Art)
	Category string `json:"category"`
	// If the collection is currently able to be bought or sold using OpenSea
	IsDisabled bool `json:"is_disabled"`
	// If the collection is currently classified as 'Not Safe for Work' by OpenSea
	IsNsfw bool `json:"is_nsfw"`
	// If trait offers are currently being accepted for the collection
	TraitOffersEnabled bool `json:"trait_offers_enabled"`
	// OpenSea Link to collection
	OpenSeaUrl string `json:"opensea_url"`
	// External URL for the collection's website
	ProjectUrl string `json:"project_url"`
	// External URL for the collection's wiki
	WikiUrl string `json:"wiki_url"`
	// External URL for the collection's Discord server
	DiscordUrl string `json:"discord_url"`
	// External URL for the collection's Telegram group
	TelegramUrl string `json:"telegram_url"`
	// Username for the collection's Twitter account
	TwitterUsername string `json:"twitter_username"`
	// Username for the collection's Instagram account
	InstagramUsername string `json:"instagram_username"`
	// Define the Contract's Addresses and Chain
	Contracts []*Contract `json:"contracts"`
}

type CollectionFee struct {
	// Percentage of the sale price that is paid to the recipient
	Fee decimal.Decimal `json:"fee"`
	// The unique public blockchain identifier, address, for the recipient
	Recipient common.Address `json:"recipient"`
	// If the fee is required for the collection
	Required bool `json:"required"`
}

type SingleCollection struct {
	*Collection

	// List of editor addresses for the collection
	Editors []common.Address `json:"editors"`
	// List of fees for the collection including creator earnings and OpenSea fees
	Fees []*CollectionFee `json:"fees"`
}

type CollectionPayload struct {
	*BaseQueryParams

	CollectionSlug string `json:"collection_slug"`
}

func (p *CollectionPayload) Validate() error {
	if p.CollectionSlug == "" {
		return errors.New("collection_slug must not be empty")
	}
	return p.BaseQueryParams.Validate()
}

type ListCollectionsPayload struct {
	// The blockchain on which to filter the results
	ChainIdentifier string `json:"chain_identifier"`
	// If true, will return hidden collections. Default: false
	IncludeHidden *bool `json:"include_hidden"`
	// The number of collections to return. Must be between 1 and 100. Default: 100
	Limit int `json:"limit"`
	// The cursor for the next page of results. This is returned from a previous request.
	Next string `json:"next"`
}

func (p *ListCollectionsPayload) Validate() error {
	if _, err := chain.NewFromString(p.ChainIdentifier); err != nil {
		return err
	}

	if p.Limit != 0 && (p.Limit < 1 || p.Limit > 100) {
		return errors.New("limit must be between 0 and 100")
	}

	return nil
}

func (p *ListCollectionsPayload) ToQuery() url.Values {
	q := make(url.Values)

	if p.ChainIdentifier != "" {
		q.Set("chain_identifier", p.ChainIdentifier)
	}
	if p.IncludeHidden != nil {
		q.Set("include_hidden", strconv.FormatBool(*p.IncludeHidden))
	}
	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(p.Limit))
	}
	if p.Next != "" {
		q.Set("next", p.Next)
	}

	return q
}

type CollectionsResponse struct {
	Collections []*Collection `json:"collections"`
	Next        string        `json:"next"`
}

type CollectionStats struct {
	// The aggregate stats over the collection's lifetime
	Total *CollectionStatsTotal `json:"total"`
	// The stats for each interval
	Intervals []*CollectionStatsInterval `json:"intervals"`
}

type CollectionStatsTotal struct {
	// The all time volume of sales for the collection
	Volume decimal.Decimal `json:"volume"`
	// The all time number of sales for the collection
	Sales int64 `json:"sales"`
	// The all time average sale price of NFTs in the collection
	AveragePrice decimal.Decimal `json:"average_price"`
	// The current number of unique owners of NFTs in the collection
	NumOwners int64 `json:"num_owners"`
	// The current market cap of the collection
	MarketCap decimal.Decimal `json:"market_cap"`
	// The current lowest price of NFTs in the collection
	FloorPrice decimal.Decimal `json:"floor_price"`
	// The symbol of the payment asset for the floor price
	FloorPriceSymbol string `json:"floor_price_symbol"`
}

type CollectionStatsInterval struct {
	// The interval for which the stats are calculated
	// one_day one_week one_month
	Interval string `json:"interval"`
	// The volume of sales for the collection during the specified interval
	Volume decimal.Decimal `json:"volume"`
	// The volume differential compared to the previous interval
	VolumeDiff decimal.Decimal `json:"volume_diff"`
	// The percentage change in volume compared to the previous interval
	VolumeChange decimal.Decimal `json:"volume_change"`
	// The number of sales for the collection during the specified interval
	Sales int64 `json:"sales"`
	// The percentage change in number of sales compared to the previous interval
	SalesDiff int64 `json:"sales_diff"`
	// The average sale price of NFTs in the collection during the interval
	AveragePrice decimal.Decimal `json:"average_price"`
}
