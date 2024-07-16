package simqueryapi

import (
	"log"
	"net/http"

	"github.com/SardarAndimeh/ev101/db"
	"github.com/gin-gonic/gin"
)

func getSimData(context *gin.Context) {

	var (
		simData, bundle map[string]string
		err             error
		response        SimCard
	)

	msisdn := context.Param("msisdn")
	key := "msisdn:" + msisdn

	for _, client := range db.Clients {
		simData, err = client.HGetAll(db.Ctx, key).Result()
		if err != nil {
			log.Fatalf("could not fetch sim data from db")
		}
		if len(simData) > 0 {
			bundleKey := "bundle:" + simData["bundle"]
			bundle, err = db.CrdbClient.HGetAll(db.Ctx, bundleKey).Result()
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"Error": "could not fetch Bundle Id from db"})
			}
			break
		}
	}

	if len(simData) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "sim card data not found"})
	} else {
		response = ResponseFormatter(simData, bundle)
	}

	context.IndentedJSON(http.StatusOK, response)

}
