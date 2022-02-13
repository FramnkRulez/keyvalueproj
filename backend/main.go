// Main go module responsible for starting up the HTTP listener and router and calling the
// kvprest go package for handling incoming requests.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FramnkRulez/keyvalueproj/pkg/kvprest"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("keyvalue backend server")

	router := mux.NewRouter().StrictSlash(true)

	kvprest.SetupHandlers(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
