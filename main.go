package main

import (
	"flag"
	"log"

	"github.com/SardarAndimeh/ev101/config"
	"github.com/SardarAndimeh/ev101/db"
	simqueryapi "github.com/SardarAndimeh/ev101/sim-query-api"
	"github.com/gin-gonic/gin"
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

	db.InitRedis(configData.Rdbs.Address, configData.CrdbAddress)

	//runInsertion(configData)

	// sim query api
	server := gin.Default()
	simqueryapi.Routes(server)

	server.Run(":8080")

}

/* func runInsertion(configData *config.Config) {
	// insert data module
	err := insertdata.AddBundles(configData.BundleFile)
	if err != nil {
		log.Println(err)
	}

	err = insertdata.AddSimCards(configData.DataFile)
	if err != nil {
		log.Println(err)
	}
}
*/
