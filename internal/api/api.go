package api

import (
	pessoafisica "gobank/internal/services/pessoa_fisica"
	"gobank/internal/services/user"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router              *chi.Mux
	Sessions            *scs.SessionManager
	UserService         user.UserSerivce
	PessoaFisicaService pessoafisica.PessoaFisicaService
}
