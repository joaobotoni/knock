package main

import (
	"fmt"
	"log"

	"github.com/joaobotoni/knock/env"
	"github.com/joaobotoni/knock/yaml"
)

const (
	ManifestPath = "etc/manifest.yaml"
	EnvPath      = ".env"
	Addr         = ":8080"
)

func main() {
	if err := env.Load(EnvPath); err != nil {
		log.Fatalf("Erro ao carregar as variaveis de ambiente: %v", err)
	}

	manifest, err := yaml.Load(ManifestPath)
	if err != nil {
		log.Fatalf("Erro ao carregar os arquivos de configuração: %v", err)
	}

	fmt.Printf("Database: %s\n%s\n", manifest.Database.Host,  manifest.Database.Name)
}
