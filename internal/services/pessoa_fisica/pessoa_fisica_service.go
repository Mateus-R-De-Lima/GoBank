package pessoafisica

import (
	"context"
	"errors"
	"fmt"
	"gobank/internal/store/pgstore"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PessoaFisicaService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

type PessoaResponse struct {
	id   string
	nome string
}

var (
	ErroEmailJaExiste = errors.New("Email já existe e esta vinculado a outro usuario!")
)

func NovaPessoaFisicaServices(pool *pgxpool.Pool) PessoaFisicaService {
	return PessoaFisicaService{
		pool:    pool,
		queries: pgstore.New(pool),
	}

}

func (pfs *PessoaFisicaService) CriarNovaPessoaFisica(ctx context.Context, nomeCompleto, categoria, celular, email string, idade int32, saldo, rendaMensal float64, userId uuid.UUID) (uuid.UUID, error) {

	arg := pgstore.CreatePessoaFisicaParams{
		RendaMensal:  rendaMensal,
		Idade:        idade,
		NomeCompleto: nomeCompleto,
		Celular:      celular,
		Email:        email,
		Categoria:    categoria,
		Saldo:        saldo,
		UserID:       userId,
	}

	novaPessoaFisica, err := pfs.queries.CreatePessoaFisica(ctx, arg)

	if err != nil {
		var pgErr *pgconn.PgError
		fmt.Println("PgEror : ", pgErr)
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			fmt.Println("Aqui esta o erro", err.Error(), err)
			return uuid.UUID{}, ErroEmailJaExiste
		}

		return uuid.UUID{}, err

	}

	return novaPessoaFisica.ID, nil
}
