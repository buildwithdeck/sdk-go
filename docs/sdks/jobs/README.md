# Jobs
(*Jobs*)

## Overview

Endpoints related to jobs.

### Available Operations

* [Submit](#submit) - Send your job requests
* [AnswerMFA](#answermfa) - Provide MFA code
* [GetDocumentFile](#getdocumentfile) - Get raw file

## Submit

Provide a job code along with its input parameters to execute it

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/buildwithdeck/sdk-go/models/components"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"github.com/buildwithdeck/sdk-go/models/operations"
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

    res, err := s.Jobs.Submit(ctx, &operations.PostJobsSubmitRequest{
        JobCode: "FetchDocuments",
        Input: map[string]string{
            "access_token": "access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
            "key1": "value1",
            "someProperty": "someValue",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.JobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PostJobsSubmitRequest](../../models/operations/postjobssubmitrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostJobsSubmitResponse](../../models/operations/postjobssubmitresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.BadRequestJobResponseError     | 400                                      | application/json                         |
| apierrors.BadRequestJobResponseError     | 400                                      | text/json                                |
| apierrors.AlreadyRunningJobResponseError | 409                                      | application/json                         |
| apierrors.AlreadyRunningJobResponseError | 409                                      | text/json                                |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## AnswerMFA

Call this endpoint to send your MFA code

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

    res, err := s.Jobs.AnswerMFA(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.MfaAnswerRequest](../../models/components/mfaanswerrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.PostJobsMfaAnswerResponse](../../models/operations/postjobsmfaanswerresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ErrorMessageResponse | 400                            | application/json               |
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## GetDocumentFile

Returns the raw file for the document with the provided document ID

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

    res, err := s.Jobs.GetDocumentFile(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.RawDocumentRequest](../../models/components/rawdocumentrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.PostJobsDocumentsFileResponse](../../models/operations/postjobsdocumentsfileresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ErrorMessageResponse | 400                            | application/json               |
| apierrors.ErrorMessageResponse | 400                            | text/json                      |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |