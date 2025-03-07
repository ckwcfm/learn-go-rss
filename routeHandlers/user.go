package routeHandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/models"
	"github.com/ckwcfm/learn-go/rss/services"
	"github.com/ckwcfm/learn-go/rss/utils"
	"github.com/go-playground/validator/v10"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// TODO: need to find a way to support json and form data
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := services.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginUser")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")
	user := models.User{
		Email:    email,
		Password: password,
	}
	log.Println("User:", user)
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = services.ValidateUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := services.CreateToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Token created:", token)
	// add a cookie to the response
	cookie := utils.CreateTokenCookie(token)
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"token": token,
	}

	json.NewEncoder(w).Encode(response)

}
