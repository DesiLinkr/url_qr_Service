package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Port string

func init() {
	godotenv.Load()
	Port = getEnv("PORT", "8080")

}

func getEnv(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val

}
func getEnvAsInt(key string, fallback int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Println("⚠️ Invalid int for", key, "using fallback")
		return fallback
	}

	return val
}
