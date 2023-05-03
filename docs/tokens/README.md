# Tokens

## Overview

Operations related to token

### Available Operations

* [Purchase](#purchase) - Generate a consumer token

## Purchase

Use this API call to create a Klarna Customer Token.<br/>After having obtained an `authorization_token` for a successful authorization, this can be used to create a purchase token instead of placing the order. Creating a Klarna Customer Token results in Klarna storing customer and payment method details.
Read more on **[Generate a consumer token](https://docs.klarna.com/klarna-payments/in-depth-knowledge/customer-token/)**.

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
    res, err := s.Tokens.Purchase(ctx, operations.PurchaseTokenRequest{
        AuthorizationToken: "mollitia",
        CustomerTokenCreationRequest: &shared.CustomerTokenCreationRequest{
            BillingAddress: &shared.Address{
                Attention: klarna.String("Attn"),
                City: klarna.String("London"),
                Country: klarna.String("GB"),
                Email: klarna.String("test.sam@test.com"),
                FamilyName: klarna.String("Andersson"),
                GivenName: klarna.String("Adam"),
                OrganizationName: klarna.String("dolorem"),
                Phone: klarna.String("+44795465131"),
                PostalCode: klarna.String("W1G 0PW"),
                Region: klarna.String("OH"),
                StreetAddress: klarna.String("33 Cavendish Square"),
                StreetAddress2: klarna.String("Floor 22 / Flat 2"),
                Title: klarna.String("Mr."),
            },
            Customer: &shared.Customer{
                DateOfBirth: klarna.String("1978-12-31"),
                Gender: klarna.String("male"),
                LastFourSsn: klarna.String("culpa"),
                NationalIdentificationNumber: klarna.String("consequuntur"),
                OrganizationEntityType: shared.CustomerOrganizationEntityTypeEnumOther.ToPointer(),
                OrganizationRegistrationID: klarna.String("mollitia"),
                Title: klarna.String("Mr."),
                Type: klarna.String("organization"),
                VatID: klarna.String("occaecati"),
            },
            Description: "numquam",
            IntendedUse: shared.CustomerTokenCreationRequestIntendedUseEnumSubscription,
            Locale: "en-GB",
            PurchaseCountry: "GB",
            PurchaseCurrency: "GBP",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomerTokenCreationResponse != nil {
        // handle response
    }
}
```
