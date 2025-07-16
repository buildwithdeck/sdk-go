# JobsDocuments
(*JobsDocuments*)

## Overview

### Available Operations

* [List](#list) - List documents

## List

Returns a list of documents available for the connection associated with the provided link token

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

    res, err := s.JobsDocuments.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DocumentListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |
| `jobGUID`                                                                       | **string*                                                                       | :heavy_minus_sign:                                                              | N/A                                                                             |
| `accessTokenRequest`                                                            | [*components.AccessTokenRequest](../../models/components/accesstokenrequest.md) | :heavy_minus_sign:                                                              | N/A                                                                             |
| `opts`                                                                          | [][operations.Option](../../models/operations/option.md)                        | :heavy_minus_sign:                                                              | The options for this request.                                                   |

### Response

**[*operations.PostJobsDocumentsListResponse](../../models/operations/postjobsdocumentslistresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ErrorMessageResponse | 400                            | application/json               |
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |