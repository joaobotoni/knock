package main

import (
	"context"
	"github.com/joaobotoni/knock/env"
	"github.com/joaobotoni/knock/khttp"
	"github.com/joaobotoni/knock/postgres"
	"github.com/joaobotoni/knock/yaml"
	"log"
)

const (
	ManifestPath = "etc/manifest.yaml"
	EnvPath      = ".env"
	Addr         = ":8443"

	TLSCertPath = "etc/tls/server.crt"
	TLSKeyPath  = "etc/tls/server.key"

	JWTPrivateKeyPath = "etc/jwt/private.pem"
	JWTPublicKeyPath  = "etc/jwt/public.pem"
)

func main() {

	current := context.Background()

	if err := env.Load(EnvPath); err != nil {
		log.Fatalf("Erro ao carregar as variaveis de ambiente: %v", err)
	}

	config, err := yaml.Load[Manifest](ManifestPath)
	if err != nil {
		log.Fatalf("Erro ao carregar os arquivos de configuração: %v", err)
	}

	pool, err := postgres.Open(current, &config.Database)
	if err != nil {
		log.Fatalf("Erro ao connectar ao banco de dados: %v", err)
	}
	defer pool.Close()

	mux := khttp.NewMux(khttp.Routes{})
	server := khttp.NewServer(Addr, mux)

	log.Printf("Servidor escutando em %s", Addr)
	if err := server.ListenAndServeTLS(TLSCertPath, TLSKeyPath); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
