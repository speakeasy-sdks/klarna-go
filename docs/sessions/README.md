# Sessions

## Overview

Operations related to sessions

### Available Operations

* [Read](#read) - Read an existing payment session
* [Read](#read) - Create a new payment session
* [Update](#update) - Update an existing payment session

## Read

Use this API call to read a Klarna Payments session. You can read the Klarna Payments session at any time after it has been created, to get information about it. This will return all data that has been collected during the session.
Read more on **[Read an existing payment session](https://docs.klarna.com/klarna-payments/other-actions/check-the-details-of-a-payment-session/)**.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/klarna-go"
	"github.com/speakeasy-sdks/klarna-go/v2/pkg/models/operations"
)

func main() {
    s := klarna.New(
        klarna.WithSecurity(shared.Security{
            APIKeyAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Sessions.Read(ctx, "porro")
    if err != nil {
        log.Fatal(err)
    }

    if res.SessionRead != nil {
        // handle response
    }
}
```

## Read

Use this API call to create a Klarna Payments session.<br/>When a session is created you will receive the available `payment_method_categories` for the session, a `session_id` and a `client_token`. The `session_id` can be used to read or update the session using the REST API. The `client_token` should be passed to the browser.
Read more on **[Create a new payment session](https://docs.klarna.com/klarna-payments/integrate-with-klarna-payments/step-1-initiate-a-payment/)**.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/klarna-go"
	"github.com/speakeasy-sdks/klarna-go/v2/pkg/models/shared"
)

func main() {
    s := klarna.New(
        klarna.WithSecurity(shared.Security{
            APIKeyAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Sessions.Read(ctx, shared.SessionCreateInput{
        AcquiringChannel: shared.SessionCreateAcquiringChannelEnumEcommerce.ToPointer(),
        Attachment: &shared.Attachment{
            Body: "{"customer_account_info":[{"unique_account_identifier":"test@gmail.com","account_registration_date":"2017-02-13T10:49:20Z","account_last_modified":"2019-03-13T11:45:27Z"}]}",
            ContentType: "application/vnd.klarna.internal.emd-v2+json",
        },
        BillingAddress: &shared.Address{
            Attention: klarna.String("Attn"),
            City: klarna.String("London"),
            Country: klarna.String("GB"),
            Email: klarna.String("test.sam@test.com"),
            FamilyName: klarna.String("Andersson"),
            GivenName: klarna.String("Adam"),
            OrganizationName: klarna.String("dolorum"),
            Phone: klarna.String("+44795465131"),
            PostalCode: klarna.String("W1G 0PW"),
            Region: klarna.String("OH"),
            StreetAddress: klarna.String("33 Cavendish Square"),
            StreetAddress2: klarna.String("Floor 22 / Flat 2"),
            Title: klarna.String("Mr."),
        },
        CustomPaymentMethodIds: []string{
            "nam",
        },
        Customer: &shared.Customer{
            DateOfBirth: klarna.String("1978-12-31"),
            Gender: klarna.String("male"),
            LastFourSsn: klarna.String("officia"),
            NationalIdentificationNumber: klarna.String("occaecati"),
            OrganizationEntityType: shared.CustomerOrganizationEntityTypeEnumPublicLimitedCompany.ToPointer(),
            OrganizationRegistrationID: klarna.String("deleniti"),
            Title: klarna.String("Mr."),
            Type: klarna.String("organization"),
            VatID: klarna.String("hic"),
        },
        Design: klarna.String("optio"),
        Intent: shared.SessionCreateIntentEnumBuy.ToPointer(),
        Locale: klarna.String("en-US"),
        MerchantData: klarna.String("{"order_specific":[{"substore":"Women's Fashion","product_name":"Women Sweatshirt"}]}"),
        MerchantReference1: klarna.String("ON4711"),
        MerchantReference2: klarna.String("hdt53h-zdgg6-hdaff2"),
        MerchantUrls: &shared.MerchantUrls{
            Authorization: klarna.String("https://www.example-url.com/authorization"),
            Confirmation: klarna.String("https://www.example-url.com/confirmation"),
            Notification: klarna.String("https://www.example-url.com/notification"),
            Push: klarna.String("https://www.example-url.com/push"),
        },
        Options: &shared.Options{
            ColorBorder: klarna.String("#FF9900"),
            ColorBorderSelected: klarna.String("#FF9900"),
            ColorDetails: klarna.String("#FF9900"),
            ColorText: klarna.String("#FF9900"),
            RadiusBorder: klarna.String("5px"),
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
                    Interval: shared.SubscriptionIntervalEnumDay,
                    IntervalCount: 414662,
                    Name: "Pauline Dibbert",
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
                    IntervalCount: 216550,
                    Name: "Brandon Auer",
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
                    Interval: shared.SubscriptionIntervalEnumDay,
                    IntervalCount: 612096,
                    Name: "Faye Howe",
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
            OrganizationName: klarna.String("fuga"),
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

    if res.MerchantSession != nil {
        // handle response
    }
}
```

## Update

Use this API call to update a Klarna Payments session with new details, in case something in the order has changed and the checkout has been reloaded. Including if the consumer adds a new item to the cart or if consumer details are updated.
Read more on **[Update an existing payment session](https://docs.klarna.com/klarna-payments/other-actions/update-the-cart/)**.

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
    res, err := s.Sessions.Update(ctx, shared.SessionInput{
        AcquiringChannel: shared.SessionAcquiringChannelEnumEcommerce.ToPointer(),
        Attachment: &shared.Attachment{
            Body: "{"customer_account_info":[{"unique_account_identifier":"test@gmail.com","account_registration_date":"2017-02-13T10:49:20Z","account_last_modified":"2019-03-13T11:45:27Z"}]}",
            ContentType: "application/vnd.klarna.internal.emd-v2+json",
        },
        BillingAddress: &shared.Address{
            Attention: klarna.String("Attn"),
            City: klarna.String("London"),
            Country: klarna.String("GB"),
            Email: klarna.String("test.sam@test.com"),
            FamilyName: klarna.String("Andersson"),
            GivenName: klarna.String("Adam"),
            OrganizationName: klarna.String("in"),
            Phone: klarna.String("+44795465131"),
            PostalCode: klarna.String("W1G 0PW"),
            Region: klarna.String("OH"),
            StreetAddress: klarna.String("33 Cavendish Square"),
            StreetAddress2: klarna.String("Floor 22 / Flat 2"),
            Title: klarna.String("Mr."),
        },
        CustomPaymentMethodIds: []string{
            "iste",
            "iure",
        },
        Customer: &shared.Customer{
            DateOfBirth: klarna.String("1978-12-31"),
            Gender: klarna.String("male"),
            LastFourSsn: klarna.String("saepe"),
            NationalIdentificationNumber: klarna.String("quidem"),
            OrganizationEntityType: shared.CustomerOrganizationEntityTypeEnumPublicLimitedCompany.ToPointer(),
            OrganizationRegistrationID: klarna.String("ipsa"),
            Title: klarna.String("Mr."),
            Type: klarna.String("organization"),
            VatID: klarna.String("reiciendis"),
        },
        Design: klarna.String("est"),
        Intent: shared.SessionIntentEnumBuy.ToPointer(),
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
        Options: &shared.Options{
            ColorBorder: klarna.String("#FF9900"),
            ColorBorderSelected: klarna.String("#FF9900"),
            ColorDetails: klarna.String("#FF9900"),
            ColorText: klarna.String("#FF9900"),
            RadiusBorder: klarna.String("5px"),
        },
        OrderAmount: klarna.Int64(2500),
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
                    Interval: shared.SubscriptionIntervalEnumMonth,
                    IntervalCount: 170909,
                    Name: "Stacy Champlin",
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
                    Interval: shared.SubscriptionIntervalEnumMonth,
                    IntervalCount: 363711,
                    Name: "Velma Batz",
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
                    Interval: shared.SubscriptionIntervalEnumYear,
                    IntervalCount: 958950,
                    Name: "Angie Durgan",
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
        PurchaseCountry: klarna.String("GB"),
        PurchaseCurrency: klarna.String("GBP"),
        ShippingAddress: &shared.Address{
            Attention: klarna.String("Attn"),
            City: klarna.String("London"),
            Country: klarna.String("GB"),
            Email: klarna.String("test.sam@test.com"),
            FamilyName: klarna.String("Andersson"),
            GivenName: klarna.String("Adam"),
            OrganizationName: klarna.String("repellat"),
            Phone: klarna.String("+44795465131"),
            PostalCode: klarna.String("W1G 0PW"),
            Region: klarna.String("OH"),
            StreetAddress: klarna.String("33 Cavendish Square"),
            StreetAddress2: klarna.String("Floor 22 / Flat 2"),
            Title: klarna.String("Mr."),
        },
    }, "mollitia")
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
