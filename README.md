# quickdb

## Usage

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/sqlc-dev/quickdb"
	pb "github.com/sqlc-dev/quickdb/v1"
)

func main() {
	ctx := context.Background()
	projectID := os.Getenv("SQLC_PROJECT_ID")
	authToken := os.Getenv("SQLC_AUTH_TOKEN")
	client, err := quickdb.NewClient(projectID, authToken)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.CreateEphemeralDatabase(ctx, &pb.CreateEphemeralDatabaseRequest{
		Engine: "postgresql",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Uri)

	_, err = client.DropEphemeralDatabase(ctx, &pb.DropEphemeralDatabaseRequest{
		DatabaseId: resp.DatabaseId,
	})
	if err != nil {
		log.Fatal(err)
	}
}
```