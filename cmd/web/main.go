package main

import (
	"AzubiTool/internal/web"
	"log"
	"net/http"
)

func main() {
	web.InitSessionStore()
	handler := &web.Handler{}
	router := web.NewRouter(handler)

	port := ":8080"
	log.Printf("Starte Webserver auf http://localhost%s ...", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Fehler beim Starten des Servers: %v", err)
	}
}
