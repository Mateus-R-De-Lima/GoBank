package api

import (
	"errors"
	pessoaFisicaService "gobank/internal/services/pessoa_fisica"
	pessoafisica "gobank/internal/usecases/pessoa_fisica"
	"gobank/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (a *Api) handleCriarContaPessoaFisica(w http.ResponseWriter, r *http.Request) {
	data, problemas, err := utils.DecodificarJson[pessoafisica.CriarPessoaFisicaRequest](r)
	if err != nil {
		_ = utils.CodificarJson(w, r, http.StatusUnprocessableEntity, problemas)
		return
	}

	userID, err := a.authenticatedUserID(r)
	if err != nil {
		writeErrorJSON(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	novaPessoaFisicaId, err := a.PessoaFisicaService.CriarNovaPessoaFisica(
		r.Context(),
		data.NomeCompleto,
		data.Categoria,
		data.Celular,
		data.Email,
		data.Idade,
		data.Saldo,
		data.RendaMensal,
		userID,
	)
	if err != nil {
		if errors.Is(err, pessoaFisicaService.ErroEmailJaExiste) {
			writeErrorJSON(w, r, http.StatusBadRequest, pessoaFisicaService.ErroEmailJaExiste.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao criar a pessoa física.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusCreated, map[string]any{
		"id":       novaPessoaFisicaId,
		"mensagem": "Pessoa física criada com sucesso!",
	})
}

func (a *Api) handlerGetListaContaPessoaFisicaPorUserId(w http.ResponseWriter, r *http.Request) {
	userID, err := a.authenticatedUserID(r)
	if err != nil {
		writeErrorJSON(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	listaPessoaFisica, err := a.PessoaFisicaService.BuscarPessoaFisicaByUserId(r.Context(), userID)
	if err != nil {
		if errors.Is(err, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao buscar a lista de pessoa física.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"data": listaPessoaFisica,
	})
}

func (a *Api) handlerGetContaPessoaFisicaPorId(w http.ResponseWriter, r *http.Request) {
	pessoaIdParam := chi.URLParam(r, "conta_id")

	pessoaId, err := uuid.Parse(pessoaIdParam)
	if err != nil {
		writeErrorJSON(w, r, http.StatusBadRequest, "Pessoa Id inválido!")
		return
	}

	responseBuscaPessoaFisica, err := a.PessoaFisicaService.BuscarPessoaFisicaById(r.Context(), pessoaId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao buscar a pessoa física.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"data": responseBuscaPessoaFisica,
	})
}

func (a *Api) handlerPatchSaldoContaPessoaFisicaPorId(w http.ResponseWriter, r *http.Request) {
	data, problemas, err := utils.DecodificarJson[pessoafisica.AtualizarSaldoPessoaFisicaRequest](r)
	if err != nil {
		_ = utils.CodificarJson(w, r, http.StatusUnprocessableEntity, problemas)
		return
	}

	pessoaID, err := uuid.Parse(chi.URLParam(r, "conta_id"))
	if err != nil {
		writeErrorJSON(w, r, http.StatusBadRequest, "Pessoa Id inválido!")
		return
	}

	response, err := a.PessoaFisicaService.AtualizarSaldoPessoaFisica(r.Context(), pessoaID, data.Saldo)
	if err != nil {
		if errors.Is(err, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao atualizar o saldo da pessoa física.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"data": response,
	})
}

func (a *Api) handlerUpdateContaPessoaFisicaPorId(w http.ResponseWriter, r *http.Request) {
	data, problemas, err := utils.DecodificarJson[pessoafisica.AtualizarPessoaFisicaRequest](r)
	if err != nil {
		_ = utils.CodificarJson(w, r, http.StatusUnprocessableEntity, problemas)
		return
	}

	pessoaID, err := uuid.Parse(chi.URLParam(r, "conta_id"))
	if err != nil {
		writeErrorJSON(w, r, http.StatusBadRequest, "Pessoa Id inválido!")
		return
	}

	response, err := a.PessoaFisicaService.AtualizarPessoaFisica(
		r.Context(),
		pessoaID,
		data.RendaMensal,
		data.Idade,
		data.NomeCompleto,
		data.Celular,
		data.Email,
		data.Categoria,
		data.Saldo,
	)
	if err != nil {
		if errors.Is(err, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado.Error())
			return
		}

		if errors.Is(err, pessoaFisicaService.ErroEmailJaExiste) {
			writeErrorJSON(w, r, http.StatusBadRequest, pessoaFisicaService.ErroEmailJaExiste.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao atualizar a pessoa física.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"data": response,
	})
}

func (a *Api) handlerDeleteContaPessoaFisicaPorId(w http.ResponseWriter, r *http.Request) {
	pessoaID, err := uuid.Parse(chi.URLParam(r, "conta_id"))
	if err != nil {
		writeErrorJSON(w, r, http.StatusBadRequest, "Pessoa Id inválido!")
		return
	}

	if err := a.PessoaFisicaService.ExcluirPessoaFisica(r.Context(), pessoaID); err != nil {
		if errors.Is(err, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaFisicaService.ErrorPessoaFisicaNaoEncontrado.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao excluir a pessoa física.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"mensagem": "Pessoa física removida com sucesso!",
	})
}

func (a *Api) authenticatedUserID(r *http.Request) (uuid.UUID, error) {
	userIDValue := a.Sessions.Get(r.Context(), "AuthenticatedUserId")
	if userIDValue == nil {
		return uuid.Nil, errors.New("usuário não autenticado")
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("id do usuário inválido")
	}

	return userID, nil
}

func writeErrorJSON(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	_ = utils.CodificarJson(w, r, statusCode, map[string]any{
		"error": message,
	})
}
