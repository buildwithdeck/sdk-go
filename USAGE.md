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
			ClientID: sdkgo.String(os.Getenv("DECK_CLIENT_ID")),
			Secret:   sdkgo.String(os.Getenv("DECK_SECRET")),
		}),
	)

	res, err := s.Jobs.Submit(ctx, nil, &operations.PostJobsSubmitRequestBody2{
		JobCode: "FetchDocuments",
		Input: map[string]string{
			"access_token": "access-development-6599f8dd-1a1c-4586-39d1-08ddb97283f7",
			"key1":         "value1",
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
<!-- End SDK Example Usage [usage] -->