package main

import (
	"log"

	nginxcachepurger "github.com/dkvz/nginx-cache-purger"
	"github.com/dkvz/nginx-cache-purger/http"
)

func main() {
	conf, err := nginxcachepurger.ConfigFromDotEnv()
	if err != nil {
		log.Fatal("Could not load configuration: " + err.Error())
	}

	server := http.NewServer(conf)
	log.Println("starting server...")
	server.ListenAndServe(3000)

}
