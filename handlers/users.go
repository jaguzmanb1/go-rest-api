package handlers

import (
	"log"
	"net/http"
	"rest-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// Users handler for getting and updating users
type Users struct {
	UserService *data.UserService
	l           *log.Logger
	v           *data.Validation
}

// New crea un handler de usuario con el logger y servicio dado
func New(us *data.UserService, l *log.Logger, v *data.Validation) *Users {
	return &Users{us, l, v}
}

// KeyUser usada para el middleware
type KeyUser struct{}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
