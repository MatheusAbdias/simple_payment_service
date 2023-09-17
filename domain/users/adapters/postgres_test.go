package adapters

import (
	"context"
	"log"
	"os"
	"testing"

	migrations "github.com/MatheusAbdias/simple_payment_service/cmd"
	"github.com/MatheusAbdias/simple_payment_service/domain/users"
	"github.com/MatheusAbdias/simple_payment_service/pkg/postgres"
	utils "github.com/MatheusAbdias/simple_payment_service/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestPostgresRepository(t *testing.T) {
	ctx := context.Background()

	container, err := postgres.NewPostgresContainer(ctx)
	if err != nil {
		log.Fatalf("error creating container: %s", err)
	}
	defer container.Terminate(ctx)

	db, err := postgres.NewDB(container.ConnectionString)
	if err != nil {
		log.Fatalf("error creating database: %s", err)
	}
	defer db.Close()

	healthChecker := postgres.NewHealthChecker(db)
	err = healthChecker.CheckHealth(ctx)
	if err != nil {
		log.Fatalf("error checking health: %s", err)
	}

	os.Setenv("DATABASE_URL", container.ConnectionString)
	dir, _ := os.Getwd()
	os.Setenv("BASE_DIR", dir)
	migrations.Migrate()

	repo := NewPostgresRepository(db)
	user := &users.User{
		Id:       utils.NewUUID(),
		FullName: "Jon Doe",
		Email:    "test@test.com",
		Document: "21472605000155",
	}
	err = repo.CreateUser(ctx, user)
	require.Nil(t, err)

	find := repo.FindUserByEmail(ctx, user.Email)
	require.True(t, find)

	find = repo.FindUserByDocument(ctx, user.Document)
	require.True(t, find)
}
