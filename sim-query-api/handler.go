package simqueryapi

import (
	"log"
	"net/http"

	"github.com/SardarAndimeh/ev101/db"
	"github.com/gin-gonic/gin"
)

func getSimData(context *gin.Context) {

	var (
		record, bundle map[string]string
		err            error
		response       gin.H
	)

	msisdn := context.Param("msisdn")
	key := "msisdn:" + msisdn

	for _, client := range db.Clients {
		record, err = client.HGetAll(db.Ctx, key).Result()
		if err != nil {
			log.Fatalf("could not fetch sim data from db")
		}
		if len(record) > 0 {
			bundleKey := "bundle:" + record["bundle"]
			bundle, err = db.CrdbClient.HGetAll(db.Ctx, bundleKey).Result()
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"Error": "could not fetch Bundle Id from db"})
			}
			break
		}
	}

	if len(record) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "sim card data not found"})
	} else {

		// Create the response in the desired format
		response = gin.H{
			"msisdn": msisdn,
			"imsi":   record["imsi"],
			"iccid":  record["iccid"],
			"secret": record["secret"],
			"tac":    record["tac"],
			"eid":    record["eid"],
			"cid":    record["cid"],
			"imei":   record["imei"],
			"bundle": gin.H{
				"id":       bundle["id"],
				"ul":       bundle["ul"],
				"dl":       bundle["dl"],
				"quota":    bundle["quota"],
				"duration": bundle["duration"],
				"label":    bundle["label"],
				"type":     bundle["type"], // capitalize first letter
			},
		}
	}

	context.IndentedJSON(http.StatusOK, gin.H{
		"message":  "data founded successfully ",
		"sim card": response,
	})

}
