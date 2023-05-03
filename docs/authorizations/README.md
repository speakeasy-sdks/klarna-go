# Authorizations

## Overview

Operations related to authorizations

### Available Operations

* [Cancel](#cancel) - Cancel an existing authorization

## Cancel

Use this API call to cancel/release an authorization. If the `authorization_token` received during a Klarna Payments won’t be used to place an order immediately you could release the authorization.
Read more on **[Cancel an existing authorization](https://docs.klarna.com/klarna-payments/other-actions/cancel-an-authorization/)**.

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
    res, err := s.Authorizations.Cancel(ctx, "temporibus")
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
