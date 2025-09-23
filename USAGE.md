<!-- Start SDK Example Usage [usage] -->
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