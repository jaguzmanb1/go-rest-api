package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rest-api/root"

	"github.com/gorilla/mux"
)

// Handler API para requests
type Handler struct {
	UserService root.UserService
}

var (
	handler *Handler
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

//CreateRoutes crea las rutas del API
func (h *Handler) CreateRoutes() {
	handler = h
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewareCors)
	router.HandleFunc("/users", getUsersRequest).Methods("GET")
	router.HandleFunc("/users", createUserRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))

}

//GetUsersRequest consulta todos los usuarios
func getUsersRequest(w http.ResponseWriter, r *http.Request) {
	users, err := handler.UserService.GetUsers()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func createUserRequest(w http.ResponseWriter, r *http.Request) {
	var newUser root.User
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Insert a valid user")
	}
	json.Unmarshal(reqBody, &newUser)
	handler.UserService.CreateUser(newUser)

}
