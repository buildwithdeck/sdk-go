# Agents
(*Agents*)

## Overview

Agent management

### Available Operations

* [PostAgentsPrepare](#postagentsprepare) - Prepares an agent in the background so that it is ready for a subsequent ensure connection job request.

## PostAgentsPrepare

Prepares an agent in the background so that it is ready for a subsequent ensure connection job request.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/agents/prepare" method="post" path="/agents/prepare" -->
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

    res, err := s.Agents.PostAgentsPrepare(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.PrepareAgentRequest](../../models/components/prepareagentrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PostAgentsPrepareResponse](../../models/operations/postagentsprepareresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ProblemDetailsError | 400, 404                      | application/json              |
| apierrors.ProblemDetailsError | 400, 404                      | application/json+encrypted    |
| apierrors.ProblemDetailsError | 400, 404                      | text/json                     |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |