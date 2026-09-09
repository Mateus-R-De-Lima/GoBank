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
		return
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
				"error": usuario.ErroEmailOuUserNameJaExiste.Error(),
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
func (a *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problemas, err := utils.DecodificarJson[user.LoginUserRequest](r)

	if err != nil {
		utils.CodificarJson(w, r, http.StatusUnprocessableEntity, problemas)
	}

	id, err := a.UserService.AuthenticarUsuario(r.Context(), data.Email, data.Password)

	if err != nil {
		if errors.Is(err, usuario.ErrorCredenciasInvalidas) {
			utils.CodificarJson(w, r, http.StatusUnauthorized, map[string]any{
				"error": usuario.ErrorCredenciasInvalidas.Error(),
			})
			return
		}

		_ = utils.CodificarJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "Erro  internoao realizar o login.",
		})

		return
	}

	err = a.Sessions.RenewToken(r.Context())

	if err != nil {
		_ = utils.CodificarJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "Erro interno ao realizar o login.",
		})

		return
	}

	a.Sessions.Put(r.Context(), "AuthenticatedUserId", id)
	utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"mensagem": "login realizado com sucesso",
	})
}
func (a *Api) handlerLogoutUser(w http.ResponseWriter, r *http.Request) {
	err := a.Sessions.RenewToken(r.Context())

	if err != nil {
		utils.CodificarJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "Erro interno ao realizar o logout",
		})
		return
	}

	a.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	utils.CodificarJson(w, r, http.StatusOK, map[string]any{
		"mensagem": "logout realizado com sucesso",
	})
}
