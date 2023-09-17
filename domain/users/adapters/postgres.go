package adapters

import (
	"context"
	"database/sql"

	domain "github.com/MatheusAbdias/simple_payment_service/domain/users"
)

type PostgresRepo struct {
	Conn *sql.DB
}

func NewPostgresRepository(conn *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		Conn: conn,
	}
}

func (postgresRepository *PostgresRepo) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, full_name, email, document) VALUES ($1, $2, $3, $4)`
	_, err := postgresRepository.Conn.ExecContext(ctx, query, user.Id, user.FullName, user.Email, user.Document)
	if err != nil {
		return err
	}
	return nil
}

func (postgresRepository *PostgresRepo) FindUserByEmail(ctx context.Context, email string) bool {
	query := `SELECT 1 FROM users WHERE email = $1`
	row := postgresRepository.Conn.QueryRowContext(ctx, query, email)

	return row.Err() == nil
}

func (postgresRepository *PostgresRepo) FindUserByDocument(ctx context.Context, document string) bool {
	query := `SELECT 1 FROM users WHERE document = $1`
	row := postgresRepository.Conn.QueryRowContext(ctx, query, document)

	return row.Err() == nil
}
