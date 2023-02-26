package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"

	"dockerized/api/tools/postgres"
)

func main() {

	pathPtr := flag.String("file", "./cmd/importer/data/names.csv", "path to the csv data file")

	f, err := os.Open(*pathPtr)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	config, err := postgres.NewPostgresConfig("")
	if err != nil {
		panic(err)
	}

	client, err := postgres.NewPostgresClient(config)
	if err != nil {
		panic(err)
	}
	defer client.DB.Close()
	fmt.Printf("connecting to: %s\n", client.Config.Hostname)

	_, err = client.DB.Exec(`CREATE TABLE IF NOT EXISTS names (
		id SERIAL PRIMARY KEY,
		first_name TEXT,
		last_name TEXT
	);`)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(f)
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		_, err = client.DB.Exec(`INSERT INTO names (first_name, last_name) VALUES ($1, $2);`, rec[0], rec[1])
		if err != nil {
			panic(err)
		}
	}

}
