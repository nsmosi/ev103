package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SardarAndimeh/ev101/config"
)

func main() {

	// parsing commad line arguments
	configPath := flag.String("config", "", "path to the config file")
	flag.Parse()

	if *configPath == "" {
		log.Fatalf("no arguments is provided , usage : go run main.go -config  path to config file")
	}

	// Loading Config data
	configData, err := config.LoadConfigData(*configPath)
	if err != nil {
		log.Fatalf("failed to  load config data")
	}

	fmt.Println(configData)

}
