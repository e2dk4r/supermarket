package cockroachdb

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type DbErrorService struct{}

func (des *DbErrorService) IsDuplicateError(err error) bool {
	var pgError *pgconn.PgError
	return errors.As(err, &pgError) && pgError.Code == "23505"
}

func (des *DbErrorService) IsNotFoundError(err error) bool {
	return err == pgx.ErrNoRows
}
