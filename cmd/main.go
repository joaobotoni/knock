package main

import (
	"fmt"
	"log"
	"github.com/joaobotoni/knock/env"
	"github.com/joaobotoni/knock/yaml"
)

const (
	PATH = "etc/config.yaml"
	ENV  = ".env"
)

func main() {

	if err := env.Load(ENV); err != nil {
		log.Fatalf("Erro ao carregar as variaveis de ambiente: %v", err)
	}

	data, err := yaml.Load(PATH)
	if err != nil {
		log.Fatalf("Erro ao carregar os arquivos de configuração: %v", err)
	}
	
	fmt.Printf("Data: %s\n", data.Database.Host)
}
