package pessoafisica

import (
	"context"
	"errors"
	"gobank/internal/store/pgstore"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PessoaFisicaService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

type PessoaResponse struct {
	Id    string  `json:"id"`
	Nome  string  `json:"nome"`
	Email string  `json:"email"`
	Saldo float64 `json:"saldo"`
}

var (
	ErroEmailJaExiste              = errors.New("Email já existe e esta vinculado a outro usuario!")
	ErrorPessoaFisicaNaoEncontrado = errors.New("Não foi encontrado nenhuma pessoa fisica com os dados informados!")
)

func NovaPessoaFisicaServices(pool *pgxpool.Pool) PessoaFisicaService {
	return PessoaFisicaService{
		pool:    pool,
		queries: pgstore.New(pool),
	}

}

func (pfs *PessoaFisicaService) CriarNovaPessoaFisica(ctx context.Context, nomeCompleto, categoria, celular, email string, idade int32, saldo, rendaMensal float64, userId uuid.UUID) (uuid.UUID, error) {
	if existe, err := pfs.emailJaExiste(ctx, email, nil); err != nil {
		return uuid.UUID{}, err
	} else if existe {
		return uuid.UUID{}, ErroEmailJaExiste
	}

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

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {

			return uuid.UUID{}, ErroEmailJaExiste
		}

		return uuid.UUID{}, err

	}

	return novaPessoaFisica.ID, nil
}

func (pfs *PessoaFisicaService) BuscarPessoaFisicaByUserId(ctx context.Context, userId uuid.UUID) ([]PessoaResponse, error) {

	pessoaFisica, err := pfs.queries.GetPessoaFisicaByUserId(ctx, userId)

	if err != nil {

		if err == pgx.ErrNoRows {
			return nil, ErrorPessoaFisicaNaoEncontrado
		}

		return nil, err
	}

	if len(pessoaFisica) == 0 {
		return nil, ErrorPessoaFisicaNaoEncontrado
	}

	response := make([]PessoaResponse, 0, len(pessoaFisica))

	for _, pessoa := range pessoaFisica {
		response = append(response, PessoaResponse{
			Id:    pessoa.ID.String(),
			Nome:  pessoa.NomeCompleto,
			Saldo: pessoa.Saldo,
			Email: pessoa.Email,
		})
	}

	return response, nil

}

func (pfs *PessoaFisicaService) BuscarPessoaFisicaById(ctx context.Context, pessoaFisicaId uuid.UUID) (PessoaResponse, error) {
	pessoaFisica, err := pfs.queries.GetPessoaFisica(ctx, pessoaFisicaId)

	if err != nil {
		if err == pgx.ErrNoRows {
			return PessoaResponse{}, ErrorPessoaFisicaNaoEncontrado
		}

		return PessoaResponse{}, err
	}

	return mapPessoaFisicaResponse(pessoaFisica), nil

}

func (pfs *PessoaFisicaService) AtualizarPessoaFisica(ctx context.Context, pessoaFisicaId uuid.UUID, rendaMensal float64, idade int32, nomeCompleto, celular, email, categoria string, saldo float64) (PessoaResponse, error) {
	if existe, err := pfs.emailJaExiste(ctx, email, &pessoaFisicaId); err != nil {
		return PessoaResponse{}, err
	} else if existe {
		return PessoaResponse{}, ErroEmailJaExiste
	}

	pessoaAtualizada, err := pfs.queries.UpdatePessoaFisica(ctx, pgstore.UpdatePessoaFisicaParams{
		ID:           pessoaFisicaId,
		RendaMensal:  rendaMensal,
		Idade:        idade,
		NomeCompleto: nomeCompleto,
		Celular:      celular,
		Email:        email,
		Categoria:    categoria,
		Saldo:        saldo,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return PessoaResponse{}, ErrorPessoaFisicaNaoEncontrado
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return PessoaResponse{}, ErroEmailJaExiste
		}

		return PessoaResponse{}, err
	}

	return mapPessoaFisicaResponse(pessoaAtualizada), nil
}

func (pfs *PessoaFisicaService) AtualizarSaldoPessoaFisica(ctx context.Context, pessoaFisicaId uuid.UUID, saldo float64) (PessoaResponse, error) {
	pessoaAtualizada, err := pfs.queries.UpdateSaldoPessoaFisica(ctx, pgstore.UpdateSaldoPessoaFisicaParams{
		ID:    pessoaFisicaId,
		Saldo: saldo,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return PessoaResponse{}, ErrorPessoaFisicaNaoEncontrado
		}

		return PessoaResponse{}, err
	}

	return mapPessoaFisicaResponse(pessoaAtualizada), nil
}

func (pfs *PessoaFisicaService) ExcluirPessoaFisica(ctx context.Context, pessoaFisicaId uuid.UUID) error {
	if _, err := pfs.queries.GetPessoaFisica(ctx, pessoaFisicaId); err != nil {
		if err == pgx.ErrNoRows {
			return ErrorPessoaFisicaNaoEncontrado
		}
		return err
	}

	if err := pfs.queries.DeletePessoaFisica(ctx, pessoaFisicaId); err != nil {
		return err
	}

	return nil
}

func mapPessoaFisicaResponse(pessoaFisica pgstore.PessoaFisica) PessoaResponse {
	return PessoaResponse{
		Id:    pessoaFisica.ID.String(),
		Nome:  pessoaFisica.NomeCompleto,
		Saldo: pessoaFisica.Saldo,
		Email: pessoaFisica.Email,
	}
}

func (pfs *PessoaFisicaService) emailJaExiste(ctx context.Context, email string, pessoaID *uuid.UUID) (bool, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return false, nil
	}

	pessoa, err := pfs.queries.GetPessoaFisicaByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	if pessoaID != nil && pessoa.ID == *pessoaID {
		return false, nil
	}

	return true, nil
}
