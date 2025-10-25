package main

import (
	"log"
	"net/http"

	"credsystem/integration"
)

func main() {
	http.HandleFunc("/hello", integration.APIHandler)

	port := ":8080"
	log.Printf("Servidor rodando em http://localhost%s/hello", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
