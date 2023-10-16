package openseaapiutils

import (
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/xTransact/openseaapi/chain"
	"github.com/xTransact/openseaapi/openseaconsts"
)

func GetBaseURL(testnets bool) string {
	if testnets {
		return openseaconsts.BaseUrlTest
	} else {
		return openseaconsts.BaseUrlProd
	}
}

func GetBaseURLByChain(c chain.Chain) string {
	if c.IsTestNet() {
		return openseaconsts.BaseUrlTest
	} else {
		return openseaconsts.BaseUrlProd
	}
}

func GetMethod(contractMethodDef string) string {
	i := strings.Index(contractMethodDef, "(")
	if i > 0 {
		return contractMethodDef[:i]
	}
	return contractMethodDef
}

func ValidateAddress(addr common.Address) bool {
	return common.IsHexAddress(addr.String())
}

func ValidateAmount(key string, amount any) error {
	if amount == nil {
		return fmt.Errorf("%s must not be nil", key)
	}
	amountStr, ok := amount.(string)
	if !ok {
		return fmt.Errorf("invalid %s", key)
	}
	_, ok = new(big.Int).SetString(amountStr, 10)
	if !ok {
		return fmt.Errorf("invalid %s", key)
	}
	return nil
}

func ParseFailureResponse(resp *http.Response, respBody []byte) error {
	if len(respBody) != 0 {
		return fmt.Errorf("%s: %s", resp.Status, respBody)
	}

	return fmt.Errorf("unexpected http response: %s", resp.Status)
}

func StringPtr(v string) *string {
	return &v
}

func PtrToString(p *string) (v string) {
	if p == nil {
		return v
	}

	return *p
}

func PtrToInt(p *int) (v int) {
	if p == nil {
		return v
	}

	return *p
}

func PtrToBool(p *bool) (v bool) {
	if p == nil {
		return v
	}

	return *p
}
