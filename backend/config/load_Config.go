package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JWTSecretKey string
}

var configuration Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("Service is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("Http Port is required")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("PORT must be number in env")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	if jwtSecretKey == "" {
		fmt.Println("JWT secret key is required")
		os.Exit(1)
	}

	configuration = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}

}

func init() {
	loadConfig()
}

// reciver function for config
func GetCongfig() Config {
	return configuration
}
