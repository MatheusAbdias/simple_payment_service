package adapters

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	migrations "github.com/MatheusAbdias/simple_payment_service/cmd"
	"github.com/MatheusAbdias/simple_payment_service/domain/users"
	wallets "github.com/MatheusAbdias/simple_payment_service/domain/wallets"
	"github.com/MatheusAbdias/simple_payment_service/pkg/postgres"
	utils "github.com/MatheusAbdias/simple_payment_service/pkg/utils"
	"github.com/stretchr/testify/require"
)

func setupDataBase(ctx context.Context) (*sql.DB, *postgres.PostgresContainer) {
	container, err := postgres.NewPostgresContainer(ctx)
	if err != nil {
		log.Fatalf("error creating container: %s", err)
	}

	db, err := postgres.NewDB(container.ConnectionString)
	if err != nil {
		log.Fatalf("error creating database: %s", err)
	}

	healthChecker := postgres.NewHealthChecker(db)
	err = healthChecker.CheckHealth(ctx)
	if err != nil {
		log.Fatalf("error checking health: %s", err)
	}

	os.Setenv("DATABASE_URL", container.ConnectionString)
	dir, _ := os.Getwd()
	os.Setenv("BASE_DIR", dir)
	migrations.Migrate()

	return db, container
}

func tearDownDataBase(ctx context.Context, db *sql.DB, container *postgres.PostgresContainer) {
	db.Close()
	container.Terminate(ctx)

}

func TestCreateUserWithWallet(t *testing.T) {
	ctx := context.Background()
	db, container := setupDataBase(ctx)
	defer tearDownDataBase(ctx, db, container)

	repo := NewPostgresRepository(db)

	userID := utils.NewUUID()
	wallet := &wallets.Wallet{
		Id:      utils.NewUUID(),
		OwnerID: userID,
	}
	user := &users.User{
		Id:       userID,
		FullName: "Jon Doe",
		Email:    "test@test.com",
		Document: "21472605000155",
		Wallet:   wallet,
	}

	_, err := repo.CreateUserWithWallet(ctx, user)
	require.Nil(t, err)

	find := repo.FindUserByEmailOrDocument(ctx, user.Email, user.Document)
	require.True(t, find)
}

func TestFailCreateUserWithEmailIsAlreadyRegister(t *testing.T) {
	ctx := context.Background()
	db, container := setupDataBase(ctx)
	defer tearDownDataBase(ctx, db, container)

	repo := NewPostgresRepository(db)

	userID := utils.NewUUID()
	wallet := &wallets.Wallet{
		Id:      utils.NewUUID(),
		OwnerID: userID,
	}
	user := &users.User{
		Id:       userID,
		FullName: "Jon Doe",
		Email:    "jon@email.com",
		Document: "21472605000155",
		Wallet:   wallet,
	}

	_, err := repo.CreateUserWithWallet(ctx, user)
	require.Nil(t, err)

	otherUserID := utils.NewUUID()
	otherWallet := &wallets.Wallet{
		Id:      utils.NewUUID(),
		OwnerID: otherUserID,
	}
	invalidUser := &users.User{
		Id:       otherUserID,
		FullName: "Mark Doe",
		Email:    "JON@email.com",
		Document: "57271936050",
		Wallet:   otherWallet,
	}

	_, err = repo.CreateUserWithWallet(ctx, invalidUser)
	require.NotNil(t, err)
}

func TestFailCreateUserWithDocumentIsAlreadyRegister(t *testing.T) {
	ctx := context.Background()
	db, container := setupDataBase(ctx)
	defer tearDownDataBase(ctx, db, container)

	repo := NewPostgresRepository(db)

	userID := utils.NewUUID()
	wallet := &wallets.Wallet{
		Id:      utils.NewUUID(),
		OwnerID: userID,
	}
	user := &users.User{
		Id:       userID,
		FullName: "Jon Doe",
		Email:    "jon@email.com",
		Document: "21472605000155",
		Wallet:   wallet,
	}

	_, err := repo.CreateUserWithWallet(ctx, user)
	require.Nil(t, err)

	otherUserID := utils.NewUUID()
	otherWallet := &wallets.Wallet{
		Id:      utils.NewUUID(),
		OwnerID: otherUserID,
	}
	invalidUser := &users.User{
		Id:       otherUserID,
		FullName: "Mark Doe",
		Email:    "mark@email.com",
		Document: "21472605000155",
		Wallet:   otherWallet,
	}

	_, err = repo.CreateUserWithWallet(ctx, invalidUser)
	require.NotNil(t, err)
}
