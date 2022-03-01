package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "config: ", log.LstdFlags)
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatalf("Fatal read config file, error: %s ", err)
	}

	// Create postgresql connection
	logger.SetPrefix("postgres: ")
	logger.Println("Connecting postgresql...")
	dbconnect := viper.GetString("mysql_connection.connection")
	db, err := sql.Open("mysql", dbconnect)
	if err != nil {
		logger.Fatalf("Fatal postgresql connection, error: %s ", err)
	}
	if err != nil {
		logger.Fatalf("Fatal postgresql ping, error: %s ", err)
	}
	logger.Println("Successfully connected to postgresql!")
	defer db.Close()

}
