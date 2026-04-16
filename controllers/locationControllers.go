package controllers

import (
	"encoding/json"
	"go-calender/models"
	u "go-calender/utils"
	"net/http"
	"strconv"
	"fmt"
)

var CreateLocation = func(w http.ResponseWriter, r *http.Request) {
	location := &models.Location{}
	err := json.NewDecoder(r.Body).Decode(location)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := location.Create()
	fmt.Print(resp["status"])
	if resp["status"] == true {
		w.WriteHeader(http.StatusCreated)
	}else{
		w.WriteHeader(http.StatusBadRequest)
	}
	u.Respond(w, resp)
}

var PatchLocation = func(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")
	if sid == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		u.Respond(w, u.Message(false, "No Location entry to patch"))
		return
	}
	id,err := strconv.ParseUint(sid,10,32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Could not parse given ID: " + sid))
		return
	}
	location := &models.Location{}
	err = json.NewDecoder(r.Body).Decode(location)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	location.ID = uint(id)
        fmt.Print(location)
	data := models.GetLocation(uint(id))
	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Location not found with ID: " + sid))
		return
	}
	resp := location.Update()
	if resp["status"] == true {
		w.WriteHeader(http.StatusAccepted)
	}else{
		w.WriteHeader(http.StatusBadRequest)
	}
	u.Respond(w, resp)
}

var DeleteLocation = func(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Context().Value("user"))
	sid := r.PathValue("id")
	if sid == "" {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "No Location entry to patch"))
		return
	}
	id,err := strconv.ParseUint(sid,10,32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Could not parse given ID: " + sid))
		return
	}
	data := models.GetLocation(uint(id))
	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "No location with ID: " + sid))
		return
	}
	resp := data.Delete()
	u.Respond(w, resp)
}

var GetLocations = func(w http.ResponseWriter, r *http.Request) {
	//id := r.Context().Value("user").(uint)
	data := models.GetLocations()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetLocation = func(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")
	id,err := strconv.ParseUint(sid,10,32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Could not parse given ID: " + sid))
		return
	}
	data := models.GetLocation(uint(id))
	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Location not found with ID: " + sid))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
