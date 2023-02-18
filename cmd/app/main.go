package main

import (
	"dockerized/api/pkg/controller/rest"
	repository "dockerized/api/pkg/repository/postgres"

	"dockerized/api/tools/http"
	"dockerized/api/tools/postgres"
)

func main() {

	httpConfig, err := http.NewHttpConfig("")
	if err != nil {
		panic(err)
	}

	postgresConfig, err := postgres.NewPostgresConfig("")
	if err != nil {
		panic(err)
	}

	client, err := postgres.NewPostgresClient(postgresConfig)
	if err != nil {
		panic(err)
	}
	defer client.DB.Close()

	database := repository.NewRepository(client)

	restHandler := rest.NewRestHandler(httpConfig, database)

	err = restHandler.Serve()
	if err != nil {
		panic(err)
	}
}
