<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sdkgo "github.com/buildwithdeck/sdk-go"
	"github.com/buildwithdeck/sdk-go/models/components"
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

	res, err := s.Authentication.PostJobsJobGUIDCancel(ctx, "839ff252-653e-4088-92e2-4bcc8c4ebe83")
	if err != nil {
		log.Fatal(err)
	}
	if res.JobCancelResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->