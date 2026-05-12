package controllers

import (
	"encoding/json"
	"go-calender/models"
	u "go-calender/utils"
	"net/http"
	"fmt"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetOwnUserInfo = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	fmt.Print(id)
	fmt.Print("\n")
	data := models.GetUser(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
