package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DBConnection() (*sql.DB, error) {

	envMap, mapErr := godotenv.Read(".env")
	if mapErr != nil {
		panic(mapErr)
	}

	dbDriver := envMap["DB_driver"]
	dbUser := envMap["DB_user"]
	dbPass := envMap["DB_pass"]
	dbName := envMap["DB_name"]

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err

}
