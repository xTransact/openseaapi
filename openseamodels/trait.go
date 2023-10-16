package openseamodels

type NftTrait struct {
	// The name of the trait category (e.g. 'Background')
	TraitType string `json:"trait_type"`
	// A field indicating how to display. None is used for string traits.
	// number boost_percentage boost_number author date None
	DisplayType *string `json:"display_type"`
	// Ceiling for possible numeric trait values
	MaxValue *string `json:"max_value"`
	// Deprecated Field. Use Get Collection API instead.
	TraitCount int `json:"trait_count"`
	// Deprecated Field
	Order *int `json:"order"`
	// The value of the trait (e.g. 'Red')
	// Type could be: number | integer | date | string
	Value any `json:"value"`
}

type Trait struct {
	// List of trait categories, e.g. Background, in the collection and their type, e.g. string
	Categories map[string]any `json:"categories"`
	// If the category type is STRING, the dict will contain each trait value and its count.
	// Otherwise, the dict will contain the min and max value seen in the collection
	Counts map[string]any `json:"counts"`
}
