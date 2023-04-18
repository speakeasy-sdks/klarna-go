// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/klarna-go/v2/pkg/models/shared"
	"net/http"
)

type PurchaseTokenRequest struct {
	AuthorizationToken           string                               `pathParam:"style=simple,explode=false,name=authorizationToken"`
	CustomerTokenCreationRequest *shared.CustomerTokenCreationRequest `request:"mediaType=application/json"`
}

type PurchaseTokenResponse struct {
	ContentType string
	// We were unable to create a customer token with the provided data. Some field constraint was violated.
	ErrorV2     *shared.ErrorV2
	StatusCode  int
	RawResponse *http.Response
	// Token was successfully created.
	CustomerTokenCreationResponse *shared.CustomerTokenCreationResponse
}