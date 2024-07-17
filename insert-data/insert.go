package insertdata

import (
	"fmt"
	"log"
	"time"

	"github.com/SardarAndimeh/ev101/db"
)

func AddSimCards(filePath string) error {

	start := time.Now()

	records, err := LoadCSV(filePath)
	if err != nil {
		log.Fatalf("could not Load CSV data: %v", err)
	}

	headers := []string{"msisdn", "imsi", "iccid", "secret", "tac", "eid", "cid", "imei", "bundleID"}

	totalRedis := len(db.Clients)
	totalInsertedRecords := 0

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
				if header == "bundleID" {
					recordMap[header] = getRandomBundleId()
				} else {
					recordMap[header] = record[index]
				}
			}
		}

		err = db.Clients[shardIndex].HSet(db.Ctx, key, recordMap).Err()
		if err != nil {
			return fmt.Errorf("inserting sim cards data into db failed %w", err)
		}
		totalInsertedRecords++
	}

	log.Printf("inserting Sim cards was successful !")

	// a brief report
	elapsed := time.Since(start)

	fmt.Printf("Duration : %v \n", elapsed)
	fmt.Printf("Number of inserted records : %v", totalInsertedRecords)

	return nil
}

func AddBundles(filePath string) error {

	bundleHeader := []string{"id", "ul", "dl", "quota", "duration", "label", "type"}

	bundles, err := LoadCSV(filePath)
	if err != nil {
		log.Fatalf("could not Load CSV data")
	}

	for _, bundle := range bundles {
		key := "bundle:" + bundle[0]

		// create a map of key and values for each bundle record ->  "ID": 1021 , "UL": 234534 ,....
		recordMap := make(map[string]interface{})
		for index, header := range bundleHeader {
			if index < len(bundle) {
				recordMap[header] = bundle[index]
			}
		}

		err := db.CrdbClient.HSet(db.Ctx, key, recordMap).Err()
		if err != nil {
			return fmt.Errorf("inserting bundles into db failed: %w", err)
		}
	}

	log.Printf("inserting bundles was successful !")

	return nil

}
