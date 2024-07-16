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

	// Loop through CSV data (sim cards data)
	for _, record := range records {

		msisdn := record[0]
		key := "msisdn:" + msisdn

		// create a map of key and values for each Sim card record ->  "MSISDN": 811502214250 , "IMSI": 217500013105250 ,....
		recordMap := make(map[string]interface{})
		for index, header := range headers {
			if index < len(record) {
				recordMap[header] = record[index]
			}
		}

		//insert vreated Map into database
		err = db.Clients[0].HSet(db.Ctx, key, recordMap).Err()
		if err != nil {
			return fmt.Errorf("inserting data into db (HSet) failed %w", err)
		}
	}
	return nil
}

func getShard(msisdn string, totalShards int) int {

	// Use the last digit of the MSISDN as the shard key
	lastDigit := msisdn[len(msisdn)-1]
	shard := int(lastDigit-'0') % totalShards
	return shard
}
