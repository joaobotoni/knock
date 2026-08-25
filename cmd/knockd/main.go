package main

import (
	"log"
	"context"
	"github.com/joaobotoni/knock/env"
	"github.com/joaobotoni/knock/yaml"
	"github.com/joaobotoni/knock/postgres"

)

const (
	ManifestPath = "etc/manifest.yaml"
	EnvPath      = ".env"
	Addr         = ":8080"
)

func main() {

	current := context.Background()

	if err := env.Load(EnvPath); err != nil {
		log.Fatalf("Erro ao carregar as variaveis de ambiente: %v", err)
	}

	manifest, err := yaml.Load(ManifestPath)
	if err != nil {
		log.Fatalf("Erro ao carregar os arquivos de configuração: %v", err)
	}

	pool, err := postgres.Open(&current, &manifest.Database)
	if err != nil {
		log.Fatalf("Erro ao connectar ao banco de dados: %v", err)
	}
	defer pool.Close()
}
