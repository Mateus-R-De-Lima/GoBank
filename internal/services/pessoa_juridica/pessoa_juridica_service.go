package pessoajuridica

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

type PessoaJuridicaService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

type PessoaJuridicaResponse struct {
	Id    string  `json:"id"`
	Nome  string  `json:"nome"`
	Email string  `json:"email"`
	Saldo float64 `json:"saldo"`
}

var (
	ErroEmailJaExiste                = errors.New("Email já existe e esta vinculado a outro usuario!")
	ErrorPessoaJuridicaNaoEncontrado = errors.New("Não foi encontrado nenhuma pessoa juridica com os dados informados!")
)

func NovaPessoaJuridicaServices(pool *pgxpool.Pool) PessoaJuridicaService {
	return PessoaJuridicaService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (pjs *PessoaJuridicaService) CriarNovaPessoaJuridica(ctx context.Context, nomeFantasia, categoria, celular, emailCorporativo string, idade int32, saldo, faturamento float64, userId uuid.UUID) (uuid.UUID, error) {
	if existe, err := pjs.emailJaExiste(ctx, emailCorporativo, nil); err != nil {
		return uuid.UUID{}, err
	} else if existe {
		return uuid.UUID{}, ErroEmailJaExiste
	}

	arg := pgstore.CreatePessoaJuridicaParams{
		Faturamento:      faturamento,
		Idade:            idade,
		NomeFantasia:     nomeFantasia,
		Celular:          celular,
		EmailCorporativo: emailCorporativo,
		Categoria:        categoria,
		Saldo:            saldo,
		UserID:           userId,
	}

	novaPessoaJuridica, err := pjs.queries.CreatePessoaJuridica(ctx, arg)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErroEmailJaExiste
		}
		return uuid.UUID{}, err
	}

	return novaPessoaJuridica.ID, nil
}

func (pjs *PessoaJuridicaService) BuscarPessoaJuridicaByUserId(ctx context.Context, userId uuid.UUID) ([]PessoaJuridicaResponse, error) {
	pessoasJuridicas, err := pjs.queries.GetPessoaJuridicaByUserId(ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrorPessoaJuridicaNaoEncontrado
		}
		return nil, err
	}

	if len(pessoasJuridicas) == 0 {
		return nil, ErrorPessoaJuridicaNaoEncontrado
	}

	response := make([]PessoaJuridicaResponse, 0, len(pessoasJuridicas))
	for _, pessoa := range pessoasJuridicas {
		response = append(response, PessoaJuridicaResponse{
			Id:    pessoa.ID.String(),
			Nome:  pessoa.NomeFantasia,
			Saldo: pessoa.Saldo,
			Email: pessoa.EmailCorporativo,
		})
	}

	return response, nil
}

func (pjs *PessoaJuridicaService) BuscarPessoaJuridicaById(ctx context.Context, pessoaJuridicaId uuid.UUID) (PessoaJuridicaResponse, error) {
	pessoaJuridica, err := pjs.queries.GetPessoaJuridica(ctx, pessoaJuridicaId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return PessoaJuridicaResponse{}, ErrorPessoaJuridicaNaoEncontrado
		}
		return PessoaJuridicaResponse{}, err
	}

	return mapPessoaJuridicaResponse(pessoaJuridica), nil
}

func (pjs *PessoaJuridicaService) AtualizarPessoaJuridica(ctx context.Context, pessoaJuridicaId uuid.UUID, faturamento float64, idade int32, nomeFantasia, celular, emailCorporativo, categoria string, saldo float64) (PessoaJuridicaResponse, error) {
	if existe, err := pjs.emailJaExiste(ctx, emailCorporativo, &pessoaJuridicaId); err != nil {
		return PessoaJuridicaResponse{}, err
	} else if existe {
		return PessoaJuridicaResponse{}, ErroEmailJaExiste
	}

	pessoaAtualizada, err := pjs.queries.UpdatePessoaJuridica(ctx, pgstore.UpdatePessoaJuridicaParams{
		ID:               pessoaJuridicaId,
		Faturamento:      faturamento,
		Idade:            idade,
		NomeFantasia:     nomeFantasia,
		Celular:          celular,
		EmailCorporativo: emailCorporativo,
		Categoria:        categoria,
		Saldo:            saldo,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return PessoaJuridicaResponse{}, ErrorPessoaJuridicaNaoEncontrado
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return PessoaJuridicaResponse{}, ErroEmailJaExiste
		}

		return PessoaJuridicaResponse{}, err
	}

	return mapPessoaJuridicaResponse(pessoaAtualizada), nil
}

func (pjs *PessoaJuridicaService) AtualizarSaldoPessoaJuridica(ctx context.Context, pessoaJuridicaId uuid.UUID, saldo float64) (PessoaJuridicaResponse, error) {
	pessoaAtualizada, err := pjs.queries.UpdateSaldoPessoaJuridica(ctx, pgstore.UpdateSaldoPessoaJuridicaParams{
		ID:    pessoaJuridicaId,
		Saldo: saldo,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return PessoaJuridicaResponse{}, ErrorPessoaJuridicaNaoEncontrado
		}
		return PessoaJuridicaResponse{}, err
	}

	return mapPessoaJuridicaResponse(pessoaAtualizada), nil
}

func (pjs *PessoaJuridicaService) ExcluirPessoaJuridica(ctx context.Context, pessoaJuridicaId uuid.UUID) error {
	if _, err := pjs.queries.GetPessoaJuridica(ctx, pessoaJuridicaId); err != nil {
		if err == pgx.ErrNoRows {
			return ErrorPessoaJuridicaNaoEncontrado
		}
		return err
	}

	if err := pjs.queries.DeletePessoaJuridica(ctx, pessoaJuridicaId); err != nil {
		return err
	}

	return nil
}

func mapPessoaJuridicaResponse(pessoaJuridica pgstore.PessoaJuridica) PessoaJuridicaResponse {
	return PessoaJuridicaResponse{
		Id:    pessoaJuridica.ID.String(),
		Nome:  pessoaJuridica.NomeFantasia,
		Saldo: pessoaJuridica.Saldo,
		Email: pessoaJuridica.EmailCorporativo,
	}
}

func (pjs *PessoaJuridicaService) emailJaExiste(ctx context.Context, emailCorporativo string, pessoaID *uuid.UUID) (bool, error) {
	emailCorporativo = strings.TrimSpace(emailCorporativo)
	if emailCorporativo == "" {
		return false, nil
	}

	pessoa, err := pjs.queries.GetPessoaJuridicaByEmail(ctx, emailCorporativo)
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
