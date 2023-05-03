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
    res, err := s.Orders.Create(ctx, "corrupti", &shared.CreateOrderRequestInput{
        AutoCapture: klarna.Bool(false),
        BillingAddress: &shared.Address{
            Attention: klarna.String("Attn"),
            City: klarna.String("London"),
            Country: klarna.String("GB"),
            Email: klarna.String("test.sam@test.com"),
            FamilyName: klarna.String("Andersson"),
            GivenName: klarna.String("Adam"),
            OrganizationName: klarna.String("provident"),
            Phone: klarna.String("+44795465131"),
            PostalCode: klarna.String("W1G 0PW"),
            Region: klarna.String("OH"),
            StreetAddress: klarna.String("33 Cavendish Square"),
            StreetAddress2: klarna.String("Floor 22 / Flat 2"),
            Title: klarna.String("Mr."),
        },
        CustomPaymentMethodIds: []string{
            "quibusdam",
            "unde",
            "nulla",
        },
        Customer: &shared.Customer{
            DateOfBirth: klarna.String("1978-12-31"),
            Gender: klarna.String("male"),
            LastFourSsn: klarna.String("corrupti"),
            NationalIdentificationNumber: klarna.String("illum"),
            OrganizationEntityType: shared.CustomerOrganizationEntityTypeEnumLimitedPartnership.ToPointer(),
            OrganizationRegistrationID: klarna.String("error"),
            Title: klarna.String("Mr."),
            Type: klarna.String("organization"),
            VatID: klarna.String("deserunt"),
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
                    Interval: shared.SubscriptionIntervalEnumWeek,
                    IntervalCount: 297534,
                    Name: "Larry Windler",
                },
                TaxRate: klarna.Int64(1900),
                TotalAmount: 2500,
                TotalDiscountAmount: klarna.Int64(500),
                TotalTaxAmount: klarna.Int64(475),
                Type: klarna.String("physical"),
                UnitPrice: 2500,
            },
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
                    Interval: shared.SubscriptionIntervalEnumWeek,
                    IntervalCount: 791725,
                    Name: "Ken Kshlerin",
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
            OrganizationName: klarna.String("recusandae"),
            Phone: klarna.String("+44795465131"),
            PostalCode: klarna.String("W1G 0PW"),
            Region: klarna.String("OH"),
            StreetAddress: klarna.String("33 Cavendish Square"),
            StreetAddress2: klarna.String("Floor 22 / Flat 2"),
            Title: klarna.String("Mr."),
        },
    })
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


### [Authorizations](docs/authorizations/README.md)

* [Cancel](docs/authorizations/README.md#cancel) - Cancel an existing authorization

### [Orders](docs/orders/README.md)

* [Create](docs/orders/README.md#create) - Create a new order

### [Sessions](docs/sessions/README.md)

* [Create](docs/sessions/README.md#create) - Create a new payment session
* [Read](docs/sessions/README.md#read) - Read an existing payment session
* [Update](docs/sessions/README.md#update) - Update an existing payment session

### [Tokens](docs/tokens/README.md)

* [Purchase](docs/tokens/README.md#purchase) - Generate a consumer token
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
