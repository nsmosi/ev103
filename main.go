package main

import (
	"flag"
	"log"
)

func main() {
	// parsing commad line arguments
	configPath := flag.String("config", "", "path to the config file")
	flag.Parse()

	if *configPath == "" {
		log.Fatalf("no arguments is provided , usage : go run main.go -config  path to config file")
	}

}
