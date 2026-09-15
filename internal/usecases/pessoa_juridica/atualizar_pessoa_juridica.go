package pessoajuridica

import (
	"context"
	"gobank/internal/validador"
)

type AtualizarPessoaJuridicaRequest struct {
	Faturamento      float64 `json:"faturamento"`
	Idade            int32   `json:"idade"`
	NomeFantasia     string  `json:"nome_fantasia"`
	Celular          string  `json:"celular"`
	EmailCorporativo string  `json:"email_corporativo"`
	Categoria        string  `json:"categoria"`
	Saldo            float64 `json:"saldo"`
}

type AtualizarSaldoPessoaJuridicaRequest struct {
	Saldo float64 `json:"saldo"`
}

func (req AtualizarPessoaJuridicaRequest) Valid(ctx context.Context) validador.Avaliador {
	return validarPessoaJuridicaBase(req.Faturamento, req.Idade, req.NomeFantasia, req.Celular, req.EmailCorporativo, req.Categoria, req.Saldo)
}

func (req AtualizarSaldoPessoaJuridicaRequest) Valid(ctx context.Context) validador.Avaliador {
	var a validador.Avaliador

	a.VerificarCampo(
		validador.MaiorLimiteDecimal(req.Saldo, 0),
		"saldo",
		"o saldo mínimo para atualização tem que ser maior que 0",
	)

	return a
}

func validarPessoaJuridicaBase(faturamento float64, idade int32, nomeFantasia, celular, emailCorporativo, categoria string, saldo float64) validador.Avaliador {
	var a validador.Avaliador

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(nomeFantasia),
		"nome_fantasia",
		"o nome fantasia não pode estar em branco",
	)

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(celular),
		"celular",
		"o celular não pode estar em branco",
	)

	a.VerificarCampo(
		validador.NaoPoderEstaEmBranco(emailCorporativo),
		"email_corporativo",
		"o e-mail corporativo não pode estar em branco",
	)

	a.VerificarCampo(
		validador.Corresponde(emailCorporativo, validador.EmailRX),
		"email_corporativo",
		"o e-mail corporativo informado é inválido",
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
		validador.MaiorLimiteDecimal(faturamento, 0),
		"faturamento",
		"o faturamento mínimo para criação tem que ser maior que 1 reais",
	)

	a.VerificarCampo(
		validador.MaiorLimiteDecimal(saldo, 0),
		"saldo",
		"o saldo mínimo para criação tem que ser maior que 1 reais",
	)

	return a
}
