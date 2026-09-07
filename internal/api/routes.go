package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *Api) BindRoutes() {
	a.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	a.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", a.handleSignupUser)
				r.Post("/login", a.handleLoginUser)
				r.Group(func(r chi.Router) {
					r.Use(a.AuthMiddleware)
					r.Post("/logout", a.handlerLogoutUser)
				})
			})
			r.Route("/pessoa-fisica", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(a.AuthMiddleware)
					r.Post("/", a.handleCriarContaPessoaFisica)
					r.Route("/{conta_id}", func(r chi.Router) {
						r.Get("/", a.handlerGetContaPessoaFisicaPorId)
						r.Patch("/saldo", a.handlerPatchSaldoContaPessoaFisicaPorId)
						r.Put("/", a.handlerUpdateContaPessoaFisicaPorId)
						r.Delete("/", a.handlerDeleteContaPessoaFisicaPorId)
					})

				})
			})
		})
	})
}
