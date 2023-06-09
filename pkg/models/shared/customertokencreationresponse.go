// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CustomerTokenCreationResponse - Token was successfully created.
type CustomerTokenCreationResponse struct {
	BillingAddress *Address                 `json:"billing_address,omitempty"`
	Customer       *CustomerReadCreateToken `json:"customer,omitempty"`
	// Used to connect customers with payment method when it is present.
	PaymentMethodReference *string `json:"payment_method_reference,omitempty"`
	// URL to redirect the customer to after placing the order. This is a Klarna URL where Klarna will place a cookie in the customer’s browser (if redirected) and redirect the customer back to the confirmation URL provided by the merchant. This is not a mandatory step but a recommended one to improve the returning customer’s experience.
	RedirectURL *string `json:"redirect_url,omitempty"`
	// Generated customer token. This token will be used to create a new order for the subscription using the Create a New order using token API.
	TokenID string `json:"token_id"`
}
