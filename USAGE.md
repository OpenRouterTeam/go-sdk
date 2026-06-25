<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	openrouter "github.com/OpenRouterTeam/go-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := openrouter.New(
		openrouter.WithSecurity(os.Getenv("OPENROUTER_API_KEY")),
	)

	res, err := s.Analytics.GetUserActivity(ctx, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->