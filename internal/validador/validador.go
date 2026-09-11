package validador

import (
	"context"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validador interface {
	Valid(context.Context) Avaliador
}

type Avaliador map[string]string

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$")

func (a *Avaliador) AddError(campo, mensagem string) {

	if *a == nil {
		*a = make(map[string]string)
	}

	if _, existe := (*a)[campo]; !existe {
		(*a)[campo] = mensagem
	}

}

// Função que verifica se um campo é válido, onde ok é um booleano que indica se o campo é válido ou não, chave é o nome do campo e mensagem é a mensagem de erro. Se ok for false, a mensagem de erro é adicionada ao Avaliador.
func (e *Avaliador) VerificarCampo(ok bool, chave, mensagem string) {
	if !ok {
		e.AddError(chave, mensagem)
	}
}

// Função que verifica se o valor de uma string é diferente de vazio, removendo os espaços em branco. Retorna true se a string não for vazia, e false caso contrário.
func NaoPoderEstaEmBranco(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaximoCaracteres(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func MinimoCaracteres(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func Corresponde(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func MaiorOuIgualLimite(value int32, limite int) bool {

	return int(value) >= limite
}

func MaximoValor(value int32, max int) bool {

	return int(value) <= max
}

func MaiorLimiteDecimal(value float64, limite float64) bool {

	return float64(value) > limite
}
