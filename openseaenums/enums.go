package openseaenums

type ItemType int

const (
	ItemTypeNative ItemType = iota
	ItemTypeERC20
	ItemTypeERC721
	ItemTypeERC1155
	ItemTypeERC721WithCriteria
	ItemTypeERC1155WithCriteria
)

type OrderType int

const (
	// OrderTypeFullOpen : No partial fills, anyone can execute
	OrderTypeFullOpen OrderType = iota
	// OrderTypePartialOpen : Partial fills supported, anyone can execute
	OrderTypePartialOpen
	// OrderTypeFullRestricted : No partial fills, only offerer or zone can check if it can be executed
	OrderTypeFullRestricted
	// OrderTypePartialRestricted : Partial fills supported, only offerer or zone can check if it can be executed
	OrderTypePartialRestricted
	// OrderTypeContract : Contract order type, for contract offerers that can dynamically generate orders. Introduced in Seaport v1.4 and currently unsupported
	OrderTypeContract
)

type Type int

const (
	TypeBasic Type = iota
	TypeDutch
	TypeEnglish
	TypeCriteria
)

func (t Type) Value() string {
	switch t {
	case TypeBasic:
		return "basic"
	case TypeDutch:
		return "dutch"
	case TypeEnglish:
		return "english"
	case TypeCriteria:
		return "criteria"
	default:
		return ""
	}
}

// SafelistStatus is the status of the collection verification requests.
type SafelistStatus string

const (
	SafelistStatusNotRequested        SafelistStatus = "not_requested"
	SafelistStatusRequested           SafelistStatus = "requested"
	SafelistStatusApproved            SafelistStatus = "approved"
	SafelistStatusVerified            SafelistStatus = "verified"
	SafelistStatusDisabledTopTrending SafelistStatus = "disabled_top_trending"
)
