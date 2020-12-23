package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

//Users describe un recurso de usuario
type Users struct {
	UserService *data.UserService
	l           *log.Logger
}

// New crea un handler de usuario con el logger y servicio dado
func New(us *data.UserService, l *log.Logger) *Users {
	return &Users{us, l}
}

//GetUser obtiene un solo usuario
func (h *Users) GetUser(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle GET User")

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	user, err := h.UserService.GetUser(userID)
	if err != nil {
		http.Error(w, "No se ha podid obtener al usuario", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

//GetUsers consulta todos los usuarios
func (h *Users) GetUsers(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle GET Users")

	users, err := h.UserService.GetUsers()
	if err != nil {
		http.Error(w, "No se ha podido realizar una conexión al a base de datos: ", http.StatusInternalServerError)
	}
	err = users.ToJSON(w)
	if err != nil {
		http.Error(w, "No se ha podido leer el JSON de salida", http.StatusInternalServerError)
	}
}

//CreateUser crear un usuario en la persistencia
func (h *Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle POST Users")

	user := r.Context().Value(KeyUser{}).(data.User)
	h.UserService.CreateUser(user)
}

//DeleteUser obtiene un solo usuario con el id especificado
func (h *Users) DeleteUser(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle DELETE Users")

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprint(w, "ID no valido")
		return
	}

	h.UserService.DeleteUser(userID)
}

// KeyUser usada para el middleware
type KeyUser struct{}

//MiddlewareValidateUser  verificacion para los request
func (h *Users) MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := data.User{}

		err := user.FromJSON(r.Body)
		if err != nil {
			h.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
