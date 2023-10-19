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

func ValidateItemType(itemType int) bool {
	return itemType >= int(ItemTypeNative) && itemType <= int(ItemTypeERC1155WithCriteria)
}

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

func ValidateOrderType(orderType int) bool {
	return orderType >= int(OrderTypeFullOpen) && orderType <= int(OrderTypeContract)
}

type OrderSide uint8

const (
	OrderSideBuy OrderSide = 0
	OrderSideSell
)

func ValidateOrderSide(orderSide uint8) bool {
	return orderSide >= uint8(OrderSideBuy) && orderSide <= uint8(OrderSideSell)
}

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

func ValidateType(t int) bool {
	return t >= int(TypeBasic) && t <= int(TypeCriteria)
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

var safelistStatusList = []SafelistStatus{
	SafelistStatusNotRequested, SafelistStatusRequested, SafelistStatusApproved,
	SafelistStatusVerified, SafelistStatusDisabledTopTrending,
}

func ValidateSafelistStatus(safelistStatus string) bool {
	for _, s := range safelistStatusList {
		if string(s) == safelistStatus {
			return true
		}
	}
	return false
}
