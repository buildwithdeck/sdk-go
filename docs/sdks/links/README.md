# Links
(*Links*)

## Overview

### Available Operations

* [GetConnectionStatus](#getconnectionstatus) - Get Token Status

## GetConnectionStatus

Returns the current status of an SDK token.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/link/connection/status" method="post" path="/link/connection/status" -->
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

    res, err := s.Links.GetConnectionStatus(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LinkConnectionStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.LinkTokenRequestResponse](../../models/components/linktokenrequestresponse.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PostLinkConnectionStatusResponse](../../models/operations/postlinkconnectionstatusresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ErrorMessageResponse | 400                            | application/json               |
| apierrors.ErrorMessageResponse | 400                            | application/json+encrypted     |
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |