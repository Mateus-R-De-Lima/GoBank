package api

import (
	"github.com/go-chi/chi/middleware"
)

func (a *Api) BindRoutes() {
	a.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)
}
