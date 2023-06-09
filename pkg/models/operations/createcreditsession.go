// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/klarna-go/v2/pkg/models/shared"
	"net/http"
)

type CreateCreditSessionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// successful operation
	MerchantSession *shared.MerchantSession
}
