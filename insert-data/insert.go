package insertdata

import (
	"fmt"
	"log"

	"github.com/SardarAndimeh/ev101/db"
)

func AddSimCards(filePath string) error {

	records, err := LoadCSV(filePath)
	if err != nil {
		log.Fatalf("could not Load CSV data: %v", err)
	}

	headers := []string{"MSISDN", "IMSI", "ICCID", "Secret", "TAC", "EID", "CID", "IMEI", "BundleID"}
	totalRedis := len(db.Clients)

	// Loop through CSV data (sim cards data)
	for _, record := range records {

		msisdn := record[0]
		key := "msisdn:" + msisdn
		shardIndex := getShard(msisdn, totalRedis) // chose Redis Client according to Shard index (last 2 digits of MSISDN)

		if checkKeyExists(db.Clients[shardIndex], key) {
			continue
		}

		// create a map of key and values for each Sim card record ->  "MSISDN": 811502214250 , "IMSI": 217500013105250 ,....
		recordMap := make(map[string]interface{})
		for index, header := range headers {
			if index < len(record) {
				recordMap[header] = record[index]
			}
		}

		//insert vreated Map into database
		err = db.Clients[shardIndex].HSet(db.Ctx, key, recordMap).Err()
		if err != nil {
			return fmt.Errorf("inserting data into db (HSet) failed %w", err)
		}
	}
	return nil
}
