# Jobs
(*Jobs*)

## Overview

Endpoints related to jobs.

### Available Operations

* [Submit](#submit) - Run a Job
* [GetJobsJobIDStatus](#getjobsjobidstatus) - Retrieve Job Run Status
* [GetJobsJobID](#getjobsjobid) - Retrieve Job Run Output and Status
* [GetJobsJobIDAll](#getjobsjobidall) - Retrieve all Job Information

## Submit

Runs a job with the specified parameters and returns a job ID for tracking its status.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/jobs/submit" method="post" path="/jobs/submit" -->
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
            ClientID: sdkgo.Pointer(os.Getenv("DECK_CLIENT_ID")),
            Secret: sdkgo.Pointer(os.Getenv("DECK_SECRET")),
        }),
    )

    res, err := s.Jobs.Submit(ctx, nil, nil, &operations.PostJobsSubmitRequestBody{
        JobCode: "FetchDocuments",
        Input: map[string]operations.InputUnion{
            "access_token": operations.CreateInputUnionStr(
                "access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
            ),
            "key1": operations.CreateInputUnionStr(
                "value1",
            ),
            "someNumber": operations.CreateInputUnionNumber(
                123.45,
            ),
            "someBoolean": operations.CreateInputUnionBoolean(
                true,
            ),
            "someArray": operations.CreateInputUnionArrayOfStr(
                []string{
                    "a",
                    "b",
                },
            ),
            "nestedObject": operations.CreateInputUnionBoolean(
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

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `xDeckCorrelationID`                                                                          | **string*                                                                                     | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `xDeckSandbox`                                                                                | **string*                                                                                     | :heavy_minus_sign:                                                                            | Can be used against the sandbox API with special values to test error use cases               |
| `requestBody`                                                                                 | [*operations.PostJobsSubmitRequestBody](../../models/operations/postjobssubmitrequestbody.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |

### Response

**[*operations.PostJobsSubmitResponse](../../models/operations/postjobssubmitresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.BadRequestJobResponseError     | 400                                      | application/json                         |
| apierrors.AlreadyRunningJobResponseError | 409                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## GetJobsJobIDStatus

Returns the current status for a specific job run.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/jobs/{jobId}/status" method="get" path="/jobs/{jobId}/status" -->
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

    res, err := s.Jobs.GetJobsJobIDStatus(ctx, "859ab723-d15b-4fa5-bde1-7c39c294108c")
    if err != nil {
        log.Fatal(err)
    }
    if res.JobStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `jobID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetJobsJobIDStatusResponse](../../models/operations/getjobsjobidstatusresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.BadRequestJobResponseError | 400, 404                             | application/json                     |
| apierrors.BadRequestJobResponseError | 400, 404                             | application/json+encrypted           |
| apierrors.BadRequestJobResponseError | 400, 404                             | text/json                            |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## GetJobsJobID

Returns information for the specific job, including its output and its status.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/jobs/{jobId}" method="get" path="/jobs/{jobId}" -->
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

    res, err := s.Jobs.GetJobsJobID(ctx, "6462a33a-1b55-4e48-89c9-c5d8fa5b7d36")
    if err != nil {
        log.Fatal(err)
    }
    if res.JobDetailsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `jobID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetJobsJobIDResponse](../../models/operations/getjobsjobidresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.BadRequestJobResponseError | 400, 404                             | application/json                     |
| apierrors.BadRequestJobResponseError | 400, 404                             | application/json+encrypted           |
| apierrors.BadRequestJobResponseError | 400, 404                             | text/json                            |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## GetJobsJobIDAll

Returns information for the specific job, including all play-back artifacts.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/jobs/{jobId}/all" method="get" path="/jobs/{jobId}/all" -->
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

    res, err := s.Jobs.GetJobsJobIDAll(ctx, "c1e3f7a0-60a4-4a56-87bf-6a310de9f405")
    if err != nil {
        log.Fatal(err)
    }
    if res.JobArtifactsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `jobID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetJobsJobIDAllResponse](../../models/operations/getjobsjobidallresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.BadRequestJobResponseError | 400, 404                             | application/json                     |
| apierrors.BadRequestJobResponseError | 400, 404                             | application/json+encrypted           |
| apierrors.BadRequestJobResponseError | 400, 404                             | text/json                            |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |