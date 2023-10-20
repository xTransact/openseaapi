# openseaapi

An SDK for [OpenSea API](https://docs.opensea.io/reference/api-overview).


## Supported API

### NFT Endpoints

- [x] [Get Account](https://docs.opensea.io/reference/get_account)
- [x] [Get NFTs (by account)](https://docs.opensea.io/reference/list_nfts_by_account)
- [x] [Get Contract](https://docs.opensea.io/reference/get_contract)
- [x] [Get NFTs (by contract)](https://docs.opensea.io/reference/list_nfts_by_contract)
- [x] [Get an NFT](https://docs.opensea.io/reference/get_nft)
- [x] [Refresh NFT Metadata](https://docs.opensea.io/reference/refresh_nft)
- [x] [Get NFTs (by collection)](https://docs.opensea.io/reference/list_nfts_by_collection)
- [x] [Get Collections](https://docs.opensea.io/reference/list_collections)
- [x] [Get a Collection](https://docs.opensea.io/reference/get_collection)
- [x] [Get Traits](https://docs.opensea.io/reference/get_traits)


### Analytics Endpoints

- [x] [Get Collection Stats](https://docs.opensea.io/reference/get_collection_stats)
- [x] [Get Events (by account)](https://docs.opensea.io/reference/list_events_by_account)
- [x] [Get Events (by NFT)](https://docs.opensea.io/reference/list_events_by_nft)
- [x] [Get Events (by collection)](https://docs.opensea.io/reference/list_events_by_collection)


### OpenSea Marketplace Endpoints

- [x] [Build an Offer](https://docs.opensea.io/reference/build_offer_v2)
- [x] [Get Collection Offers](https://docs.opensea.io/reference/get_collection_offers_v2)
- [x] [Create Criteria Offer](https://docs.opensea.io/reference/post_criteria_offer_v2)
- [x] [Create Individual Offer](https://docs.opensea.io/reference/post_offer)
- [x] [Create Listing](https://docs.opensea.io/reference/post_listing)
- [x] [Fulfill a Listing](https://docs.opensea.io/reference/generate_listing_fulfillment_data_v2)
- [x] [Fullfill an Offer](https://docs.opensea.io/reference/generate_offer_fulfillment_data_v2)
- [x] [Get All Listings (by collection)](https://docs.opensea.io/reference/get_all_listings_on_collection_v2)
- [x] [Get All Offers (by collection)](https://docs.opensea.io/reference/get_all_offers_on_collection_v2)
- [x] [Get Individual Offers](https://docs.opensea.io/reference/get_offers)
- [x] [Get Listings](https://docs.opensea.io/reference/get_listings)
- [x] [Get Order](https://docs.opensea.io/reference/get_order)
- [x] [Get Trait Offers](https://docs.opensea.io/reference/get_trait_offers_v2)


## Getting Started

### Get it

```shell
go get -u github.com/reedchan7/openseaapi
```

### Examples

```go
cli := NewClient(
	WithApiKey(os.Getenv("OPENSEA_API_KEY")),
)

// Get account on mainnet by wallet address
resp, err := cli.GetAccount(ctx, addr)

// Get account on goerli testnet by wallet address
resp, err := cli.GetAccount(ctx, addr,UseTestnets())
```
