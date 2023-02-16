package main

import (
	"context"
	"go-redis/curd"
	"go-redis/dao"
)

func main() {
	client := dao.Init()
	curd.C(context.Background(), client)
}
