package main

import (
	"context"
	"fmt"
	"clone/rent_car_us/api"
	"clone/rent_car_us/config"
	"clone/rent_car_us/pkg/logger"
	"clone/rent_car_us/service"
	"clone/rent_car_us/storage/postgres"
	"clone/rent_car_us/storage/redis"

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

