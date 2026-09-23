package auth

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(server *http.ServeMux) {
	// Guests should be able to view register and login pages
	http.HandleFunc("POST /register", registerHandler)
	http.HandleFunc("POST /login", loginHandler)

	// Users need to be able to logout, view profile, update profile
	http.HandleFunc("GET /logout", logoutHandler)
	http.HandleFunc("GET /profile/{id}", viewProfileHandler)
	http.HandleFunc("PATCH /profile/{id}/edit", editProfileHandler)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Register User")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login User")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logout User")
}

func viewProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "View Profile")
}

func editProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit Profile")
}