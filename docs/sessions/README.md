# Sessions

## Overview

Operations related to sessions

### Available Operations

* [Create](#create) - Create a new payment session
* [Read](#read) - Read an existing payment session
* [Update](#update) - Update an existing payment session

## Create

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
    req := shared.SessionCreateInput{
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
            OrganizationName: klarna.String("porro"),
            Phone: klarna.String("+44795465131"),
            PostalCode: klarna.String("W1G 0PW"),
            Region: klarna.String("OH"),
            StreetAddress: klarna.String("33 Cavendish Square"),
            StreetAddress2: klarna.String("Floor 22 / Flat 2"),
            Title: klarna.String("Mr."),
        },
        CustomPaymentMethodIds: []string{
            "dicta",
            "nam",
            "officia",
        },
        Customer: &shared.Customer{
            DateOfBirth: klarna.String("1978-12-31"),
            Gender: klarna.String("male"),
            LastFourSsn: klarna.String("occaecati"),
            NationalIdentificationNumber: klarna.String("fugit"),
            OrganizationEntityType: shared.CustomerOrganizationEntityTypeEnumGeneralPartnership.ToPointer(),
            OrganizationRegistrationID: klarna.String("hic"),
            Title: klarna.String("Mr."),
            Type: klarna.String("organization"),
            VatID: klarna.String("optio"),
        },
        Design: klarna.String("totam"),
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
                    Interval: shared.SubscriptionIntervalEnumWeek,
                    IntervalCount: 473600,
                    Name: "Norma Ryan",
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
            OrganizationName: klarna.String("ipsum"),
            Phone: klarna.String("+44795465131"),
            PostalCode: klarna.String("W1G 0PW"),
            Region: klarna.String("OH"),
            StreetAddress: klarna.String("33 Cavendish Square"),
            StreetAddress2: klarna.String("Floor 22 / Flat 2"),
            Title: klarna.String("Mr."),
        },
    }

    res, err := s.Sessions.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.MerchantSession != nil {
        // handle response
    }
}
```

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
    req := operations.ReadCreditSessionRequest{
        SessionID: "excepturi",
    }

    res, err := s.Sessions.Read(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SessionRead != nil {
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
    req := operations.UpdateCreditSessionRequest{
        SessionInput: shared.SessionInput{
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
                OrganizationName: klarna.String("aspernatur"),
                Phone: klarna.String("+44795465131"),
                PostalCode: klarna.String("W1G 0PW"),
                Region: klarna.String("OH"),
                StreetAddress: klarna.String("33 Cavendish Square"),
                StreetAddress2: klarna.String("Floor 22 / Flat 2"),
                Title: klarna.String("Mr."),
            },
            CustomPaymentMethodIds: []string{
                "ad",
            },
            Customer: &shared.Customer{
                DateOfBirth: klarna.String("1978-12-31"),
                Gender: klarna.String("male"),
                LastFourSsn: klarna.String("natus"),
                NationalIdentificationNumber: klarna.String("sed"),
                OrganizationEntityType: shared.CustomerOrganizationEntityTypeEnumRegisteredSoleTrader.ToPointer(),
                OrganizationRegistrationID: klarna.String("dolor"),
                Title: klarna.String("Mr."),
                Type: klarna.String("organization"),
                VatID: klarna.String("natus"),
            },
            Design: klarna.String("laboriosam"),
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
                        Interval: shared.SubscriptionIntervalEnumYear,
                        IntervalCount: 681820,
                        Name: "Stacy Moore",
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
                        IntervalCount: 99280,
                        Name: "Lela Orn",
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
                        IntervalCount: 210382,
                        Name: "Rose Rolfson",
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
                        IntervalCount: 325047,
                        Name: "Brian Kessler",
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
                OrganizationName: klarna.String("sapiente"),
                Phone: klarna.String("+44795465131"),
                PostalCode: klarna.String("W1G 0PW"),
                Region: klarna.String("OH"),
                StreetAddress: klarna.String("33 Cavendish Square"),
                StreetAddress2: klarna.String("Floor 22 / Flat 2"),
                Title: klarna.String("Mr."),
            },
        },
        SessionID: "architecto",
    }

    res, err := s.Sessions.Update(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
