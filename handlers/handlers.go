package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hugoing97/twittor/middlew"
	"github.com/hugoing97/twittor/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
