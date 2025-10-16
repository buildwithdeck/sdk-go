# WebhookSubscription
(*WebhookSubscription*)

## Overview

Webhook subscriptions for the team. Use this to subscribe to and unsubscribe from webhook events. See the documentation for more information on how to use this API.

### Available Operations

* [PostWebhookSubscriptionsEventTypes](#postwebhooksubscriptionseventtypes) - Get all available webhook event types
* [PostWebhookSubscriptionsSubscriptions](#postwebhooksubscriptionssubscriptions) - Get webhook subscriptions for a specific team
* [PostWebhookSubscriptionsSubscribe](#postwebhooksubscriptionssubscribe) - Subscribe to webhook events
* [PostWebhookSubscriptionsUnsubscribe](#postwebhooksubscriptionsunsubscribe) - Unsubscribe from webhook events

## PostWebhookSubscriptionsEventTypes

Get all available webhook event types

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-subscriptions/event-types" method="post" path="/webhook-subscriptions/event-types" -->
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

    res, err := s.WebhookSubscription.PostWebhookSubscriptionsEventTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookEventTypesResponse != nil {
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

**[*operations.PostWebhookSubscriptionsEventTypesResponse](../../models/operations/postwebhooksubscriptionseventtypesresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ErrorMessageResponse | 400                            | application/json               |
| apierrors.ErrorMessageResponse | 400                            | application/json+encrypted     |
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## PostWebhookSubscriptionsSubscriptions

Get webhook subscriptions for a specific team

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-subscriptions/subscriptions" method="post" path="/webhook-subscriptions/subscriptions" -->
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

    res, err := s.WebhookSubscription.PostWebhookSubscriptionsSubscriptions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONWebhookSubscriptionResponses != nil {
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

**[*operations.PostWebhookSubscriptionsSubscriptionsResponse](../../models/operations/postwebhooksubscriptionssubscriptionsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ErrorMessageResponse | 400                            | application/json               |
| apierrors.ErrorMessageResponse | 400                            | application/json+encrypted     |
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## PostWebhookSubscriptionsSubscribe

Subscribe to webhook events

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-subscriptions/subscribe" method="post" path="/webhook-subscriptions/subscribe" -->
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

    res, err := s.WebhookSubscription.PostWebhookSubscriptionsSubscribe(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.WebhookSubscriptionRequest](../../models/components/webhooksubscriptionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PostWebhookSubscriptionsSubscribeResponse](../../models/operations/postwebhooksubscriptionssubscriberesponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ProblemDetailsError | 400                           | application/json              |
| apierrors.ProblemDetailsError | 400                           | application/json+encrypted    |
| apierrors.ProblemDetailsError | 400                           | text/json                     |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## PostWebhookSubscriptionsUnsubscribe

Unsubscribe from webhook events

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-subscriptions/unsubscribe" method="post" path="/webhook-subscriptions/unsubscribe" -->
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

    res, err := s.WebhookSubscription.PostWebhookSubscriptionsUnsubscribe(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.WebhookSubscriptionRequest](../../models/components/webhooksubscriptionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PostWebhookSubscriptionsUnsubscribeResponse](../../models/operations/postwebhooksubscriptionsunsubscriberesponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ProblemDetailsError | 400                           | application/json              |
| apierrors.ProblemDetailsError | 400                           | application/json+encrypted    |
| apierrors.ProblemDetailsError | 400                           | text/json                     |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |