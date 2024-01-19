package openseaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

func TestGetListings(t *testing.T) {
	cli := NewClient(WithApiKey(getTestApiKey()))
	listings, err := cli.GetListings(context.Background(), chain.Ethereum, &openseamodels.OrderPayload{
		AssetContractAddress: openseaapiutils.StringPtr("0xed5af388653567af2f388e6224dc7c4b3241c544"), // azuki
		TokenIDs: []json.Number{
			"4332", "7626",
		},
	})
	require.NoError(t, err)
	fmt.Println(string(listings.ToJson()))
}

func TestGetTestListings(t *testing.T) {
	cli := NewClient()
	listings, err := cli.GetListings(context.Background(), chain.Sepolia, &openseamodels.OrderPayload{
		AssetContractAddress: openseaapiutils.StringPtr("0xb31d6b5516eed64a874e9f7ab605e359e20b645f"),
		TokenIDs: []json.Number{
			"9",
		},
	})

	require.NoError(t, err)

	data, err := json.MarshalIndent(listings, "", "  ")
	require.NoError(t, err)

	fmt.Println(string(data))
}

func TestFulfillListings(t *testing.T) {
	cli := NewClient()
	// token_id: 10
	resp, err := cli.FulfillListing(context.Background(), chain.Sepolia,
		"0x39dc941236628a4ddc793a0c49f8065195b52908a98b4e337b6a51ae63965fd2", "0xeFe15c06BAE6bA30b444e6fCD6B94354057fC998")
	require.NoError(t, err)

	data, err := json.MarshalIndent(resp, "", "  ")
	require.NoError(t, err)

	fmt.Println(string(data))
	time.Sleep(time.Second)
	fmt.Println()

	// 0.001 WETH: 0x58eae5fa7fe08200c7f0d71cc16fc7cfdffd0cc6c330245eeedf5b789019c85d
	// resp, err := cli.FulfillListing(context.Background(), chain.Sepolia,
	// 	"0x58eae5fa7fe08200c7f0d71cc16fc7cfdffd0cc6c330245eeedf5b789019c85d", "0x69493301a10A06679a6771D33E8CDd3a5fdA4dB4")
	// require.NoError(t, err)
	//
	// data, err := json.MarshalIndent(resp, "", "  ")
	// require.NoError(t, err)
	//
	// fmt.Println(string(data))
	// time.Sleep(time.Second)
	// fmt.Println()

	// 0.026  ETH: 0xa57250013724312244772508096cfb4a9f7989985729ec2df6e4e75e29687d86
	// resp, err = cli.FulfillListing(context.Background(), chain.Sepolia,
	// 	"0xa57250013724312244772508096cfb4a9f7989985729ec2df6e4e75e29687d86", "0x69493301a10A06679a6771D33E8CDd3a5fdA4dB4")
	// require.NoError(t, err)
	//
	// data, err = json.MarshalIndent(resp, "", "  ")
	// require.NoError(t, err)
	//
	// fmt.Println(string(data))
}

//	type ReceivedItem struct {
//		ItemType   uint8
//		Token      common.Address
//		Identifier *big.Int
//		Amount     *big.Int
//		Recipient  common.Address
//	}
func TestUnmarshalBigInt(t *testing.T) {
	const s = `{"itemType":1,"token":"0xeFe15c06BAE6bA30b444e6fCD6B94354057fC998","identifier":1,"amount":"100","recipient":"0x69493301a10A06679a6771D33E8CDd3a5fdA4dB4"}`
	var item openseamodels.ReceivedItem
	err := json.Unmarshal([]byte(s), &item)
	require.NoError(t, err)
	fmt.Printf("%#v\n", item)
}

func TestUnmarshalFulfillmentData(t *testing.T) {
	const s = `{
    "protocol": "seaport1.5",
    "fulfillment_data":
    {
        "transaction":
        {
            "function": "fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes))",
            "chain": 5,
            "to": "0x00000000000000adc04c56bf30ac9d3c0aaf14dc",
            "value": 1000000000000000,
            "input_data":
            {
                "parameters":
                {
                    "additionalRecipients":
                    [
                        {
                            "amount": "25000000000000",
                            "recipient": "0x0000a26b00c1f0df003000390027140000faa719"
                        }
                    ],
                    "basicOrderType": 0,
                    "considerationAmount": "975000000000000",
                    "considerationIdentifier": "0",
                    "considerationToken": "0x0000000000000000000000000000000000000000",
                    "endTime": "1699581869",
                    "fulfillerConduitKey": "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
                    "offerAmount": "1",
                    "offerIdentifier": "6",
                    "offerToken": "0xb31d6b5516eed64a874e9f7ab605e359e20b645f",
                    "offerer": "0xefe15c06bae6ba30b444e6fcd6b94354057fc998",
                    "offererConduitKey": "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
                    "salt": "24446860302761739304752683030156737591518664810215442929800774530667140607197",
                    "signature": "0x59b6d8aa7c897361745cc43f25611fa65af0237bb09c5fc6d03f9a4e1248afd4fd45726622a15daf56833aec9830404f2487f5fb2cd92a8d332b7f234xxxxxxx",
                    "startTime": "1696903469",
                    "totalOriginalAdditionalRecipients": "1",
                    "zone": "0x004c00500000ad104d7dbd00e3ae0a5c00560c00",
                    "zoneHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
                }
            }
        },
        "orders":
        [
            {
                "parameters":
                {
                    "offerer": "0xefe15c06bae6ba30b444e6fcd6b94354057fc998",
                    "offer":
                    [
                        {
                            "itemType": 2,
                            "token": "0xB31D6B5516Eed64a874E9F7aB605e359e20B645F",
                            "identifierOrCriteria": "6",
                            "startAmount": "1",
                            "endAmount": "1"
                        }
                    ],
                    "consideration":
                    [
                        {
                            "itemType": 0,
                            "token": "0x0000000000000000000000000000000000000000",
                            "identifierOrCriteria": "0",
                            "startAmount": "975000000000000",
                            "endAmount": "975000000000000",
                            "recipient": "0xeFe15c06BAE6bA30b444e6fCD6B94354057fC998"
                        },
                        {
                            "itemType": 0,
                            "token": "0x0000000000000000000000000000000000000000",
                            "identifierOrCriteria": "0",
                            "startAmount": "25000000000000",
                            "endAmount": "25000000000000",
                            "recipient": "0x0000a26b00c1F0DF003000390027140000fAa719"
                        }
                    ],
                    "startTime": "1696903469",
                    "endTime": "1699581869",
                    "orderType": 0,
                    "zone": "0x004C00500000aD104D7DBd00e3ae0A5C00560C00",
                    "zoneHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
                    "salt": "0x360c6ebe000000000000000000000000000000000000000005701a6f0fxxxxxx",
                    "conduitKey": "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
                    "totalOriginalConsiderationItems": 2,
                    "counter": 0
                },
                "signature": "0x59b6d8aa7c897361745cc43f25611fa65af0237bb09c5fc6d03f9a4e1248afd4fd45726622a15daf56833aec9830404f2487f5fb2cd92a8d332b7f234xxxxxxx"
            }
        ]
    }
}`
	var resp openseamodels.FulfillmentDataResponse
	err := json.Unmarshal([]byte(s), &resp)
	require.NoError(t, err)

	p, err := resp.FulfillmentData.Transaction.ParseInputDataToBasicOrder()
	require.NoError(t, err)

	pd, err := json.MarshalIndent(&p, "", "  ")
	require.NoError(t, err)
	fmt.Println(string(pd))
}

func getTestApiKey() string {
	return os.Getenv("OPENSEA_API_KEY")
}
