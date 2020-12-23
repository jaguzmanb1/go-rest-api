package handlers

import (
	"net/http"
	"rest-api/data"
)

//CreateUser crear un usuario en la persistencia
func (h *Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.l.Println("[DEBUG] Creating user")

	user := r.Context().Value(KeyUser{}).(data.User)
	h.UserService.CreateUser(user)
}
