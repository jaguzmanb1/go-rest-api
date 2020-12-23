package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/data"
	"rest-api/handlers"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// Se crea servicio de usuario
	us := &data.UserService{DB: db}

	// Se crea el handler de usuario
	uh := handlers.New(us, l)

	// Se crea un serve mux que pueda registrar los handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", uh.GetUsers)

	s := http.Server{
		Addr:         os.Getenv("bindAddress"), // configure the bind address
		Handler:      sm,                       // set the default handler
		ErrorLog:     l,                        // set the logger for the server
		ReadTimeout:  5 * time.Second,          // max time to read request from the client
		WriteTimeout: 10 * time.Second,         // max time to write response to the client
		IdleTimeout:  120 * time.Second,        // max time for connections using TCP Keep-Alive
	}

	go func() {
		l.Println("Starting server on port ", os.Getenv("bindAddress"))

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
