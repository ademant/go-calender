package controllers

import (
	"encoding/json"
	"go-calender/models"
	u "go-calender/utils"
	"net/http"
	"strconv"
	"fmt"
)

var CreateRole = func(w http.ResponseWriter, r *http.Request) {

	role := &models.Roles{}
	err := json.NewDecoder(r.Body).Decode(role) //decode the request body into struct and failed if any error occur
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := role.Create() //Create account
	fmt.Print(resp["status"])
	if resp["status"] == true {
		w.WriteHeader(http.StatusCreated)
	}else{
		w.WriteHeader(http.StatusBadRequest)
	}
	u.Respond(w, resp)
}

var GetRole = func(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")
	id,err := strconv.ParseUint(sid,10,32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Could not parse given ID: " + sid))
		return
	}
	data := models.GetRole(uint(id))
	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false,"Role not found with ID: " + sid))
		return
	}
	resp := u.Message(true,"success")
	resp["data"] = data
	u.Respond(w,resp)
}
