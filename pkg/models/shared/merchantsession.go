// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MerchantSession - successful operation
type MerchantSession struct {
	// Client token to be passed to the JS client while initializing the JS SDK in the next step.
	ClientToken string `json:"client_token"`
	// Available payment method categories for this particular session
	PaymentMethodCategories []PaymentMethodCategory `json:"payment_method_categories,omitempty"`
	// ID of the created session. Please use this ID to share with Klarna for identifying any issues during integration.
	SessionID string `json:"session_id"`
}
