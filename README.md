# github.com/buildwithdeck/sdk-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/buildwithdeck/sdk-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/buildwithdeck/sdk-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<!-- Start Summary [summary] -->
## Summary

Deck API: # Deck API makes it straightforward for users to connect to any portal securely and quickly.

### Welcome! Looking for a quick introduction to our API? Check out the 🚀[Quickstart guide](https://framework.docs.deck.co/docs/setup).

Starting on the sandbox server is easy:

1. Create an account using the [dashboard](https://dashboard.deck.co) to get your client id and sandbox secret
2. Enter your client id and sandbox secret in the Authentication section below
3. Hit the "Try" buttons below for each endpoint.

Happy querying!
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/buildwithdeck/sdk-go](#githubcombuildwithdecksdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/buildwithdeck/sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"github.com/buildwithdeck/sdk-go/models/components"
	"github.com/buildwithdeck/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(components.Security{
			ClientID: sdkgo.Pointer(os.Getenv("DECK_CLIENT_ID")),
			Secret:   sdkgo.Pointer(os.Getenv("DECK_SECRET")),
		}),
	)

	res, err := s.Jobs.Submit(ctx, nil, &operations.PostJobsSubmitRequestBody2{
		JobCode: "FetchDocuments",
		Input: map[string]operations.InputUnion2{
			"access_token": operations.CreateInputUnion2Str(
				"access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
			),
			"key1": operations.CreateInputUnion2Str(
				"value1",
			),
			"someNumber": operations.CreateInputUnion2Number(
				123.45,
			),
			"someBoolean": operations.CreateInputUnion2Boolean(
				true,
			),
			"someArray": operations.CreateInputUnion2ArrayOfStr(
				[]string{
					"a",
					"b",
				},
			),
			"nestedObject": operations.CreateInputUnion2Boolean(
				true,
			),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.IJobResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name       | Type   | Scheme  | Environment Variable |
| ---------- | ------ | ------- | -------------------- |
| `ClientID` | apiKey | API key | `DECK_CLIENT_ID`     |
| `Secret`   | apiKey | API key | `DECK_SECRET`        |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"github.com/buildwithdeck/sdk-go/models/components"
	"github.com/buildwithdeck/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(components.Security{
			ClientID: sdkgo.Pointer(os.Getenv("DECK_CLIENT_ID")),
			Secret:   sdkgo.Pointer(os.Getenv("DECK_SECRET")),
		}),
	)

	res, err := s.Jobs.Submit(ctx, nil, &operations.PostJobsSubmitRequestBody2{
		JobCode: "FetchDocuments",
		Input: map[string]operations.InputUnion2{
			"access_token": operations.CreateInputUnion2Str(
				"access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
			),
			"key1": operations.CreateInputUnion2Str(
				"value1",
			),
			"someNumber": operations.CreateInputUnion2Number(
				123.45,
			),
			"someBoolean": operations.CreateInputUnion2Boolean(
				true,
			),
			"someArray": operations.CreateInputUnion2ArrayOfStr(
				[]string{
					"a",
					"b",
				},
			),
			"nestedObject": operations.CreateInputUnion2Boolean(
				true,
			),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.IJobResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Agents](docs/sdks/agents/README.md)

* [PostAgentsPrepare](docs/sdks/agents/README.md#postagentsprepare) - Prepares an agent in the background so that it is ready for a subsequent ensure connection job request.

### [Connection](docs/sdks/connection/README.md)

* [ExchangePublicToken](docs/sdks/connection/README.md#exchangepublictoken) - Exchange public token for an access token
* [GetContext](docs/sdks/connection/README.md#getcontext) - Get connection status details
* [GetAccounts](docs/sdks/connection/README.md#getaccounts) - Get the connection list of account numbers
* [SelectAccounts](docs/sdks/connection/README.md#selectaccounts) - Update the list of accounts to be considered
* [UpdateAutoRefresh](docs/sdks/connection/README.md#updateautorefresh) - Update connection stream config
* [Refresh](docs/sdks/connection/README.md#refresh) - Request manual refresh
* [Destroy](docs/sdks/connection/README.md#destroy) - Delete all data related to a connection, losing access to it.

### [Connections](docs/sdks/connections/README.md)

* [InvalidateAccessToken](docs/sdks/connections/README.md#invalidateaccesstoken) - Invalidate access_token
* [UpdateWebhook](docs/sdks/connections/README.md#updatewebhook) - Update the webhook url for a connection


### [Jobs](docs/sdks/jobs/README.md)

* [Submit](docs/sdks/jobs/README.md#submit) - Send your job requests
* [AnswerMFA](docs/sdks/jobs/README.md#answermfa) - Provide MFA code
* [GetDocumentFile](docs/sdks/jobs/README.md#getdocumentfile) - Get raw file

### [JobsDocuments](docs/sdks/jobsdocuments/README.md)

* [List](docs/sdks/jobsdocuments/README.md#list) - List documents

### [Link](docs/sdks/link/README.md)

* [CreateToken](docs/sdks/link/README.md#createtoken) - Create a Link Token for running a Link session
* [GetInfo](docs/sdks/link/README.md#getinfo) - Get client information
* [SearchSources](docs/sdks/link/README.md#searchsources) - Search sources
* [SelectAccount](docs/sdks/link/README.md#selectaccount) - Select accounts
* [GetToken](docs/sdks/link/README.md#gettoken) - Get information about a previously created link_token

### [Links](docs/sdks/links/README.md)

* [Connect](docs/sdks/links/README.md#connect) - Connect with credentials
* [GetConnectionStatus](docs/sdks/links/README.md#getconnectionstatus) - Get the status of an attempted connection
* [ListAccounts](docs/sdks/links/README.md#listaccounts) - Return the list of accounts found while creating connection

#### [Links.Authentication](docs/sdks/authentication/README.md)

* [GetMfaQuestion](docs/sdks/authentication/README.md#getmfaquestion) - Get the security question
* [AnswerMfa](docs/sdks/authentication/README.md#answermfa) - Provide MFA code

### [Test](docs/sdks/test/README.md)

* [GetTestAPIKeys](docs/sdks/test/README.md#gettestapikeys) - Test your API keys

### [WebhookSubscription](docs/sdks/webhooksubscription/README.md)

* [PostWebhookSubscriptionsEventTypes](docs/sdks/webhooksubscription/README.md#postwebhooksubscriptionseventtypes) - Get all available webhook event types
* [PostWebhookSubscriptionsSubscriptions](docs/sdks/webhooksubscription/README.md#postwebhooksubscriptionssubscriptions) - Get webhook subscriptions for a specific team
* [PostWebhookSubscriptionsSubscribe](docs/sdks/webhooksubscription/README.md#postwebhooksubscriptionssubscribe) - Subscribe to webhook events
* [PostWebhookSubscriptionsUnsubscribe](docs/sdks/webhooksubscription/README.md#postwebhooksubscriptionsunsubscribe) - Unsubscribe from webhook events

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"github.com/buildwithdeck/sdk-go/models/components"
	"github.com/buildwithdeck/sdk-go/models/operations"
	"github.com/buildwithdeck/sdk-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(components.Security{
			ClientID: sdkgo.Pointer(os.Getenv("DECK_CLIENT_ID")),
			Secret:   sdkgo.Pointer(os.Getenv("DECK_SECRET")),
		}),
	)

	res, err := s.Jobs.Submit(ctx, nil, &operations.PostJobsSubmitRequestBody2{
		JobCode: "FetchDocuments",
		Input: map[string]operations.InputUnion2{
			"access_token": operations.CreateInputUnion2Str(
				"access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
			),
			"key1": operations.CreateInputUnion2Str(
				"value1",
			),
			"someNumber": operations.CreateInputUnion2Number(
				123.45,
			),
			"someBoolean": operations.CreateInputUnion2Boolean(
				true,
			),
			"someArray": operations.CreateInputUnion2ArrayOfStr(
				[]string{
					"a",
					"b",
				},
			),
			"nestedObject": operations.CreateInputUnion2Boolean(
				true,
			),
		},
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.IJobResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"github.com/buildwithdeck/sdk-go/models/components"
	"github.com/buildwithdeck/sdk-go/models/operations"
	"github.com/buildwithdeck/sdk-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		sdkgo.WithSecurity(components.Security{
			ClientID: sdkgo.Pointer(os.Getenv("DECK_CLIENT_ID")),
			Secret:   sdkgo.Pointer(os.Getenv("DECK_SECRET")),
		}),
	)

	res, err := s.Jobs.Submit(ctx, nil, &operations.PostJobsSubmitRequestBody2{
		JobCode: "FetchDocuments",
		Input: map[string]operations.InputUnion2{
			"access_token": operations.CreateInputUnion2Str(
				"access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
			),
			"key1": operations.CreateInputUnion2Str(
				"value1",
			),
			"someNumber": operations.CreateInputUnion2Number(
				123.45,
			),
			"someBoolean": operations.CreateInputUnion2Boolean(
				true,
			),
			"someArray": operations.CreateInputUnion2ArrayOfStr(
				[]string{
					"a",
					"b",
				},
			),
			"nestedObject": operations.CreateInputUnion2Boolean(
				true,
			),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.IJobResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Submit` function may return the following errors:

| Error Type                               | Status Code | Content Type               |
| ---------------------------------------- | ----------- | -------------------------- |
| apierrors.BadRequestJobResponseError     | 400         | application/json           |
| apierrors.BadRequestJobResponseError     | 400         | application/json+encrypted |
| apierrors.BadRequestJobResponseError     | 400         | text/json                  |
| apierrors.AlreadyRunningJobResponseError | 409         | application/json           |
| apierrors.AlreadyRunningJobResponseError | 409         | application/json+encrypted |
| apierrors.AlreadyRunningJobResponseError | 409         | text/json                  |
| apierrors.APIError                       | 4XX, 5XX    | \*/\*                      |

### Example

```go
package main

import (
	"context"
	"errors"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"github.com/buildwithdeck/sdk-go/models/apierrors"
	"github.com/buildwithdeck/sdk-go/models/components"
	"github.com/buildwithdeck/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(components.Security{
			ClientID: sdkgo.Pointer(os.Getenv("DECK_CLIENT_ID")),
			Secret:   sdkgo.Pointer(os.Getenv("DECK_SECRET")),
		}),
	)

	res, err := s.Jobs.Submit(ctx, nil, &operations.PostJobsSubmitRequestBody2{
		JobCode: "FetchDocuments",
		Input: map[string]operations.InputUnion2{
			"access_token": operations.CreateInputUnion2Str(
				"access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
			),
			"key1": operations.CreateInputUnion2Str(
				"value1",
			),
			"someNumber": operations.CreateInputUnion2Number(
				123.45,
			),
			"someBoolean": operations.CreateInputUnion2Boolean(
				true,
			),
			"someArray": operations.CreateInputUnion2ArrayOfStr(
				[]string{
					"a",
					"b",
				},
			),
			"nestedObject": operations.CreateInputUnion2Boolean(
				true,
			),
		},
	})
	if err != nil {

		var e *apierrors.BadRequestJobResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.BadRequestJobResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.BadRequestJobResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.AlreadyRunningJobResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.AlreadyRunningJobResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.AlreadyRunningJobResponseError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                           | Description      |
| --- | -------------------------------- | ---------------- |
| 0   | `https://sandbox.deck.co/api/v1` | Deck Sandbox API |
| 1   | `https://live.deck.co/api/v1`    | Deck API         |

#### Example

```go
package main

import (
	"context"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"github.com/buildwithdeck/sdk-go/models/components"
	"github.com/buildwithdeck/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithServerIndex(1),
		sdkgo.WithSecurity(components.Security{
			ClientID: sdkgo.Pointer(os.Getenv("DECK_CLIENT_ID")),
			Secret:   sdkgo.Pointer(os.Getenv("DECK_SECRET")),
		}),
	)

	res, err := s.Jobs.Submit(ctx, nil, &operations.PostJobsSubmitRequestBody2{
		JobCode: "FetchDocuments",
		Input: map[string]operations.InputUnion2{
			"access_token": operations.CreateInputUnion2Str(
				"access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
			),
			"key1": operations.CreateInputUnion2Str(
				"value1",
			),
			"someNumber": operations.CreateInputUnion2Number(
				123.45,
			),
			"someBoolean": operations.CreateInputUnion2Boolean(
				true,
			),
			"someArray": operations.CreateInputUnion2ArrayOfStr(
				[]string{
					"a",
					"b",
				},
			),
			"nestedObject": operations.CreateInputUnion2Boolean(
				true,
			),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.IJobResponse != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"github.com/buildwithdeck/sdk-go/models/components"
	"github.com/buildwithdeck/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithServerURL("https://live.deck.co/api/v1"),
		sdkgo.WithSecurity(components.Security{
			ClientID: sdkgo.Pointer(os.Getenv("DECK_CLIENT_ID")),
			Secret:   sdkgo.Pointer(os.Getenv("DECK_SECRET")),
		}),
	)

	res, err := s.Jobs.Submit(ctx, nil, &operations.PostJobsSubmitRequestBody2{
		JobCode: "FetchDocuments",
		Input: map[string]operations.InputUnion2{
			"access_token": operations.CreateInputUnion2Str(
				"access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
			),
			"key1": operations.CreateInputUnion2Str(
				"value1",
			),
			"someNumber": operations.CreateInputUnion2Number(
				123.45,
			),
			"someBoolean": operations.CreateInputUnion2Boolean(
				true,
			),
			"someArray": operations.CreateInputUnion2ArrayOfStr(
				[]string{
					"a",
					"b",
				},
			),
			"nestedObject": operations.CreateInputUnion2Boolean(
				true,
			),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.IJobResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/buildwithdeck/sdk-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdkgo.New(sdkgo.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/buildwithdeck/sdk-go&utm_campaign=go)
