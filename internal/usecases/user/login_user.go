package user

import (
	"context"
	"gobank/internal/validador"
)

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserRequest) Valid(ctx context.Context) validador.Avaliador {
	var a validador.Avaliador

	a.VerificarCampo(validador.NaoPoderEstaEmBranco(req.Password), "password", "Password não pode esta em branco")
	a.VerificarCampo(validador.NaoPoderEstaEmBranco(req.Email), "email", "Email não pode esta em branco")

	a.VerificarCampo(validador.MinimoCaracteres(req.Password, 8) && validador.MaximoCaracteres(req.Password, 255), "passwoerd", "Password deve ter um comprimento entre 8 e 255 ")

	a.VerificarCampo(validador.Corresponde(req.Email, validador.EmailRX), "email", "Email não é valido!")
	return a
}
