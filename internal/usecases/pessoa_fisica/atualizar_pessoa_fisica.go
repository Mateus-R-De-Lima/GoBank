package pessoafisica

import (
	"context"
	"gobank/internal/validador"
)

type AtualizarPessoaFisicaRequest struct {
	RendaMensal  float64 `json:"renda_mensal"`
	Idade        int32   `json:"idade"`
	NomeCompleto string  `json:"nome_completo"`
	Celular      string  `json:"celular"`
	Email        string  `json:"email"`
	Categoria    string  `json:"categoria"`
	Saldo        float64 `json:"saldo"`
}

type AtualizarSaldoPessoaFisicaRequest struct {
	Saldo float64 `json:"saldo"`
}

func (req AtualizarPessoaFisicaRequest) Valid(ctx context.Context) validador.Avaliador {
	return validarPessoaFisicaBase(req.RendaMensal, req.Idade, req.NomeCompleto, req.Celular, req.Email, req.Categoria, req.Saldo)
}

func (req AtualizarSaldoPessoaFisicaRequest) Valid(ctx context.Context) validador.Avaliador {
	var a validador.Avaliador

	a.VerificarCampo(
		validador.MaiorLimiteDecimal(req.Saldo, 0),
		"saldo",
		"o saldo mínimo para atualização tem que ser maior que 0",
	)

	return a
}

func validarPessoaFisicaBase(rendaMensal float64, idade int32, nomeCompleto, celular, email, categoria string, saldo float64) validador.Avaliador {
	var a validador.Avaliador

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(nomeCompleto),
		"nome_completo",
		"o nome completo não pode estar em branco",
	)

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(celular),
		"celular",
		"o celular não pode estar em branco",
	)

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(email),
		"email",
		"o e-mail não pode estar em branco",
	)

	a.VerificarCampo(
		validador.Corresponde(email, validador.EmailRX),
		"email",
		"o e-mail informado é inválido",
	)

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(categoria),
		"categoria",
		"a categoria não pode estar em branco",
	)

	a.VerificarCampo(
		validador.MaximoValor(idade, 80),
		"idade",
		"a idade máxima para criação da conta é de 80 anos",
	)

	a.VerificarCampo(
		validador.MaiorOuIgualLimite(idade, 10),
		"idade",
		"a idade mínima para criação da conta é de 10 anos",
	)

	a.VerificarCampo(
		validador.MaiorLimiteDecimal(rendaMensal, 0),
		"renda_mensal",
		"a renda mensal mínima para criação tem que ser maior que 1 reais",
	)

	a.VerificarCampo(
		validador.MaiorLimiteDecimal(saldo, 0),
		"saldo",
		"o saldo mínimo para criação tem que ser maior que 1 reais",
	)

	return a
}
