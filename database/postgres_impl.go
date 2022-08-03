package database

import (
	"fmt"
	"github.com/galileoluna/tareas/repository"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type PostgresConfig struct {
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
}

func Init() {
	var configuration PostgresConfig
	err := envconfig.Process("", &configuration)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", configuration.PostgresUser, configuration.PostgresPassword, configuration.PostgresDB)
	repo, err := NewPostgresRepository(addr)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)

}
