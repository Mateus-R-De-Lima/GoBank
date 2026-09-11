package api

import (
	"errors"
	pessoaFisicaService "gobank/internal/services/pessoa_fisica"
	pessoafisica "gobank/internal/usecases/pessoa_fisica"
	"gobank/internal/utils"
	"net/http"

	"github.com/google/uuid"
)

func (a *Api) handleCriarContaPessoaFisica(w http.ResponseWriter, r *http.Request) {

	data, problemas, err := utils.DecodificarJson[pessoafisica.CriarPessoaFisicaRequest](r)

	if err != nil {
		_ = utils.CodificarJson(w, r, http.StatusUnprocessableEntity, problemas)
		return
	}

	userID := a.Sessions.Get(r.Context(), "AuthenticatedUserId")

	if userID == nil {
		_ = utils.CodificarJson(w, r, http.StatusUnauthorized, map[string]any{
			"error": "usuário não autenticado",
		})
		return
	}

	id, ok := userID.(uuid.UUID)
	if !ok {

		_ = utils.CodificarJson(w, r, http.StatusUnauthorized, map[string]any{
			"error": "id do usuário inválido",
		})
		return
	}

	novaPessoaFisicaId, err := a.PessoaFisicaService.CriarNovaPessoaFisica(r.Context(),
		data.NomeCompleto,
		data.Categoria,
		data.Celular,
		data.Email,
		data.Idade,
		data.Saldo,
		data.RendaMensal,
		id,
	)

	if err != nil {
		if errors.Is(err, pessoaFisicaService.ErroEmailJaExiste) {
			utils.CodificarJson(w, r, http.StatusBadRequest, map[string]any{
				"error": pessoaFisicaService.ErroEmailJaExiste.Error(),
			})
			return
		}

		_ = utils.CodificarJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "Erro internal ao criar a pessoa fisica.",
		})
		return

	}

	utils.CodificarJson(w, r, http.StatusCreated, map[string]any{
		"id":       novaPessoaFisicaId,
		"mensagem": "Pessoa fisica criada com sucesso!",
	})

}
func (a *Api) handlerGetContaPessoaFisicaPorId(w http.ResponseWriter, r *http.Request)        {}
func (a *Api) handlerPatchSaldoContaPessoaFisicaPorId(w http.ResponseWriter, r *http.Request) {}
func (a *Api) handlerUpdateContaPessoaFisicaPorId(w http.ResponseWriter, r *http.Request)     {}
func (a *Api) handlerDeleteContaPessoaFisicaPorId(w http.ResponseWriter, r *http.Request)     {}
