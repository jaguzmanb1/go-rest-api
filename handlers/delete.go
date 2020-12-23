package handlers

import (
	"net/http"
)

//DeleteUser obtiene un solo usuario con el id especificado
func (h *Users) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := getProductID(r)

	h.l.Println("[DEBUG] Deleting user with id ", userID)

	h.UserService.DeleteUser(userID)
}
