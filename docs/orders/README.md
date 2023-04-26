# Orders

## Overview

Operations related to orders

### Available Operations

* [Create](#create) - Create a new order

## Create

Use this API call to create a new order. Placing an order towards Klarna means that the Klarna Payments session will be closed and that an order will be created in Klarna's system.<br/>When you have received the `authorization_token` for a successful authorization you can place the order. Among the other order details in this request, you include a URL to the confirmation page for the customer.<br/>When the Order has been successfully placed at Klarna, you need to handle it either through the Merchant Portal or using [Klarna’s Order Management API](#order-management-api).
Read more on **[Create a new order](https://docs.klarna.com/klarna-payments/integrate-with-klarna-payments/step-3-create-an-order/)**.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/klarna-go"
	"github.com/speakeasy-sdks/klarna-go/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/klarna-go/v2/pkg/models/shared"
)

func main() {
    s := klarna.New(
        klarna.WithSecurity(shared.Security{
            APIKeyAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateOrderRequest{
        AuthorizationToken: "ab",
        CreateOrderRequestInput: &shared.CreateOrderRequestInput{
            AutoCapture: klarna.Bool(false),
            BillingAddress: &shared.Address{
                Attention: klarna.String("Attn"),
                City: klarna.String("London"),
                Country: klarna.String("GB"),
                Email: klarna.String("test.sam@test.com"),
                FamilyName: klarna.String("Andersson"),
                GivenName: klarna.String("Adam"),
                OrganizationName: klarna.String("quis"),
                Phone: klarna.String("+44795465131"),
                PostalCode: klarna.String("W1G 0PW"),
                Region: klarna.String("OH"),
                StreetAddress: klarna.String("33 Cavendish Square"),
                StreetAddress2: klarna.String("Floor 22 / Flat 2"),
                Title: klarna.String("Mr."),
            },
            CustomPaymentMethodIds: []string{
                "deserunt",
            },
            Customer: &shared.Customer{
                DateOfBirth: klarna.String("1978-12-31"),
                Gender: klarna.String("male"),
                LastFourSsn: klarna.String("perferendis"),
                NationalIdentificationNumber: klarna.String("ipsam"),
                OrganizationEntityType: shared.CustomerOrganizationEntityTypeEnumPublicInstitution.ToPointer(),
                OrganizationRegistrationID: klarna.String("sapiente"),
                Title: klarna.String("Mr."),
                Type: klarna.String("organization"),
                VatID: klarna.String("quo"),
            },
            Locale: klarna.String("en-GB"),
            MerchantData: klarna.String("{"order_specific":[{"substore":"Women's Fashion","product_name":"Women Sweatshirt"}]}"),
            MerchantReference1: klarna.String("ON4711"),
            MerchantReference2: klarna.String("hdt53h-zdgg6-hdaff2"),
            MerchantUrls: &shared.MerchantUrls{
                Authorization: klarna.String("https://www.example-url.com/authorization"),
                Confirmation: klarna.String("https://www.example-url.com/confirmation"),
                Notification: klarna.String("https://www.example-url.com/notification"),
                Push: klarna.String("https://www.example-url.com/push"),
            },
            OrderAmount: 2500,
            OrderLines: []shared.OrderLine{
                shared.OrderLine{
                    ImageURL: klarna.String("https://www.exampleobjects.com/logo.png"),
                    MerchantData: klarna.String("{"customer_account_info":[{"unique_account_identifier":"test@gmail.com","account_registration_date":"2017-02-13T10:49:20Z","account_last_modified":"2019-03-13T11:45:27Z"}]}"),
                    Name: "Running shoe",
                    ProductIdentifiers: &shared.ProductIdentifiers{
                        Brand: klarna.String("shoe-brand"),
                        CategoryPath: klarna.String("Shoes > Running"),
                        Color: klarna.String("white"),
                        GlobalTradeItemNumber: klarna.String("4912345678904"),
                        ManufacturerPartNumber: klarna.String("AD6654412-334.22"),
                        Size: klarna.String("small"),
                    },
                    ProductURL: klarna.String("https://.../AD6654412.html"),
                    Quantity: 1,
                    QuantityUnit: klarna.String("pcs"),
                    Reference: klarna.String("AD6654412"),
                    Subscription: &shared.Subscription{
                        Interval: shared.SubscriptionIntervalEnumYear,
                        IntervalCount: 870088,
                        Name: "Javier Schmidt",
                    },
                    TaxRate: klarna.Int64(1900),
                    TotalAmount: 2500,
                    TotalDiscountAmount: klarna.Int64(500),
                    TotalTaxAmount: klarna.Int64(475),
                    Type: klarna.String("physical"),
                    UnitPrice: 2500,
                },
            },
            OrderTaxAmount: klarna.Int64(475),
            PurchaseCountry: "GB",
            PurchaseCurrency: "GBP",
            ShippingAddress: &shared.Address{
                Attention: klarna.String("Attn"),
                City: klarna.String("London"),
                Country: klarna.String("GB"),
                Email: klarna.String("test.sam@test.com"),
                FamilyName: klarna.String("Andersson"),
                GivenName: klarna.String("Adam"),
                OrganizationName: klarna.String("totam"),
                Phone: klarna.String("+44795465131"),
                PostalCode: klarna.String("W1G 0PW"),
                Region: klarna.String("OH"),
                StreetAddress: klarna.String("33 Cavendish Square"),
                StreetAddress2: klarna.String("Floor 22 / Flat 2"),
                Title: klarna.String("Mr."),
            },
        },
    }

    res, err := s.Orders.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```
