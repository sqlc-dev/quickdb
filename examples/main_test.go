package example

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/sqlc-dev/quickdb"
	pb "github.com/sqlc-dev/quickdb/v1"
)

func createDatabase(ctx context.Context, path string) (*pb.CreateEphemeralDatabaseResponse, error) {
	projectID := os.Getenv("SQLC_PROJECT_ID")
	authToken := os.Getenv("SQLC_AUTH_TOKEN")
	if projectID == "" || authToken == "" {
		return nil, fmt.Errorf("missing project id and auth token")
	}

	client, err := quickdb.NewClient(projectID, authToken)
	if err != nil {
		return nil, fmt.Errorf("new client: %w", err)
	}

	var migrations []string
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("new client: %w", err)
	}
	for _, f := range files {
		contents, err := os.ReadFile(filepath.Join(path, f.Name()))
		if err != nil {
			return nil, fmt.Errorf("read file: %s", err)
		}
		migrations = append(migrations, string(contents))
	}

	resp, err := client.CreateEphemeralDatabase(ctx, &pb.CreateEphemeralDatabaseRequest{
		Engine:     "postgresql",
		Region:     quickdb.GetClosestRegion(),
		Migrations: migrations,
	})
	if err != nil {
		return nil, fmt.Errorf("create db: %w", err)
	}

	return resp, nil

	// cleanup := func() {
	// 	client.DropEphemeralDatabase(ctx, &pb.DropEphemeralDatabaseRequest{
	// 		DatabaseId: resp.DatabaseId,
	// 	})
	// }
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	resp, err := createDatabase(ctx, "path/to/migrations")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Uri)

	os.Exit(m.Run())
}
