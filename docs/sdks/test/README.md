# Test
(*Test*)

## Overview

### Available Operations

* [GetTestAPIKeys](#gettestapikeys) - Test your API keys

## GetTestAPIKeys

Provide your API keys in the usual x-deck-client-id and x-deck-secret headers and see if they get validated.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/test/api_keys" method="get" path="/test/api_keys" -->
```go
package main

import(
	"context"
	"os"
	"github.com/buildwithdeck/sdk-go/models/components"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithSecurity(components.Security{
            ClientID: sdkgo.Pointer(os.Getenv("DECK_CLIENT_ID")),
            Secret: sdkgo.Pointer(os.Getenv("DECK_SECRET")),
        }),
    )

    res, err := s.Test.GetTestAPIKeys(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTestAPIKeysResponse](../../models/operations/gettestapikeysresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ErrorMessageResponse | 400                            | application/json               |
| apierrors.ErrorMessageResponse | 400                            | application/json+encrypted     |
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |