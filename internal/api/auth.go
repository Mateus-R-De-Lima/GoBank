package api

import (
	"gobank/internal/utils"
	"net/http"
)

func (a *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !a.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			utils.CodificarJson(w, r, http.StatusUnauthorized, map[string]any{
				"mensagem": "É necessário estar logado",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
