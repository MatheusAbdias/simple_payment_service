package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectingWithPostgres(t *testing.T) {
	ctx := context.Background()

	container, err := NewPostgresContainer(ctx)

	defer container.Terminate(ctx)

	require.Nil(t, err, "error should be nil")

}
