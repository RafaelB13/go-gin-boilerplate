package config

import (
	"os"
)

func ServerListenner() string {
	DotEnv().LoadEnv()

	appURL := os.Getenv("APP_URL")
	appURL = func() string {
		if appURL != "" {
			return appURL
		}
		return ":"
	}()

	port := os.Getenv("APP_PORT")
	port = func() string {
		if port != "" {
			return ":" + port // Adicione ":" antes da porta
		}
		return ":8888" // Porta padrão, se não estiver definida no arquivo .env
	}()

	return appURL + port

}
