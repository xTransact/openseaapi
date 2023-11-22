package openseamodels

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListingsConvert1(t *testing.T) {
	const data = `{
  "next": null,
  "previous": null,
  "orders": [
    {
      "created_date": "2023-11-22T08:24:35.074631",
      "closing_date": "2023-11-23T08:24:31",
      "listing_time": 1700641471,
      "expiration_time": 1700727871,
      "order_hash": "0xe5abef0267b3c4f14aac409cc98a1afff8786055123e0f08273b17592ac97c17",
      "protocol_data": {
        "parameters": {
          "offerer": "0x0097b9cfe64455eed479292671a1121f502bc954",
          "offer": [
            {
              "itemType": 2,
              "token": "0x9401518f4EBBA857BAA879D9f76E1Cc8b31ed197",
              "identifierOrCriteria": "5333",
              "startAmount": "1",
              "endAmount": "1"
            }
          ],
          "consideration": [
            {
              "itemType": 0,
              "token": "0x0000000000000000000000000000000000000000",
              "identifierOrCriteria": "0",
              "startAmount": "487505225000000000",
              "endAmount": "487505225000000000",
              "recipient": "0x0097b9cFE64455EED479292671A1121F502bc954"
            },
            {
              "itemType": 0,
              "token": "0x0000000000000000000000000000000000000000",
              "identifierOrCriteria": "0",
              "startAmount": "2449775000000000",
              "endAmount": "2449775000000000",
              "recipient": "0x0000a26b00c1F0DF003000390027140000fAa719"
            }
          ],
          "startTime": "1700641471",
          "endTime": "1700727871",
          "orderType": 0,
          "zone": "0x004C00500000aD104D7DBd00e3ae0A5C00560C00",
          "zoneHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
          "salt": "0x72db8c0b00000000000000000000000000000000000000000b2fe904d133b02c",
          "conduitKey": "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
          "totalOriginalConsiderationItems": 2,
          "counter": "0x3fbc81a0b0ffd9cf3cef372d93bfc35f5"
        },
        "signature": null
      },
      "protocol_address": "0x00000000000000adc04c56bf30ac9d3c0aaf14dc",
      "current_price": "489955000000000000",
      "maker": {
        "user": 35212416,
        "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/31.png",
        "address": "0x0097b9cfe64455eed479292671a1121f502bc954",
        "config": ""
      },
      "taker": null,
      "maker_fees": [
        {
          "account": {
            "user": null,
            "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/29.png",
            "address": "0x0000a26b00c1f0df003000390027140000faa719",
            "config": ""
          },
          "basis_points": "50"
        }
      ],
      "taker_fees": [],
      "side": "ask",
      "order_type": "basic",
      "cancelled": false,
      "finalized": false,
      "marked_invalid": false,
      "remaining_quantity": 1,
      "relay_id": "T3JkZXJWMlR5cGU6MTQzNDg4OTUyMDA=",
      "criteria_proof": null,
      "maker_asset_bundle": {
        "assets": [
          {
            "id": 284206546,
            "token_id": "5333",
            "num_sales": 35,
            "background_color": null,
            "image_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_preview_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_thumbnail_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_original_url": "ipfs://QmTVdT52JWqEgowWLVSpxeqj1u7ZayeqGvpsjWP5jWhjsY",
            "animation_url": null,
            "animation_original_url": null,
            "name": "Lil Ghost#5333",
            "description": "We are the ghost, your nightmare, your daydream.",
            "external_link": null,
            "asset_contract": {
              "address": "0x9401518f4ebba857baa879d9f76e1cc8b31ed197",
              "asset_contract_type": "non-fungible",
              "chain_identifier": "ethereum",
              "created_date": "2022-02-09T09:20:52.962749",
              "name": "Weirdo Ghost Gang",
              "nft_version": null,
              "opensea_version": null,
              "owner": 259997920,
              "schema_name": "ERC721",
              "symbol": "GHOST",
              "total_supply": "5556",
              "description": "\"We are the ghosts, your nightmare, your daydream.\"\n\nWeirdo Ghost Gang is a collection of 5555 Lil Ghosts roaming in the web3 playground together with countless frens. Lil Ghosts are not just avatars, but also the duke of a castle, the manager of a street brand, a super rapper, a virtual idol, and the protagonist of an adventure...only the boundary of imagination can limit our pace, our existence is not only in the metaverse but will eventually penetrate the reality.\n\n👻[Twitter](https://twitter.com/WeirdoGhostGang) | [Discord](https://discord.com/invite/weirdoghost) | [Website](https://www.weirdoghost.com/) 👻\n\n\n*Tip: You can check whether a Lil Ghost has claimed [WGG Haunted House](https://opensea.io/collection/wgg-haunted-house) in [Gallery](https://www.weirdoghost.com/gallery) by searching the token number.\n",
              "external_link": "https://www.weirdoghost.com/",
              "image_url": "https://i.seadn.io/gcs/files/d19a1520f0791c9b38bb70f6fb5d49cf.gif?w=500&auto=format",
              "default_to_fiat": false,
              "dev_buyer_fee_basis_points": 0,
              "dev_seller_fee_basis_points": 500,
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": 0,
              "opensea_seller_fee_basis_points": 250,
              "buyer_fee_basis_points": 0,
              "seller_fee_basis_points": 750,
              "payout_address": "0x0921d663401d11ce92a8b3b7b559b52bb05291c3"
            },
            "permalink": "https://opensea.io/assets/ethereum/0x9401518f4ebba857baa879d9f76e1cc8b31ed197/5333",
            "collection": {
              "banner_image_url": "https://i.seadn.io/gae/fPVJq87DLFgPwxKnxPGjpeNahYceZT77cNNxXjH0s76mqQKx9i4isUxIIgis6J5lbdRFjIFtrm0ibK1h175LFrCFMW2lVGOifozWg8s?w=500&auto=format",
              "chat_url": null,
              "created_date": "2022-02-09T09:40:30.006181+00:00",
              "default_to_fiat": false,
              "description": "\"We are the ghosts, your nightmare, your daydream.\"\n\nWeirdo Ghost Gang is a collection of 5555 Lil Ghosts roaming in the web3 playground together with countless frens. Lil Ghosts are not just avatars, but also the duke of a castle, the manager of a street brand, a super rapper, a virtual idol, and the protagonist of an adventure...only the boundary of imagination can limit our pace, our existence is not only in the metaverse but will eventually penetrate the reality.\n\n👻[Twitter](https://twitter.com/WeirdoGhostGang) | [Discord](https://discord.com/invite/weirdoghost) | [Website](https://www.weirdoghost.com/) 👻\n\n\n*Tip: You can check whether a Lil Ghost has claimed [WGG Haunted House](https://opensea.io/collection/wgg-haunted-house) in [Gallery](https://www.weirdoghost.com/gallery) by searching the token number.\n",
              "dev_buyer_fee_basis_points": "0",
              "dev_seller_fee_basis_points": "500",
              "discord_url": "https://discord.gg/weirdoghost",
              "display_data": {
                "card_display_style": "cover",
                "images": null
              },
              "external_url": "https://www.weirdoghost.com/",
              "featured": false,
              "featured_image_url": "https://i.seadn.io/gae/FBO6vfVx1DsBER60kO0vL-r7fouTI--iQkFMKLf3E8Vqa5SIksIvFg9_vHGM1iY6IAoKgYrxuwgB8AePWonriD02oi_ri3ZbVRuRi9I?w=500&auto=format",
              "hidden": false,
              "safelist_request_status": "verified",
              "image_url": "https://i.seadn.io/gcs/files/d19a1520f0791c9b38bb70f6fb5d49cf.gif?w=500&auto=format",
              "is_subject_to_whitelist": false,
              "large_image_url": "https://i.seadn.io/gae/FBO6vfVx1DsBER60kO0vL-r7fouTI--iQkFMKLf3E8Vqa5SIksIvFg9_vHGM1iY6IAoKgYrxuwgB8AePWonriD02oi_ri3ZbVRuRi9I?w=500&auto=format",
              "medium_username": null,
              "name": "The Weirdo Ghost Gang",
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": "0",
              "opensea_seller_fee_basis_points": 250,
              "payout_address": "0x0921d663401d11ce92a8b3b7b559b52bb05291c3",
              "require_email": false,
              "short_description": null,
              "slug": "the-weirdo-ghost-gang",
              "telegram_url": null,
              "twitter_username": null,
              "instagram_username": null,
              "wiki_url": null,
              "is_nsfw": false,
              "fees": {
                "seller_fees": {
                  "0x0921d663401d11ce92a8b3b7b559b52bb05291c3": 500
                },
                "opensea_fees": {
                  "0x0000a26b00c1f0df003000390027140000faa719": 250
                }
              },
              "is_rarity_enabled": true,
              "is_creator_fees_enforced": false
            },
            "decimals": 0,
            "token_metadata": "https://ipfs.io/ipfs/QmU61BwmB9fm3kN4EWS14YxrB1FFJcMWj9GRrf4hsEvaYE/5333",
            "is_nsfw": false,
            "owner": null
          }
        ],
        "maker": null,
        "slug": null,
        "name": null,
        "description": null,
        "external_link": null,
        "asset_contract": null,
        "permalink": null,
        "seaport_sell_orders": null
      },
      "taker_asset_bundle": {
        "assets": [
          {
            "id": 13689077,
            "token_id": "0",
            "num_sales": 11,
            "background_color": null,
            "image_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_preview_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_thumbnail_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_original_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "animation_url": null,
            "animation_original_url": null,
            "name": "Ether",
            "description": "",
            "external_link": null,
            "asset_contract": {
              "address": "0x0000000000000000000000000000000000000000",
              "asset_contract_type": "fungible",
              "chain_identifier": "ethereum",
              "created_date": "2019-08-02T23:41:09.503168",
              "name": "Ether",
              "nft_version": null,
              "opensea_version": null,
              "owner": null,
              "schema_name": "ERC20",
              "symbol": "ETH",
              "total_supply": null,
              "description": "",
              "external_link": null,
              "image_url": null,
              "default_to_fiat": false,
              "dev_buyer_fee_basis_points": 0,
              "dev_seller_fee_basis_points": 0,
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": 0,
              "opensea_seller_fee_basis_points": 250,
              "buyer_fee_basis_points": 0,
              "seller_fee_basis_points": 250,
              "payout_address": null
            },
            "permalink": "https://opensea.io/assets/ethereum/0x0000000000000000000000000000000000000000/0",
            "collection": {
              "banner_image_url": null,
              "chat_url": null,
              "created_date": "2022-08-11T13:34:04.673691+00:00",
              "default_to_fiat": false,
              "description": "",
              "dev_buyer_fee_basis_points": "0",
              "dev_seller_fee_basis_points": "0",
              "discord_url": null,
              "display_data": {
                "card_display_style": "contain",
                "images": []
              },
              "external_url": null,
              "featured": false,
              "featured_image_url": null,
              "hidden": false,
              "safelist_request_status": "approved",
              "image_url": null,
              "is_subject_to_whitelist": false,
              "large_image_url": null,
              "medium_username": null,
              "name": "OpenSea PaymentAssets",
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": "0",
              "opensea_seller_fee_basis_points": 250,
              "payout_address": null,
              "require_email": false,
              "short_description": null,
              "slug": "opensea-paymentassets",
              "telegram_url": null,
              "twitter_username": null,
              "instagram_username": null,
              "wiki_url": null,
              "is_nsfw": false,
              "fees": {
                "seller_fees": {},
                "opensea_fees": {
                  "0x0000a26b00c1f0df003000390027140000faa719": 250
                }
              },
              "is_rarity_enabled": false,
              "is_creator_fees_enforced": false
            },
            "decimals": 18,
            "token_metadata": null,
            "is_nsfw": false,
            "owner": null
          }
        ],
        "maker": null,
        "slug": null,
        "name": null,
        "description": null,
        "external_link": null,
        "asset_contract": null,
        "permalink": null,
        "seaport_sell_orders": null
      }
    },
    {
      "created_date": "2023-11-22T02:36:53.899072",
      "closing_date": "2023-11-23T02:36:49",
      "listing_time": 1700620609,
      "expiration_time": 1700707009,
      "order_hash": "0x577798ea26383d7cb7dafb25bf29de65b20342ab5c7b05bb875f834bfb809153",
      "protocol_data": {
        "parameters": {
          "offerer": "0x0097b9cfe64455eed479292671a1121f502bc954",
          "offer": [
            {
              "itemType": 2,
              "token": "0x9401518f4EBBA857BAA879D9f76E1Cc8b31ed197",
              "identifierOrCriteria": "5333",
              "startAmount": "1",
              "endAmount": "1"
            }
          ],
          "consideration": [
            {
              "itemType": 0,
              "token": "0x0000000000000000000000000000000000000000",
              "identifierOrCriteria": "0",
              "startAmount": "487505225000000000",
              "endAmount": "487505225000000000",
              "recipient": "0x0097b9cFE64455EED479292671A1121F502bc954"
            },
            {
              "itemType": 0,
              "token": "0x0000000000000000000000000000000000000000",
              "identifierOrCriteria": "0",
              "startAmount": "2449775000000000",
              "endAmount": "2449775000000000",
              "recipient": "0x0000a26b00c1F0DF003000390027140000fAa719"
            }
          ],
          "startTime": "1700620609",
          "endTime": "1700707009",
          "orderType": 0,
          "zone": "0x004C00500000aD104D7DBd00e3ae0A5C00560C00",
          "zoneHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
          "salt": "0x72db8c0b0000000000000000000000000000000000000000925276caacecf068",
          "conduitKey": "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
          "totalOriginalConsiderationItems": 2,
          "counter": "0x3fbc81a0b0ffd9cf3cef372d93bfc35f5"
        },
        "signature": null
      },
      "protocol_address": "0x00000000000000adc04c56bf30ac9d3c0aaf14dc",
      "current_price": "489955000000000000",
      "maker": {
        "user": 35212416,
        "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/31.png",
        "address": "0x0097b9cfe64455eed479292671a1121f502bc954",
        "config": ""
      },
      "taker": null,
      "maker_fees": [
        {
          "account": {
            "user": null,
            "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/29.png",
            "address": "0x0000a26b00c1f0df003000390027140000faa719",
            "config": ""
          },
          "basis_points": "50"
        }
      ],
      "taker_fees": [],
      "side": "ask",
      "order_type": "basic",
      "cancelled": false,
      "finalized": false,
      "marked_invalid": false,
      "remaining_quantity": 1,
      "relay_id": "T3JkZXJWMlR5cGU6MTQzNDE0NjY4ODU=",
      "criteria_proof": null,
      "maker_asset_bundle": {
        "assets": [
          {
            "id": 284206546,
            "token_id": "5333",
            "num_sales": 35,
            "background_color": null,
            "image_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_preview_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_thumbnail_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_original_url": "ipfs://QmTVdT52JWqEgowWLVSpxeqj1u7ZayeqGvpsjWP5jWhjsY",
            "animation_url": null,
            "animation_original_url": null,
            "name": "Lil Ghost#5333",
            "description": "We are the ghost, your nightmare, your daydream.",
            "external_link": null,
            "asset_contract": {
              "address": "0x9401518f4ebba857baa879d9f76e1cc8b31ed197",
              "asset_contract_type": "non-fungible",
              "chain_identifier": "ethereum",
              "created_date": "2022-02-09T09:20:52.962749",
              "name": "Weirdo Ghost Gang",
              "nft_version": null,
              "opensea_version": null,
              "owner": 259997920,
              "schema_name": "ERC721",
              "symbol": "GHOST",
              "total_supply": "5556",
              "description": "\"We are the ghosts, your nightmare, your daydream.\"\n\nWeirdo Ghost Gang is a collection of 5555 Lil Ghosts roaming in the web3 playground together with countless frens. Lil Ghosts are not just avatars, but also the duke of a castle, the manager of a street brand, a super rapper, a virtual idol, and the protagonist of an adventure...only the boundary of imagination can limit our pace, our existence is not only in the metaverse but will eventually penetrate the reality.\n\n👻[Twitter](https://twitter.com/WeirdoGhostGang) | [Discord](https://discord.com/invite/weirdoghost) | [Website](https://www.weirdoghost.com/) 👻\n\n\n*Tip: You can check whether a Lil Ghost has claimed [WGG Haunted House](https://opensea.io/collection/wgg-haunted-house) in [Gallery](https://www.weirdoghost.com/gallery) by searching the token number.\n",
              "external_link": "https://www.weirdoghost.com/",
              "image_url": "https://i.seadn.io/gcs/files/d19a1520f0791c9b38bb70f6fb5d49cf.gif?w=500&auto=format",
              "default_to_fiat": false,
              "dev_buyer_fee_basis_points": 0,
              "dev_seller_fee_basis_points": 500,
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": 0,
              "opensea_seller_fee_basis_points": 250,
              "buyer_fee_basis_points": 0,
              "seller_fee_basis_points": 750,
              "payout_address": "0x0921d663401d11ce92a8b3b7b559b52bb05291c3"
            },
            "permalink": "https://opensea.io/assets/ethereum/0x9401518f4ebba857baa879d9f76e1cc8b31ed197/5333",
            "collection": {
              "banner_image_url": "https://i.seadn.io/gae/fPVJq87DLFgPwxKnxPGjpeNahYceZT77cNNxXjH0s76mqQKx9i4isUxIIgis6J5lbdRFjIFtrm0ibK1h175LFrCFMW2lVGOifozWg8s?w=500&auto=format",
              "chat_url": null,
              "created_date": "2022-02-09T09:40:30.006181+00:00",
              "default_to_fiat": false,
              "description": "\"We are the ghosts, your nightmare, your daydream.\"\n\nWeirdo Ghost Gang is a collection of 5555 Lil Ghosts roaming in the web3 playground together with countless frens. Lil Ghosts are not just avatars, but also the duke of a castle, the manager of a street brand, a super rapper, a virtual idol, and the protagonist of an adventure...only the boundary of imagination can limit our pace, our existence is not only in the metaverse but will eventually penetrate the reality.\n\n👻[Twitter](https://twitter.com/WeirdoGhostGang) | [Discord](https://discord.com/invite/weirdoghost) | [Website](https://www.weirdoghost.com/) 👻\n\n\n*Tip: You can check whether a Lil Ghost has claimed [WGG Haunted House](https://opensea.io/collection/wgg-haunted-house) in [Gallery](https://www.weirdoghost.com/gallery) by searching the token number.\n",
              "dev_buyer_fee_basis_points": "0",
              "dev_seller_fee_basis_points": "500",
              "discord_url": "https://discord.gg/weirdoghost",
              "display_data": {
                "card_display_style": "cover",
                "images": null
              },
              "external_url": "https://www.weirdoghost.com/",
              "featured": false,
              "featured_image_url": "https://i.seadn.io/gae/FBO6vfVx1DsBER60kO0vL-r7fouTI--iQkFMKLf3E8Vqa5SIksIvFg9_vHGM1iY6IAoKgYrxuwgB8AePWonriD02oi_ri3ZbVRuRi9I?w=500&auto=format",
              "hidden": false,
              "safelist_request_status": "verified",
              "image_url": "https://i.seadn.io/gcs/files/d19a1520f0791c9b38bb70f6fb5d49cf.gif?w=500&auto=format",
              "is_subject_to_whitelist": false,
              "large_image_url": "https://i.seadn.io/gae/FBO6vfVx1DsBER60kO0vL-r7fouTI--iQkFMKLf3E8Vqa5SIksIvFg9_vHGM1iY6IAoKgYrxuwgB8AePWonriD02oi_ri3ZbVRuRi9I?w=500&auto=format",
              "medium_username": null,
              "name": "The Weirdo Ghost Gang",
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": "0",
              "opensea_seller_fee_basis_points": 250,
              "payout_address": "0x0921d663401d11ce92a8b3b7b559b52bb05291c3",
              "require_email": false,
              "short_description": null,
              "slug": "the-weirdo-ghost-gang",
              "telegram_url": null,
              "twitter_username": null,
              "instagram_username": null,
              "wiki_url": null,
              "is_nsfw": false,
              "fees": {
                "seller_fees": {
                  "0x0921d663401d11ce92a8b3b7b559b52bb05291c3": 500
                },
                "opensea_fees": {
                  "0x0000a26b00c1f0df003000390027140000faa719": 250
                }
              },
              "is_rarity_enabled": true,
              "is_creator_fees_enforced": false
            },
            "decimals": 0,
            "token_metadata": "https://ipfs.io/ipfs/QmU61BwmB9fm3kN4EWS14YxrB1FFJcMWj9GRrf4hsEvaYE/5333",
            "is_nsfw": false,
            "owner": null
          }
        ],
        "maker": null,
        "slug": null,
        "name": null,
        "description": null,
        "external_link": null,
        "asset_contract": null,
        "permalink": null,
        "seaport_sell_orders": null
      },
      "taker_asset_bundle": {
        "assets": [
          {
            "id": 13689077,
            "token_id": "0",
            "num_sales": 11,
            "background_color": null,
            "image_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_preview_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_thumbnail_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_original_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "animation_url": null,
            "animation_original_url": null,
            "name": "Ether",
            "description": "",
            "external_link": null,
            "asset_contract": {
              "address": "0x0000000000000000000000000000000000000000",
              "asset_contract_type": "fungible",
              "chain_identifier": "ethereum",
              "created_date": "2019-08-02T23:41:09.503168",
              "name": "Ether",
              "nft_version": null,
              "opensea_version": null,
              "owner": null,
              "schema_name": "ERC20",
              "symbol": "ETH",
              "total_supply": null,
              "description": "",
              "external_link": null,
              "image_url": null,
              "default_to_fiat": false,
              "dev_buyer_fee_basis_points": 0,
              "dev_seller_fee_basis_points": 0,
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": 0,
              "opensea_seller_fee_basis_points": 250,
              "buyer_fee_basis_points": 0,
              "seller_fee_basis_points": 250,
              "payout_address": null
            },
            "permalink": "https://opensea.io/assets/ethereum/0x0000000000000000000000000000000000000000/0",
            "collection": {
              "banner_image_url": null,
              "chat_url": null,
              "created_date": "2022-08-11T13:34:04.673691+00:00",
              "default_to_fiat": false,
              "description": "",
              "dev_buyer_fee_basis_points": "0",
              "dev_seller_fee_basis_points": "0",
              "discord_url": null,
              "display_data": {
                "card_display_style": "contain",
                "images": []
              },
              "external_url": null,
              "featured": false,
              "featured_image_url": null,
              "hidden": false,
              "safelist_request_status": "approved",
              "image_url": null,
              "is_subject_to_whitelist": false,
              "large_image_url": null,
              "medium_username": null,
              "name": "OpenSea PaymentAssets",
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": "0",
              "opensea_seller_fee_basis_points": 250,
              "payout_address": null,
              "require_email": false,
              "short_description": null,
              "slug": "opensea-paymentassets",
              "telegram_url": null,
              "twitter_username": null,
              "instagram_username": null,
              "wiki_url": null,
              "is_nsfw": false,
              "fees": {
                "seller_fees": {},
                "opensea_fees": {
                  "0x0000a26b00c1f0df003000390027140000faa719": 250
                }
              },
              "is_rarity_enabled": false,
              "is_creator_fees_enforced": false
            },
            "decimals": 18,
            "token_metadata": null,
            "is_nsfw": false,
            "owner": null
          }
        ],
        "maker": null,
        "slug": null,
        "name": null,
        "description": null,
        "external_link": null,
        "asset_contract": null,
        "permalink": null,
        "seaport_sell_orders": null
      }
    },
    {
      "created_date": "2023-11-21T16:00:41.616704",
      "closing_date": "2023-11-22T16:00:32",
      "listing_time": 1700582432,
      "expiration_time": 1700668832,
      "order_hash": "0x6c3734de11682bd58043050c42b346add584b19de35d4bbb6e52b7aa62586c88",
      "protocol_data": {
        "parameters": {
          "offerer": "0x0097b9cfe64455eed479292671a1121f502bc954",
          "offer": [
            {
              "itemType": 2,
              "token": "0x9401518f4EBBA857BAA879D9f76E1Cc8b31ed197",
              "identifierOrCriteria": "5333",
              "startAmount": "1",
              "endAmount": "1"
            }
          ],
          "consideration": [
            {
              "itemType": 0,
              "token": "0x0000000000000000000000000000000000000000",
              "identifierOrCriteria": "0",
              "startAmount": "487505225000000000",
              "endAmount": "487505225000000000",
              "recipient": "0x0097b9cFE64455EED479292671A1121F502bc954"
            },
            {
              "itemType": 0,
              "token": "0x0000000000000000000000000000000000000000",
              "identifierOrCriteria": "0",
              "startAmount": "2449775000000000",
              "endAmount": "2449775000000000",
              "recipient": "0x0000a26b00c1F0DF003000390027140000fAa719"
            }
          ],
          "startTime": "1700582432",
          "endTime": "1700668832",
          "orderType": 0,
          "zone": "0x004C00500000aD104D7DBd00e3ae0A5C00560C00",
          "zoneHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
          "salt": "0x72db8c0b0000000000000000000000000000000000000000af5f337e40daff38",
          "conduitKey": "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
          "totalOriginalConsiderationItems": 2,
          "counter": "0x3fbc81a0b0ffd9cf3cef372d93bfc35f5"
        },
        "signature": null
      },
      "protocol_address": "0x00000000000000adc04c56bf30ac9d3c0aaf14dc",
      "current_price": "489955000000000000",
      "maker": {
        "user": 35212416,
        "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/31.png",
        "address": "0x0097b9cfe64455eed479292671a1121f502bc954",
        "config": ""
      },
      "taker": null,
      "maker_fees": [
        {
          "account": {
            "user": null,
            "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/29.png",
            "address": "0x0000a26b00c1f0df003000390027140000faa719",
            "config": ""
          },
          "basis_points": "50"
        }
      ],
      "taker_fees": [],
      "side": "ask",
      "order_type": "basic",
      "cancelled": false,
      "finalized": false,
      "marked_invalid": false,
      "remaining_quantity": 1,
      "relay_id": "T3JkZXJWMlR5cGU6MTQzMjgzOTU5ODA=",
      "criteria_proof": null,
      "maker_asset_bundle": {
        "assets": [
          {
            "id": 284206546,
            "token_id": "5333",
            "num_sales": 35,
            "background_color": null,
            "image_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_preview_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_thumbnail_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_original_url": "ipfs://QmTVdT52JWqEgowWLVSpxeqj1u7ZayeqGvpsjWP5jWhjsY",
            "animation_url": null,
            "animation_original_url": null,
            "name": "Lil Ghost#5333",
            "description": "We are the ghost, your nightmare, your daydream.",
            "external_link": null,
            "asset_contract": {
              "address": "0x9401518f4ebba857baa879d9f76e1cc8b31ed197",
              "asset_contract_type": "non-fungible",
              "chain_identifier": "ethereum",
              "created_date": "2022-02-09T09:20:52.962749",
              "name": "Weirdo Ghost Gang",
              "nft_version": null,
              "opensea_version": null,
              "owner": 259997920,
              "schema_name": "ERC721",
              "symbol": "GHOST",
              "total_supply": "5556",
              "description": "\"We are the ghosts, your nightmare, your daydream.\"\n\nWeirdo Ghost Gang is a collection of 5555 Lil Ghosts roaming in the web3 playground together with countless frens. Lil Ghosts are not just avatars, but also the duke of a castle, the manager of a street brand, a super rapper, a virtual idol, and the protagonist of an adventure...only the boundary of imagination can limit our pace, our existence is not only in the metaverse but will eventually penetrate the reality.\n\n👻[Twitter](https://twitter.com/WeirdoGhostGang) | [Discord](https://discord.com/invite/weirdoghost) | [Website](https://www.weirdoghost.com/) 👻\n\n\n*Tip: You can check whether a Lil Ghost has claimed [WGG Haunted House](https://opensea.io/collection/wgg-haunted-house) in [Gallery](https://www.weirdoghost.com/gallery) by searching the token number.\n",
              "external_link": "https://www.weirdoghost.com/",
              "image_url": "https://i.seadn.io/gcs/files/d19a1520f0791c9b38bb70f6fb5d49cf.gif?w=500&auto=format",
              "default_to_fiat": false,
              "dev_buyer_fee_basis_points": 0,
              "dev_seller_fee_basis_points": 500,
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": 0,
              "opensea_seller_fee_basis_points": 250,
              "buyer_fee_basis_points": 0,
              "seller_fee_basis_points": 750,
              "payout_address": "0x0921d663401d11ce92a8b3b7b559b52bb05291c3"
            },
            "permalink": "https://opensea.io/assets/ethereum/0x9401518f4ebba857baa879d9f76e1cc8b31ed197/5333",
            "collection": {
              "banner_image_url": "https://i.seadn.io/gae/fPVJq87DLFgPwxKnxPGjpeNahYceZT77cNNxXjH0s76mqQKx9i4isUxIIgis6J5lbdRFjIFtrm0ibK1h175LFrCFMW2lVGOifozWg8s?w=500&auto=format",
              "chat_url": null,
              "created_date": "2022-02-09T09:40:30.006181+00:00",
              "default_to_fiat": false,
              "description": "\"We are the ghosts, your nightmare, your daydream.\"\n\nWeirdo Ghost Gang is a collection of 5555 Lil Ghosts roaming in the web3 playground together with countless frens. Lil Ghosts are not just avatars, but also the duke of a castle, the manager of a street brand, a super rapper, a virtual idol, and the protagonist of an adventure...only the boundary of imagination can limit our pace, our existence is not only in the metaverse but will eventually penetrate the reality.\n\n👻[Twitter](https://twitter.com/WeirdoGhostGang) | [Discord](https://discord.com/invite/weirdoghost) | [Website](https://www.weirdoghost.com/) 👻\n\n\n*Tip: You can check whether a Lil Ghost has claimed [WGG Haunted House](https://opensea.io/collection/wgg-haunted-house) in [Gallery](https://www.weirdoghost.com/gallery) by searching the token number.\n",
              "dev_buyer_fee_basis_points": "0",
              "dev_seller_fee_basis_points": "500",
              "discord_url": "https://discord.gg/weirdoghost",
              "display_data": {
                "card_display_style": "cover",
                "images": null
              },
              "external_url": "https://www.weirdoghost.com/",
              "featured": false,
              "featured_image_url": "https://i.seadn.io/gae/FBO6vfVx1DsBER60kO0vL-r7fouTI--iQkFMKLf3E8Vqa5SIksIvFg9_vHGM1iY6IAoKgYrxuwgB8AePWonriD02oi_ri3ZbVRuRi9I?w=500&auto=format",
              "hidden": false,
              "safelist_request_status": "verified",
              "image_url": "https://i.seadn.io/gcs/files/d19a1520f0791c9b38bb70f6fb5d49cf.gif?w=500&auto=format",
              "is_subject_to_whitelist": false,
              "large_image_url": "https://i.seadn.io/gae/FBO6vfVx1DsBER60kO0vL-r7fouTI--iQkFMKLf3E8Vqa5SIksIvFg9_vHGM1iY6IAoKgYrxuwgB8AePWonriD02oi_ri3ZbVRuRi9I?w=500&auto=format",
              "medium_username": null,
              "name": "The Weirdo Ghost Gang",
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": "0",
              "opensea_seller_fee_basis_points": 250,
              "payout_address": "0x0921d663401d11ce92a8b3b7b559b52bb05291c3",
              "require_email": false,
              "short_description": null,
              "slug": "the-weirdo-ghost-gang",
              "telegram_url": null,
              "twitter_username": null,
              "instagram_username": null,
              "wiki_url": null,
              "is_nsfw": false,
              "fees": {
                "seller_fees": {
                  "0x0921d663401d11ce92a8b3b7b559b52bb05291c3": 500
                },
                "opensea_fees": {
                  "0x0000a26b00c1f0df003000390027140000faa719": 250
                }
              },
              "is_rarity_enabled": true,
              "is_creator_fees_enforced": false
            },
            "decimals": 0,
            "token_metadata": "https://ipfs.io/ipfs/QmU61BwmB9fm3kN4EWS14YxrB1FFJcMWj9GRrf4hsEvaYE/5333",
            "is_nsfw": false,
            "owner": null
          }
        ],
        "maker": null,
        "slug": null,
        "name": null,
        "description": null,
        "external_link": null,
        "asset_contract": null,
        "permalink": null,
        "seaport_sell_orders": null
      },
      "taker_asset_bundle": {
        "assets": [
          {
            "id": 13689077,
            "token_id": "0",
            "num_sales": 11,
            "background_color": null,
            "image_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_preview_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_thumbnail_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_original_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "animation_url": null,
            "animation_original_url": null,
            "name": "Ether",
            "description": "",
            "external_link": null,
            "asset_contract": {
              "address": "0x0000000000000000000000000000000000000000",
              "asset_contract_type": "fungible",
              "chain_identifier": "ethereum",
              "created_date": "2019-08-02T23:41:09.503168",
              "name": "Ether",
              "nft_version": null,
              "opensea_version": null,
              "owner": null,
              "schema_name": "ERC20",
              "symbol": "ETH",
              "total_supply": null,
              "description": "",
              "external_link": null,
              "image_url": null,
              "default_to_fiat": false,
              "dev_buyer_fee_basis_points": 0,
              "dev_seller_fee_basis_points": 0,
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": 0,
              "opensea_seller_fee_basis_points": 250,
              "buyer_fee_basis_points": 0,
              "seller_fee_basis_points": 250,
              "payout_address": null
            },
            "permalink": "https://opensea.io/assets/ethereum/0x0000000000000000000000000000000000000000/0",
            "collection": {
              "banner_image_url": null,
              "chat_url": null,
              "created_date": "2022-08-11T13:34:04.673691+00:00",
              "default_to_fiat": false,
              "description": "",
              "dev_buyer_fee_basis_points": "0",
              "dev_seller_fee_basis_points": "0",
              "discord_url": null,
              "display_data": {
                "card_display_style": "contain",
                "images": []
              },
              "external_url": null,
              "featured": false,
              "featured_image_url": null,
              "hidden": false,
              "safelist_request_status": "approved",
              "image_url": null,
              "is_subject_to_whitelist": false,
              "large_image_url": null,
              "medium_username": null,
              "name": "OpenSea PaymentAssets",
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": "0",
              "opensea_seller_fee_basis_points": 250,
              "payout_address": null,
              "require_email": false,
              "short_description": null,
              "slug": "opensea-paymentassets",
              "telegram_url": null,
              "twitter_username": null,
              "instagram_username": null,
              "wiki_url": null,
              "is_nsfw": false,
              "fees": {
                "seller_fees": {},
                "opensea_fees": {
                  "0x0000a26b00c1f0df003000390027140000faa719": 250
                }
              },
              "is_rarity_enabled": false,
              "is_creator_fees_enforced": false
            },
            "decimals": 18,
            "token_metadata": null,
            "is_nsfw": false,
            "owner": null
          }
        ],
        "maker": null,
        "slug": null,
        "name": null,
        "description": null,
        "external_link": null,
        "asset_contract": null,
        "permalink": null,
        "seaport_sell_orders": null
      }
    },
    {
      "created_date": "2023-11-21T14:46:27.918946",
      "closing_date": "2023-11-22T14:46:24",
      "listing_time": 1700577984,
      "expiration_time": 1700664384,
      "order_hash": "0x4b59204b1df9d50608b57682d43da320c087f19b60a9671086f4d8088dbf9bdd",
      "protocol_data": {
        "parameters": {
          "offerer": "0x0097b9cfe64455eed479292671a1121f502bc954",
          "offer": [
            {
              "itemType": 2,
              "token": "0x9401518f4EBBA857BAA879D9f76E1Cc8b31ed197",
              "identifierOrCriteria": "5333",
              "startAmount": "1",
              "endAmount": "1"
            }
          ],
          "consideration": [
            {
              "itemType": 0,
              "token": "0x0000000000000000000000000000000000000000",
              "identifierOrCriteria": "0",
              "startAmount": "547150500000000000",
              "endAmount": "547150500000000000",
              "recipient": "0x0097b9cFE64455EED479292671A1121F502bc954"
            },
            {
              "itemType": 0,
              "token": "0x0000000000000000000000000000000000000000",
              "identifierOrCriteria": "0",
              "startAmount": "2749500000000000",
              "endAmount": "2749500000000000",
              "recipient": "0x0000a26b00c1F0DF003000390027140000fAa719"
            }
          ],
          "startTime": "1700577984",
          "endTime": "1700664384",
          "orderType": 0,
          "zone": "0x004C00500000aD104D7DBd00e3ae0A5C00560C00",
          "zoneHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
          "salt": "0x72db8c0b000000000000000000000000000000000000000078d775837a82ee70",
          "conduitKey": "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
          "totalOriginalConsiderationItems": 2,
          "counter": "0x3fbc81a0b0ffd9cf3cef372d93bfc35f5"
        },
        "signature": null
      },
      "protocol_address": "0x00000000000000adc04c56bf30ac9d3c0aaf14dc",
      "current_price": "549900000000000000",
      "maker": {
        "user": 35212416,
        "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/31.png",
        "address": "0x0097b9cfe64455eed479292671a1121f502bc954",
        "config": ""
      },
      "taker": null,
      "maker_fees": [
        {
          "account": {
            "user": null,
            "profile_img_url": "https://storage.googleapis.com/opensea-static/opensea-profile/29.png",
            "address": "0x0000a26b00c1f0df003000390027140000faa719",
            "config": ""
          },
          "basis_points": "50"
        }
      ],
      "taker_fees": [],
      "side": "ask",
      "order_type": "basic",
      "cancelled": false,
      "finalized": false,
      "marked_invalid": false,
      "remaining_quantity": 1,
      "relay_id": "T3JkZXJWMlR5cGU6MTQzMjY5MDU2MzU=",
      "criteria_proof": null,
      "maker_asset_bundle": {
        "assets": [
          {
            "id": 284206546,
            "token_id": "5333",
            "num_sales": 35,
            "background_color": null,
            "image_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_preview_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_thumbnail_url": "https://i.seadn.io/gcs/files/2e36bf41ffa835fa6b00a08b22530cd0.png?w=500&auto=format",
            "image_original_url": "ipfs://QmTVdT52JWqEgowWLVSpxeqj1u7ZayeqGvpsjWP5jWhjsY",
            "animation_url": null,
            "animation_original_url": null,
            "name": "Lil Ghost#5333",
            "description": "We are the ghost, your nightmare, your daydream.",
            "external_link": null,
            "asset_contract": {
              "address": "0x9401518f4ebba857baa879d9f76e1cc8b31ed197",
              "asset_contract_type": "non-fungible",
              "chain_identifier": "ethereum",
              "created_date": "2022-02-09T09:20:52.962749",
              "name": "Weirdo Ghost Gang",
              "nft_version": null,
              "opensea_version": null,
              "owner": 259997920,
              "schema_name": "ERC721",
              "symbol": "GHOST",
              "total_supply": "5556",
              "description": "\"We are the ghosts, your nightmare, your daydream.\"\n\nWeirdo Ghost Gang is a collection of 5555 Lil Ghosts roaming in the web3 playground together with countless frens. Lil Ghosts are not just avatars, but also the duke of a castle, the manager of a street brand, a super rapper, a virtual idol, and the protagonist of an adventure...only the boundary of imagination can limit our pace, our existence is not only in the metaverse but will eventually penetrate the reality.\n\n👻[Twitter](https://twitter.com/WeirdoGhostGang) | [Discord](https://discord.com/invite/weirdoghost) | [Website](https://www.weirdoghost.com/) 👻\n\n\n*Tip: You can check whether a Lil Ghost has claimed [WGG Haunted House](https://opensea.io/collection/wgg-haunted-house) in [Gallery](https://www.weirdoghost.com/gallery) by searching the token number.\n",
              "external_link": "https://www.weirdoghost.com/",
              "image_url": "https://i.seadn.io/gcs/files/d19a1520f0791c9b38bb70f6fb5d49cf.gif?w=500&auto=format",
              "default_to_fiat": false,
              "dev_buyer_fee_basis_points": 0,
              "dev_seller_fee_basis_points": 500,
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": 0,
              "opensea_seller_fee_basis_points": 250,
              "buyer_fee_basis_points": 0,
              "seller_fee_basis_points": 750,
              "payout_address": "0x0921d663401d11ce92a8b3b7b559b52bb05291c3"
            },
            "permalink": "https://opensea.io/assets/ethereum/0x9401518f4ebba857baa879d9f76e1cc8b31ed197/5333",
            "collection": {
              "banner_image_url": "https://i.seadn.io/gae/fPVJq87DLFgPwxKnxPGjpeNahYceZT77cNNxXjH0s76mqQKx9i4isUxIIgis6J5lbdRFjIFtrm0ibK1h175LFrCFMW2lVGOifozWg8s?w=500&auto=format",
              "chat_url": null,
              "created_date": "2022-02-09T09:40:30.006181+00:00",
              "default_to_fiat": false,
              "description": "\"We are the ghosts, your nightmare, your daydream.\"\n\nWeirdo Ghost Gang is a collection of 5555 Lil Ghosts roaming in the web3 playground together with countless frens. Lil Ghosts are not just avatars, but also the duke of a castle, the manager of a street brand, a super rapper, a virtual idol, and the protagonist of an adventure...only the boundary of imagination can limit our pace, our existence is not only in the metaverse but will eventually penetrate the reality.\n\n👻[Twitter](https://twitter.com/WeirdoGhostGang) | [Discord](https://discord.com/invite/weirdoghost) | [Website](https://www.weirdoghost.com/) 👻\n\n\n*Tip: You can check whether a Lil Ghost has claimed [WGG Haunted House](https://opensea.io/collection/wgg-haunted-house) in [Gallery](https://www.weirdoghost.com/gallery) by searching the token number.\n",
              "dev_buyer_fee_basis_points": "0",
              "dev_seller_fee_basis_points": "500",
              "discord_url": "https://discord.gg/weirdoghost",
              "display_data": {
                "card_display_style": "cover",
                "images": null
              },
              "external_url": "https://www.weirdoghost.com/",
              "featured": false,
              "featured_image_url": "https://i.seadn.io/gae/FBO6vfVx1DsBER60kO0vL-r7fouTI--iQkFMKLf3E8Vqa5SIksIvFg9_vHGM1iY6IAoKgYrxuwgB8AePWonriD02oi_ri3ZbVRuRi9I?w=500&auto=format",
              "hidden": false,
              "safelist_request_status": "verified",
              "image_url": "https://i.seadn.io/gcs/files/d19a1520f0791c9b38bb70f6fb5d49cf.gif?w=500&auto=format",
              "is_subject_to_whitelist": false,
              "large_image_url": "https://i.seadn.io/gae/FBO6vfVx1DsBER60kO0vL-r7fouTI--iQkFMKLf3E8Vqa5SIksIvFg9_vHGM1iY6IAoKgYrxuwgB8AePWonriD02oi_ri3ZbVRuRi9I?w=500&auto=format",
              "medium_username": null,
              "name": "The Weirdo Ghost Gang",
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": "0",
              "opensea_seller_fee_basis_points": 250,
              "payout_address": "0x0921d663401d11ce92a8b3b7b559b52bb05291c3",
              "require_email": false,
              "short_description": null,
              "slug": "the-weirdo-ghost-gang",
              "telegram_url": null,
              "twitter_username": null,
              "instagram_username": null,
              "wiki_url": null,
              "is_nsfw": false,
              "fees": {
                "seller_fees": {
                  "0x0921d663401d11ce92a8b3b7b559b52bb05291c3": 500
                },
                "opensea_fees": {
                  "0x0000a26b00c1f0df003000390027140000faa719": 250
                }
              },
              "is_rarity_enabled": true,
              "is_creator_fees_enforced": false
            },
            "decimals": 0,
            "token_metadata": "https://ipfs.io/ipfs/QmU61BwmB9fm3kN4EWS14YxrB1FFJcMWj9GRrf4hsEvaYE/5333",
            "is_nsfw": false,
            "owner": null
          }
        ],
        "maker": null,
        "slug": null,
        "name": null,
        "description": null,
        "external_link": null,
        "asset_contract": null,
        "permalink": null,
        "seaport_sell_orders": null
      },
      "taker_asset_bundle": {
        "assets": [
          {
            "id": 13689077,
            "token_id": "0",
            "num_sales": 11,
            "background_color": null,
            "image_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_preview_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_thumbnail_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "image_original_url": "https://openseauserdata.com/files/6f8e2979d428180222796ff4a33ab929.svg",
            "animation_url": null,
            "animation_original_url": null,
            "name": "Ether",
            "description": "",
            "external_link": null,
            "asset_contract": {
              "address": "0x0000000000000000000000000000000000000000",
              "asset_contract_type": "fungible",
              "chain_identifier": "ethereum",
              "created_date": "2019-08-02T23:41:09.503168",
              "name": "Ether",
              "nft_version": null,
              "opensea_version": null,
              "owner": null,
              "schema_name": "ERC20",
              "symbol": "ETH",
              "total_supply": null,
              "description": "",
              "external_link": null,
              "image_url": null,
              "default_to_fiat": false,
              "dev_buyer_fee_basis_points": 0,
              "dev_seller_fee_basis_points": 0,
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": 0,
              "opensea_seller_fee_basis_points": 250,
              "buyer_fee_basis_points": 0,
              "seller_fee_basis_points": 250,
              "payout_address": null
            },
            "permalink": "https://opensea.io/assets/ethereum/0x0000000000000000000000000000000000000000/0",
            "collection": {
              "banner_image_url": null,
              "chat_url": null,
              "created_date": "2022-08-11T13:34:04.673691+00:00",
              "default_to_fiat": false,
              "description": "",
              "dev_buyer_fee_basis_points": "0",
              "dev_seller_fee_basis_points": "0",
              "discord_url": null,
              "display_data": {
                "card_display_style": "contain",
                "images": []
              },
              "external_url": null,
              "featured": false,
              "featured_image_url": null,
              "hidden": false,
              "safelist_request_status": "approved",
              "image_url": null,
              "is_subject_to_whitelist": false,
              "large_image_url": null,
              "medium_username": null,
              "name": "OpenSea PaymentAssets",
              "only_proxied_transfers": false,
              "opensea_buyer_fee_basis_points": "0",
              "opensea_seller_fee_basis_points": 250,
              "payout_address": null,
              "require_email": false,
              "short_description": null,
              "slug": "opensea-paymentassets",
              "telegram_url": null,
              "twitter_username": null,
              "instagram_username": null,
              "wiki_url": null,
              "is_nsfw": false,
              "fees": {
                "seller_fees": {},
                "opensea_fees": {
                  "0x0000a26b00c1f0df003000390027140000faa719": 250
                }
              },
              "is_rarity_enabled": false,
              "is_creator_fees_enforced": false
            },
            "decimals": 18,
            "token_metadata": null,
            "is_nsfw": false,
            "owner": null
          }
        ],
        "maker": null,
        "slug": null,
        "name": null,
        "description": null,
        "external_link": null,
        "asset_contract": null,
        "permalink": null,
        "seaport_sell_orders": null
      }
    }
  ]
}`
	var resp OrdersResponse
	err := json.Unmarshal([]byte(data), &resp)
	require.NoError(t, err)
}
