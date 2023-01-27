package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"database/sql"
	"github.com/prueba_ingreso/database"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var DB *sql.DB = database.ConnectDB()

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/API/login/v1", addUserExample).Methods("POST")

	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token", "Authorization", "Accept", "Accept-Language"},
		AllowedMethods:   []string{"GET", "PATCH", "POST", "PUT", "OPTIONS", "DELETE", "COPY"},
		Debug:            true,
		AllowCredentials: true,
	})

	handler := cor.Handler(router)

	log.Fatal(http.ListenAndServe(":"+"8080", handler))
}