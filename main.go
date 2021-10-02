package main

import (
	"log"
	"net/http"

	"go-migration/router"
)

func main() {

	r := router.Router()

	log.Fatal(http.ListenAndServe(":8080", r))

}
