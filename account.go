package openseaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/xTransact/errx/v3"

	"github.com/xTransact/openseaapi/openseaapiutils"
	"github.com/xTransact/openseaapi/openseamodels"
)

// GetAccount gets an OpenSea Account Profile including details such as bio, social media usernames, and profile image.
// DOC: https://docs.opensea.io/reference/get_account
func (c *client) GetAccount(ctx context.Context, address common.Address,
	opts ...RequestOptionFn) (resp *openseamodels.Account, err error) {

	if !common.IsHexAddress(address.String()) {
		return nil, errx.New("invalid address")
	}

	o := new(requestOptions)
	for _, apply := range opts {
		apply(o)
	}

	// GET /api/v2/accounts/{address}
	url := fmt.Sprintf("%s/api/v2/accounts/%s", openseaapiutils.GetBaseURL(o.testnets), address.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	c.acceptJson(req)
	body, err := c.doRequest(req, o.testnets)
	if err != nil {
		return nil, errx.WithStack(err)
	}

	resp = new(openseamodels.Account)
	if err = json.Unmarshal(body, resp); err != nil {
		return nil, errx.Wrap(err, "unmarshal response body")
	}

	return resp, nil
}
