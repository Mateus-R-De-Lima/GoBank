package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Erro ao carregar o arquivo .env")
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Falha ao executar o comando: ", err)
		fmt.Println("Saída do comando: ", string(output))
		panic("Erro ao executar a migração: " + err.Error())
	}

	fmt.Println("Comando executado com sucesso: ", string(output))

}
