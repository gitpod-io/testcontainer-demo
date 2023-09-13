package trial

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
)

func TestDocker(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			t.Parallel()

			req := testcontainers.GenericContainerRequest{
				ContainerRequest: testcontainers.ContainerRequest{
					Image: "redis:latest",
				},
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
			defer cancel()
			container, err := testcontainers.GenericContainer(ctx, req)
			if err != nil {
				t.Fatalf("Failed to create container: %v", err)
			}

			defer container.Terminate(ctx)

			err = container.Start(ctx)
			if err != nil {
				t.Fatalf("Failed to start container: %v", err)
			}
		})
	}
}
