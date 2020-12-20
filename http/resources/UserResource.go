package resources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-api/root"
	"strconv"

	"github.com/gorilla/mux"
)

//UserResource describe un recurso de usuario
type UserResource struct {
	UserService root.UserService
}

//GetUser obtiene un solo usuario
func (h *UserResource) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	user, err := h.UserService.GetUser(userID)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if user == nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "No se ha encontrado el usuario")
		return
	}

	json.NewEncoder(w).Encode(user)
}

//GetUsersRequest consulta todos los usuarios
func (h *UserResource) GetUsersRequest(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json.NewEncoder(w).Encode(users)
}

//CreateUserRequest crear un usuario en la persistencia
func (h *UserResource) CreateUserRequest(w http.ResponseWriter, r *http.Request) {
	var newUser root.User
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Insert a valid user")
	}
	json.Unmarshal(reqBody, &newUser)

	if !(len(newUser.Name) > 0) {
		fmt.Fprint(w, "El usuario necesita un nombre")
		return
	}

	h.UserService.CreateUser(newUser)
}

//DeleteUser obtiene un solo usuario con el id especificado
func (h *UserResource) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprint(w, "ID no valido")
		return
	}

	h.UserService.DeleteUser(userID)
}
