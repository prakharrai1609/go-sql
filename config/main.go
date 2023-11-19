package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Port     string
	DbSource string
	DbDriver string
	Db       *sql.DB
}

var config Config

func init() {
	var err error

	if err = godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	config.Port = os.Getenv("PORT")
	config.DbDriver = os.Getenv("DB_DRIVER")
	config.DbSource = os.Getenv("DB_SOURCE")
	config.Db, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to check if the connection is successful
	err = config.Db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
}

func Get() Config {
	return config
}
