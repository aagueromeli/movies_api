package main

import (
	"log"

	"github.com/aagueromeli/movies_api/cmd/server"
)

func main() {
	app, err := server.LoadServerConf()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
