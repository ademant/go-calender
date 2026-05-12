package controllers

import (
	"encoding/json"
	"go-calender/models"
	u "go-calender/utils"
	"net/http"
	"fmt"
)

var DBAccountInit = func() {
	ac := &models.Account{}
	ac.DBAccountInit()
}

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	fmt.Print(account.User)
	resp := models.Login(account.User, account.Password)
	u.Respond(w, resp)
}

var DeAuthenticate = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	fmt.Print(id)
	resp := models.Logout(id)
	u.Respond(w,resp)

}

