package handlers

import (
	"net/http"
	"rest-api/data"
)

//GetUsers consulta todos los usuarios
func (h *Users) GetUsers(w http.ResponseWriter, r *http.Request) {
	h.l.Println("[DEBUG] get all records")

	users, err := h.UserService.GetUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	err = data.ToJSON(users, w)
	if err != nil {
		h.l.Println("[ERROR] serializing user", err)
	}
}

//GetUser obtiene un solo usuario
func (h *Users) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := getProductID(r)
	h.l.Println("[DEBUG] get record id", userID)

	user, err := h.UserService.GetUser(userID)

	switch err {
	case nil:
	case data.ErrProductNotFound:
		h.l.Println("[ERROR] fetching product", err)
		return
	default:
		h.l.Println("[ERROR] fetching product", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(user, w)
	if err != nil {
		h.l.Println("[ERROR] serializing product", err)
	}
}
