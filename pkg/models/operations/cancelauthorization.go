// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CancelAuthorizationRequest struct {
	AuthorizationToken string `pathParam:"style=simple,explode=false,name=authorizationToken"`
}

type CancelAuthorizationResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
