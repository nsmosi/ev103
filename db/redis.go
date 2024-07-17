package db

import (
	"context"
	"strings"

	"github.com/go-redis/redis/v8"
)

var (
	Clients    []*redis.Client
	CrdbClient *redis.Client
	Ctx        context.Context
)

func InitRedis(rdbs string, crdb string) {

	Ctx = context.Background()

	rdbAddrs := strings.Split(rdbs, ",")

	Clients = make([]*redis.Client, len(rdbAddrs))

	// Rdb Clients
	for i, rdbAddr := range rdbAddrs {
		client := redis.NewClient(&redis.Options{
			Addr:     rdbAddr,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		Clients[i] = client

	}

	CrdbClient = redis.NewClient(&redis.Options{
		Addr:     crdb,
		Password: "",
		DB:       0,
	})

}
