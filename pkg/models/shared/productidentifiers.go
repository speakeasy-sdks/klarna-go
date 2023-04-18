// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ProductIdentifiers struct {
	// The product's brand name as generally recognized by consumers. If no brand is available for a product, do not supply any value.
	Brand *string `json:"brand,omitempty"`
	// The product's category path as used in the merchant's webshop. Include the full and most detailed category and separate the segments with ' > '
	CategoryPath *string `json:"category_path,omitempty"`
	// Color to be shown to the end customer (max 64 characters).
	Color *string `json:"color,omitempty"`
	// The product's Global Trade Item Number (GTIN). Common types of GTIN are EAN, ISBN or UPC. Exclude dashes and spaces, where possible
	GlobalTradeItemNumber *string `json:"global_trade_item_number,omitempty"`
	// The product's Manufacturer Part Number (MPN), which - together with the brand - uniquely identifies a product. Only submit MPNs assigned by a manufacturer and use the most specific MPN possible
	ManufacturerPartNumber *string `json:"manufacturer_part_number,omitempty"`
	// Size to be shown to the end customer (max 64 characters).
	Size *string `json:"size,omitempty"`
}