package curd

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func C(ctx context.Context, r *redis.Client) {
	rdb = r
	create(ctx)
}

func create(ctx context.Context) {
	err := rdb.Set(ctx, "t1", "t1v", 0).Err()
	if err != nil {
		log.Fatalln("redis set error:", err)
	}
}
