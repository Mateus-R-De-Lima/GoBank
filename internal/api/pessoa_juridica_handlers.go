package api

import (
	"errors"
	pessoaJuridicaService "gobank/internal/services/pessoa_juridica"
	pessoajuridica "gobank/internal/usecases/pessoa_juridica"
	"gobank/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (a *Api) handleCriarContaPessoaJuridica(w http.ResponseWriter, r *http.Request) {
	data, problemas, err := utils.DecodificarJson[pessoajuridica.CriarPessoaJuridicaRequest](r)
	if err != nil {
		_ = utils.CodificarJson(w, r, http.StatusUnprocessableEntity, problemas)
		return
	}

	userID, err := a.authenticatedUserID(r)
	if err != nil {
		writeErrorJSON(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	novaPessoaJuridicaId, err := a.PessoaJuridicaService.CriarNovaPessoaJuridica(
		r.Context(),
		data.NomeFantasia,
		data.Categoria,
		data.Celular,
		data.EmailCorporativo,
		data.Idade,
		data.Saldo,
		data.Faturamento,
		userID,
	)
	if err != nil {
		if errors.Is(err, pessoaJuridicaService.ErroEmailJaExiste) {
			writeErrorJSON(w, r, http.StatusBadRequest, pessoaJuridicaService.ErroEmailJaExiste.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao criar a pessoa jurídica.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusCreated, map[string]any{
		"id":       novaPessoaJuridicaId,
		"mensagem": "Pessoa jurídica criada com sucesso!",
	})
}

func (a *Api) handlerGetListaContaPessoaJuridicaPorUserId(w http.ResponseWriter, r *http.Request) {
	userID, err := a.authenticatedUserID(r)
	if err != nil {
		writeErrorJSON(w, r, http.StatusUnauthorized, err.Error())
		return
	}

	listaPessoaJuridica, err := a.PessoaJuridicaService.BuscarPessoaJuridicaByUserId(r.Context(), userID)
	if err != nil {
		if errors.Is(err, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao buscar a lista de pessoa jurídica.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"data": listaPessoaJuridica,
	})
}

func (a *Api) handlerGetContaPessoaJuridicaPorId(w http.ResponseWriter, r *http.Request) {
	pessoaIdParam := chi.URLParam(r, "conta_id")

	pessoaId, err := uuid.Parse(pessoaIdParam)
	if err != nil {
		writeErrorJSON(w, r, http.StatusBadRequest, "Pessoa Id inválido!")
		return
	}

	responseBuscaPessoaJuridica, err := a.PessoaJuridicaService.BuscarPessoaJuridicaById(r.Context(), pessoaId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao buscar a pessoa jurídica.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"data": responseBuscaPessoaJuridica,
	})
}

func (a *Api) handlerPatchSaldoContaPessoaJuridicaPorId(w http.ResponseWriter, r *http.Request) {
	data, problemas, err := utils.DecodificarJson[pessoajuridica.AtualizarSaldoPessoaJuridicaRequest](r)
	if err != nil {
		_ = utils.CodificarJson(w, r, http.StatusUnprocessableEntity, problemas)
		return
	}

	pessoaID, err := uuid.Parse(chi.URLParam(r, "conta_id"))
	if err != nil {
		writeErrorJSON(w, r, http.StatusBadRequest, "Pessoa Id inválido!")
		return
	}

	response, err := a.PessoaJuridicaService.AtualizarSaldoPessoaJuridica(r.Context(), pessoaID, data.Saldo)
	if err != nil {
		if errors.Is(err, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao atualizar o saldo da pessoa jurídica.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"data": response,
	})
}

func (a *Api) handlerUpdateContaPessoaJuridicaPorId(w http.ResponseWriter, r *http.Request) {
	data, problemas, err := utils.DecodificarJson[pessoajuridica.AtualizarPessoaJuridicaRequest](r)
	if err != nil {
		_ = utils.CodificarJson(w, r, http.StatusUnprocessableEntity, problemas)
		return
	}

	pessoaID, err := uuid.Parse(chi.URLParam(r, "conta_id"))
	if err != nil {
		writeErrorJSON(w, r, http.StatusBadRequest, "Pessoa Id inválido!")
		return
	}

	response, err := a.PessoaJuridicaService.AtualizarPessoaJuridica(
		r.Context(),
		pessoaID,
		data.Faturamento,
		data.Idade,
		data.NomeFantasia,
		data.Celular,
		data.EmailCorporativo,
		data.Categoria,
		data.Saldo,
	)
	if err != nil {
		if errors.Is(err, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado.Error())
			return
		}

		if errors.Is(err, pessoaJuridicaService.ErroEmailJaExiste) {
			writeErrorJSON(w, r, http.StatusBadRequest, pessoaJuridicaService.ErroEmailJaExiste.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao atualizar a pessoa jurídica.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"data": response,
	})
}

func (a *Api) handlerDeleteContaPessoaJuridicaPorId(w http.ResponseWriter, r *http.Request) {
	pessoaID, err := uuid.Parse(chi.URLParam(r, "conta_id"))
	if err != nil {
		writeErrorJSON(w, r, http.StatusBadRequest, "Pessoa Id inválido!")
		return
	}

	if err := a.PessoaJuridicaService.ExcluirPessoaJuridica(r.Context(), pessoaID); err != nil {
		if errors.Is(err, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado) {
			writeErrorJSON(w, r, http.StatusNotFound, pessoaJuridicaService.ErrorPessoaJuridicaNaoEncontrado.Error())
			return
		}

		writeErrorJSON(w, r, http.StatusInternalServerError, "Erro interno ao excluir a pessoa jurídica.")
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"mensagem": "Pessoa jurídica removida com sucesso!",
	})
}
