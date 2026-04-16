package main

import (
	"fmt"
//	"github.com/gorilla/mux"
	"go-calender/app"
	"go-calender/controllers"
	"net/http"
	"os"
)

func main() {

//	router := mux.NewRouter()
	router := http.NewServeMux()

	router.HandleFunc("POST /api/role", controllers.CreateRole)
	router.HandleFunc("GET /api/role/{id}", controllers.GetRole)
	router.HandleFunc("POST /api/user/new", controllers.CreateAccount)
	router.HandleFunc("POST /api/user/login", controllers.Authenticate)
	router.HandleFunc("POST /api/contacts/new", controllers.CreateContact)
	router.HandleFunc("POST /api/v1/location", controllers.CreateLocation)
	router.HandleFunc("GET /api/v1/location", controllers.GetLocations)
	router.HandleFunc("GET /api/v1/location/{id}", controllers.GetLocation)
	router.HandleFunc("PATCH /api/v1/location/{id}", controllers.PatchLocation)
	router.HandleFunc("DELETE /api/v1/location/{id}", controllers.DeleteLocation)
	router.HandleFunc("GET /api/me/contacts", controllers.GetContactsFor) //  user/2/contacts

//	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, app.JwtAuthentication(router)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
