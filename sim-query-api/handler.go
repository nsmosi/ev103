package simqueryapi

import (
	"net/http"

	"github.com/SardarAndimeh/ev101/db"
	"github.com/gin-gonic/gin"
)

func getSimCard(context *gin.Context) {

	var (
		simData, bundle map[string]string
		err             error
		response        SimCard
	)

	msisdn := context.Param("msisdn")
	key := "msisdn:" + msisdn

	// loop over redis Clients
	for _, client := range db.Clients {
		simData, err = client.HGetAll(db.Ctx, key).Result()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch sim card from db"})
		}
		if len(simData) > 0 {
			bundleKey := "bundle:" + simData["bundleID"]
			bundle, err = db.CrdbClient.HGetAll(db.Ctx, bundleKey).Result()
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
					"Error":  "could not fetch Bundle from db",
					"Detail": err.Error(),
				})
			}
			break
		}
	}

	if len(simData) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "sim card does not exist"})
	} else {
		response = ResponseFormatter(simData, bundle)
	}

	context.IndentedJSON(http.StatusFound, response)

}

func deleteSimCard(context *gin.Context) {

	msisdn := context.Param("msisdn")
	key := "msisdn:" + msisdn

	var result int64
	var err error

	for _, client := range db.Clients {
		result, err = client.Del(db.Ctx, key).Result()
		if result > 0 {
			break
		}
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"Error":  "could not delete sim card",
				"Detail": err.Error(),
			})
		}
	}

	context.JSON(http.StatusFound, gin.H{
		"status": "success",
	})

}
