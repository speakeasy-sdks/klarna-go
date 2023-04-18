<div align="center">
    <img src="https://user-images.githubusercontent.com/6267663/230347878-f2873a58-f578-4e95-86e0-7bebfd78f4f1.svg" width="300">
    <h1>Orders Go SDK</h1>
   <p>An effortless integration. Designed for growth.</p>
   <a href="https://docs.klarna.com/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/klarna-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/klarna-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/klarna-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/klarna-go?sort=semver&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/klarna-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/klarna-go"
    "github.com/speakeasy-sdks/klarna-go/pkg/models/shared"
    "github.com/speakeasy-sdks/klarna-go/pkg/models/operations"
)

func main() {
    s := klarna.New(
        klarna.WithSecurity(shared.Security{
            APIKeyAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CreateOrderRequest{
        AuthorizationToken: "corrupti",
        CreateOrderRequestInput: &shared.CreateOrderRequestInput{
            AutoCapture: false,
            BillingAddress: &shared.Address{
                Attention: "Attn",
                City: "London",
                Country: "GB",
                Email: "test.sam@test.com",
                FamilyName: "Andersson",
                GivenName: "Adam",
                OrganizationName: "provident",
                Phone: "+44795465131",
                PostalCode: "W1G 0PW",
                Region: "OH",
                StreetAddress: "33 Cavendish Square",
                StreetAddress2: "Floor 22 / Flat 2",
                Title: "Mr.",
            },
            CustomPaymentMethodIds: []string{
                "quibusdam",
                "unde",
                "nulla",
            },
            Customer: &shared.Customer{
                DateOfBirth: "1978-12-31",
                Gender: "male",
                LastFourSsn: "corrupti",
                NationalIdentificationNumber: "illum",
                OrganizationEntityType: "LIMITED_PARTNERSHIP",
                OrganizationRegistrationID: "error",
                Title: "Mr.",
                Type: "organization",
                VatID: "deserunt",
            },
            Locale: "en-GB",
            MerchantData: "{"order_specific":[{"substore":"Women's Fashion","product_name":"Women Sweatshirt"}]}",
            MerchantReference1: "ON4711",
            MerchantReference2: "hdt53h-zdgg6-hdaff2",
            MerchantUrls: &shared.MerchantUrls{
                Authorization: "https://www.example-url.com/authorization",
                Confirmation: "https://www.example-url.com/confirmation",
                Notification: "https://www.example-url.com/notification",
                Push: "https://www.example-url.com/push",
            },
            OrderAmount: 2500,
            OrderLines: []shared.OrderLine{
                shared.OrderLine{
                    ImageURL: "https://www.exampleobjects.com/logo.png",
                    MerchantData: "{"customer_account_info":[{"unique_account_identifier":"test@gmail.com","account_registration_date":"2017-02-13T10:49:20Z","account_last_modified":"2019-03-13T11:45:27Z"}]}",
                    Name: "Running shoe",
                    ProductIdentifiers: &shared.ProductIdentifiers{
                        Brand: "shoe-brand",
                        CategoryPath: "Shoes > Running",
                        Color: "white",
                        GlobalTradeItemNumber: "4912345678904",
                        ManufacturerPartNumber: "AD6654412-334.22",
                        Size: "small",
                    },
                    ProductURL: "https://.../AD6654412.html",
                    Quantity: 1,
                    QuantityUnit: "pcs",
                    Reference: "AD6654412",
                    Subscription: &shared.Subscription{
                        Interval: "WEEK",
                        IntervalCount: 297534,
                        Name: "Larry Windler",
                    },
                    TaxRate: 1900,
                    TotalAmount: 2500,
                    TotalDiscountAmount: 500,
                    TotalTaxAmount: 475,
                    Type: "physical",
                    UnitPrice: 2500,
                },
                shared.OrderLine{
                    ImageURL: "https://www.exampleobjects.com/logo.png",
                    MerchantData: "{"customer_account_info":[{"unique_account_identifier":"test@gmail.com","account_registration_date":"2017-02-13T10:49:20Z","account_last_modified":"2019-03-13T11:45:27Z"}]}",
                    Name: "Running shoe",
                    ProductIdentifiers: &shared.ProductIdentifiers{
                        Brand: "shoe-brand",
                        CategoryPath: "Shoes > Running",
                        Color: "white",
                        GlobalTradeItemNumber: "4912345678904",
                        ManufacturerPartNumber: "AD6654412-334.22",
                        Size: "small",
                    },
                    ProductURL: "https://.../AD6654412.html",
                    Quantity: 1,
                    QuantityUnit: "pcs",
                    Reference: "AD6654412",
                    Subscription: &shared.Subscription{
                        Interval: "WEEK",
                        IntervalCount: 791725,
                        Name: "Ken Kshlerin",
                    },
                    TaxRate: 1900,
                    TotalAmount: 2500,
                    TotalDiscountAmount: 500,
                    TotalTaxAmount: 475,
                    Type: "physical",
                    UnitPrice: 2500,
                },
            },
            OrderTaxAmount: 475,
            PurchaseCountry: "GB",
            PurchaseCurrency: "GBP",
            ShippingAddress: &shared.Address{
                Attention: "Attn",
                City: "London",
                Country: "GB",
                Email: "test.sam@test.com",
                FamilyName: "Andersson",
                GivenName: "Adam",
                OrganizationName: "recusandae",
                Phone: "+44795465131",
                PostalCode: "W1G 0PW",
                Region: "OH",
                StreetAddress: "33 Cavendish Square",
                StreetAddress2: "Floor 22 / Flat 2",
                Title: "Mr.",
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### Authorizations

* `Cancel` - Cancel an existing authorization

### Orders

* `Create` - Create a new order

### Sessions

* `Create` - Create a new payment session
* `Read` - Read an existing payment session
* `Update` - Update an existing payment session

### Tokens

* `Purchase` - Generate a consumer token
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
