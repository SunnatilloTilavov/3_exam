package main

import (
	"context"
	"fmt"
	"clone/3_exam/api"
	"clone/3_exam/config"
	"clone/3_exam/pkg/logger"
	"clone/3_exam/service"
	"clone/3_exam/storage/postgres"
	"clone/3_exam/storage/redis"

)

func main() {
	cfg := config.Load()
	
	newRedis := redis.New(cfg)

	log := logger.New(cfg.ServiceName)
	store, err := postgres.New(context.Background(), cfg,newRedis)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	services := service.New(store,log,newRedis)
	c := api.New(services,log)

	fmt.Println("programm is running on localhost:8080...")
	c.Run(":8080")
}

