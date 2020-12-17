package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var (
	//ConnectionString cadena de conexión a la base de datos
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

func getDatabase() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}
