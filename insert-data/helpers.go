package insertdata

import (
	"log"
	"time"

	"math/rand"

	"github.com/SardarAndimeh/ev101/db"
	"github.com/go-redis/redis/v8"
)

func getShard(msisdn string, totalShards int) int {

	// Use the last digit of the MSISDN as the shard key
	lastDigit := msisdn[len(msisdn)-1]
	shard := int(lastDigit-'0') % totalShards
	return shard
}

func checkKeyExists(client *redis.Client, key string) bool {

	exists, err := client.Exists(db.Ctx, key).Result()
	if err != nil {
		log.Fatalf("failed to check if key exists ")
	}
	if exists > 0 {
		return true
	}
	return false
}

func getRandomBundleId() string {

	var bundleIDs []string
	keys, err := db.CrdbClient.Keys(db.Ctx, "*").Result()
	if err != nil {
		log.Fatalf("error geting bundle keys")
	}

	for _, key := range keys {
		result, err := db.CrdbClient.HGetAll(db.Ctx, key).Result()
		if err != nil {
			log.Fatalf("failed to get bundle from database")
		}
		bundleIDs = append(bundleIDs, result["ID"])
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomIndex := r.Intn(len(bundleIDs))

	return bundleIDs[randomIndex]

}
