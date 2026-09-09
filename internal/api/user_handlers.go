package api

import (
	"errors"
	usuario "gobank/internal/services/user"
	"gobank/internal/usecases/user"
	"gobank/internal/utils"
	"net/http"
)

func (a *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problemas, err := utils.DecodificarJson[user.CriarUserRequest](r)

	if err != nil {
		_ = utils.CodificarJson(w, r, http.StatusUnprocessableEntity, problemas)
	}

	id, err := a.UserService.CriarUsuario(
		r.Context(),
		data.UserName,
		data.Email,
		data.Password,
	)

	if err != nil {
		if errors.Is(err, usuario.ErroEmailOuUserNameJaExiste) {
			_ = utils.CodificarJson(w, r, http.StatusBadRequest, map[string]any{
				"error": usuario.ErroEmailOuUserNameJaExiste,
			})
			return
		}

		_ = utils.CodificarJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "Erro ao criar o usuario.",
		})
		return
	}

	_ = utils.CodificarJson(w, r, http.StatusCreated, map[string]any{
		"id": id,
	})
}
func (a *Api) handleLoginUser(w http.ResponseWriter, r *http.Request)   {}
func (a *Api) handlerLogoutUser(w http.ResponseWriter, r *http.Request) {}
