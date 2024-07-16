package simqueryapi

import (
	"log"
	"net/http"

	"github.com/SardarAndimeh/ev101/db"
	"github.com/gin-gonic/gin"
)

type SimCard struct {
	Msisdn int64  `json:"msisdn"`
	Imsi   int64  `json:"imsi"`
	Iccid  int64  `json:"iccid"`
	Secret string `json:"secret"`
	Tac    int    `json:"tac"`
	Eid    int    `json:"eid"`
	Cid    int    `json:"cid"`
	Imei   int64  `json:"imei"`
	Bundle struct {
		ID       int    `json:"id"`
		Ul       int    `json:"ul"`
		Dl       int    `json:"dl"`
		Quota    int64  `json:"quota"`
		Duration int    `json:"duration"`
		Label    string `json:"label"`
		Type     string `json:"type"`
	} `json:"bundle"`
}

func getSimData(context *gin.Context) {

	msisdn := context.Param("msisdn")
	key := "msisdn:" + msisdn

	var record map[string]string
	var err error

	for _, client := range db.Clients {
		record, err = client.HGetAll(db.Ctx, key).Result()
		if err != nil {
			log.Fatalf("failed to fetch data from db")
		}
		if len(record) > 0 {
			bundleKey := "bundle:" + record["bundle"]
			bundle, err := db.CrdbClient.HGetAll(db.Ctx, bundleKey).Result()
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"Error": "could not load Bundle Id from db"})
			}

			context.IndentedJSON(http.StatusFound, gin.H{
				"message":  "found data",
				"sim card": record,
				"bundle":   bundle,
			})
			break
		}
	}

	if len(record) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "sim card data not found"})
	}

}
