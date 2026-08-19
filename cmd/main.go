package main

import (
	"fmt"
	"github.com/joaobotoni/knock/yaml"
	"log"
)

const (
	path = "etc/config.yaml"
)

func main() {
	
	data, err := yaml.Load(path)
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}
	fmt.Printf("Values: %s %d %s\n", data.Database.Host, data.Database.Port, data.Database.Name)

}
