package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"rest-api/database"
	"rest-api/http"

	_ "github.com/go-sql-driver/mysql"
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

func main() {
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// Se crea servicio de usuario
	userService := &database.UserService{DB: db}

	// Se crea handlerHttp para las peticiones
	var httpHandler http.Handler
	httpHandler.UserService = userService
	httpHandler.CreateRoutes()

}
