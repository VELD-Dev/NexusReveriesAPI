package main

import (
	"fmt"
	"net/http"
	nexrev_routes "nexusreveries/cdn/routes"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Booting up Nexus Reveries API...")

	initRouter()
}

func initRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/dialog-hashes", nexrev_routes.DialogHashesGet)
	router.HandleFunc("/dialog-files", nexrev_routes.DialogFilesGet)
	router.HandleFunc("/localizations-hashes", nexrev_routes.GetLocalizationsHashes)
	router.HandleFunc("/localizations-files", nexrev_routes.GetLocalizationsFiles)

	http.ListenAndServe(":2772", router)
}
