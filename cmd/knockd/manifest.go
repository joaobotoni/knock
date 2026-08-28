package main

import "github.com/joaobotoni/knock/postgres"

type Manifest struct {
	Database postgres.Config `yaml:"database"`
}
