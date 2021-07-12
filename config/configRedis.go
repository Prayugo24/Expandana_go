package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"os"
	// "fmt"
)

func SetupRedis() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		panic("gagal Load env")
	}
	db_host := os.Getenv("DB_HOST")
	client := redis.NewClient(&redis.Options{
		Addr:db_host,
		Password:"",
		DB:0, 
	})
	
	return client
}