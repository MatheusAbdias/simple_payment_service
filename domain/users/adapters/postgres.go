package adapters

import (
	"context"
	"database/sql"

	"github.com/MatheusAbdias/simple_payment_service/domain/users"
)

type PostgresRepository struct {
	Conn *sql.DB
}

func NewPostgresRepository(conn *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: conn,
	}
}

func (r *PostgresRepository) CreateUserWithWallet(
	ctx context.Context,
	user *users.User,
) (*users.User, error) {
	tx, err := r.Conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	userQuery := `INSERT INTO users (id, full_name, email, document) VALUES ($1,$2,$3,$4)`
	_, err = tx.ExecContext(ctx, userQuery, user.Id, user.FullName, user.Email, user.Document)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	walletQuery := `INSERT INTO wallets (id, owner_id, amount) VALUES ($1, $2, $3)`
	_, err = tx.ExecContext(
		ctx,
		walletQuery,
		user.Wallet.Id,
		user.Wallet.OwnerID,
		user.Wallet.Amount,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresRepository) FindUserByEmailOrDocument(
	ctx context.Context,
	email string,
	document string,
) bool {
	query := `SELECT 1 FROM users WHERE email = $1 OR document = $2`
	row := r.Conn.QueryRowContext(ctx, query, email, document)

	return row.Err() == nil
}
