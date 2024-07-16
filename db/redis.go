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

}
