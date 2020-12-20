package http

import (
	"log"
	"net/http"
	"rest-api/http/resources"
	"rest-api/root"

	"github.com/gorilla/mux"
)

// Handler API para requests
type Handler struct {
	UserService root.UserService
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Content-Type", "application/json")
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
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewareCors)

	usersResource := resources.UserResource{UserService: h.UserService}

	router.HandleFunc("/users", usersResource.GetUsersRequest).Methods("GET")
	router.HandleFunc("/users", usersResource.CreateUserRequest).Methods("POST")
	router.HandleFunc("/users/{id}", usersResource.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", usersResource.GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))

}
