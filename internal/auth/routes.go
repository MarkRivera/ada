package auth

import (
	"fmt"
	"net/http"

	"github.com/MarkRivera/ada/internal/config"
)

// Guests should be able to view register and login pages
// Users need to be able to logout, view profile, update profile

func RegisterHandler(app *config.Application) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Register User")
	}
}


func LoginHandler(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Login User")
	}
}

func LogoutHandler(app *config.Application) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Logout User")
	}
}

func ViewProfileHandler(app *config.Application) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "View Profile")
	}
}

func EditProfileHandler(apap *config.Application) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Edit Profile")
	}	
}
