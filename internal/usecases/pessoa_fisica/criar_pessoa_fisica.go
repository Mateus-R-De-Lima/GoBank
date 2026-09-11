package pessoafisica

import (
	"context"
	"gobank/internal/validador"
)

type CriarPessoaFisicaRequest struct {
	RendaMensal  float64 `json:"renda_mensal"`
	Idade        int32   `json:"idade"`
	NomeCompleto string  `json:"nome_completo"`
	Celular      string  `json:"celular"`
	Email        string  `json:"email"`
	Categoria    string  `json:"categoria"`
	Saldo        float64 `json:"saldo"`
}

func (req CriarPessoaFisicaRequest) Valid(ctx context.Context) validador.Avaliador {
	var a validador.Avaliador

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(req.NomeCompleto),
		"nome_completo",
		"o nome completo não pode estar em branco",
	)

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(req.Celular),
		"celular",
		"o celular não pode estar em branco",
	)

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(req.Email),
		"email",
		"o e-mail não pode estar em branco",
	)

	a.VerificarCampo(
		validador.Corresponde(req.Email, validador.EmailRX),
		"email",
		"o e-mail informado é inválido",
	)

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(req.Categoria),
		"categoria",
		"a categoria não pode estar em branco",
	)

	a.VerificarCampo(
		validador.MaximoValor(req.Idade, 80),
		"idade",
		"a idade máxima para criação da conta é de 80 anos",
	)

	a.VerificarCampo(
		validador.MaiorOuIgualLimite(req.Idade, 10),
		"idade",
		"a idade mínima para criação da conta é de 10 anos",
	)

	a.VerificarCampo(
		validador.MaiorOuIgualLimite(req.Idade, 0),
		"idade",
		"a idade mínima para criação da conta é de 10 anos",
	)

	a.VerificarCampo(
		validador.MaiorLimiteDecimal(req.RendaMensal, 0),
		"renda_mensal",
		"a renda mensal mínima para criação tem que ser maior que 1 reais",
	)

	a.VerificarCampo(
		validador.MaiorLimiteDecimal(req.Saldo, 0),
		"saldo",
		"o saldo mínimo para criação tem que ser maior que 1 reais",
	)
	return a
}
