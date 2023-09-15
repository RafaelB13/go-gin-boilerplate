package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Env struct{}

func DotEnv() *Env {
	return &Env{}
}

func (c *Env) LoadEnv() {
	// Carregue as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}
}
