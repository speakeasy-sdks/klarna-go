// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package klarna

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/klarna-go/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/klarna-go/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/klarna-go/v2/pkg/utils"
	"net/http"
)

// orders - Operations related to orders
type orders struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newOrders(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *orders {
	return &orders{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// Read - Create a new order
// Use this API call to create a new order. Placing an order towards Klarna means that the Klarna Payments session will be closed and that an order will be created in Klarna's system.<br/>When you have received the `authorization_token` for a successful authorization you can place the order. Among the other order details in this request, you include a URL to the confirmation page for the customer.<br/>When the Order has been successfully placed at Klarna, you need to handle it either through the Merchant Portal or using [Klarna’s Order Management API](#order-management-api).
// Read more on **[Create a new order](https://docs.klarna.com/klarna-payments/integrate-with-klarna-payments/step-3-create-an-order/)**.

func (s *orders) Read(ctx context.Context, authorizationToken string, createOrderRequestInput *shared.CreateOrderRequestInput) (*operations.CreateOrderResponse, error) {
	request := operations.CreateOrderRequest{
		AuthorizationToken:      authorizationToken,
		CreateOrderRequestInput: createOrderRequestInput,
	}

	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/payments/v1/authorizations/{authorizationToken}/order", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "CreateOrderRequestInput", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateOrderResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Order
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Order = out
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 409:
	}

	return res, nil
}
