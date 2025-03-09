package main

import (
	"fmt"
	"log"

	nginxcachepurger "github.com/dkvz/nginx-cache-purger"
)

func main() {
	conf, err := nginxcachepurger.ConfigFromDotEnv()
	if err != nil {
		log.Fatal("Could not load configuration: " + err.Error())
	}

	fmt.Printf("%v\n", conf)

}
