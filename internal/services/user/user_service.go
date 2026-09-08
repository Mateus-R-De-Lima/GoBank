package user

import (
	"context"
	"errors"
	"gobank/internal/store/pgstore"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserSerivce struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

var (
	ErroEmailOuUserNameJaExiste = errors.New("User_name ou email já existe!")
	ErrorCredenciasInvalidas    = errors.New("Credencias invalidas")
)

func NovoUserServices(pool *pgxpool.Pool) UserSerivce {
	return UserSerivce{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (us *UserSerivce) CriarUsuario(ctx context.Context, nome, email, password string) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return uuid.UUID{}, err
	}

	args := pgstore.CreateUserParams{
		UserName:     nome,
		Email:        email,
		PasswordHash: hash,
	}

	id, err := us.queries.CreateUser(ctx, args)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErroEmailOuUserNameJaExiste
		}

		return uuid.UUID{}, err
	}

	return id, nil
}
