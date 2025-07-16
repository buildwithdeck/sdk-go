# Links
(*Links*)

## Overview

### Available Operations

* [Connect](#connect) - Connect with credentials
* [GetConnectionStatus](#getconnectionstatus) - Get the status of an attempted connection
* [ListAccounts](#listaccounts) - Return the list of accounts found while creating connection

## Connect

Sends credential information to a utility data source and attempts to create a connection request asynchronously.

### Example Usage

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
            ClientID: sdkgo.String(os.Getenv("DECK_CLIENT_ID")),
            Secret: sdkgo.String(os.Getenv("DECK_SECRET")),
        }),
    )

    res, err := s.Links.Connect(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LinkConnectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.LinkConnectRequest](../../models/components/linkconnectrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.PostLinkConnectResponse](../../models/operations/postlinkconnectresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ErrorMessageResponse | 400                            | application/json               |
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## GetConnectionStatus

Returns the connectivity status about a connection

### Example Usage

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
            ClientID: sdkgo.String(os.Getenv("DECK_CLIENT_ID")),
            Secret: sdkgo.String(os.Getenv("DECK_SECRET")),
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
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## ListAccounts

Return the list of accounts found while creating connection

### Example Usage

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
            ClientID: sdkgo.String(os.Getenv("DECK_CLIENT_ID")),
            Secret: sdkgo.String(os.Getenv("DECK_SECRET")),
        }),
    )

    res, err := s.Links.ListAccounts(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LinkAccountListResponse != nil {
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

**[*operations.PostLinkAccountListResponse](../../models/operations/postlinkaccountlistresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ErrorMessageResponse | 400                            | application/json               |
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |