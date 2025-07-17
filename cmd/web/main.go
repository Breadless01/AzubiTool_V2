package main

import (
	"AzubiTool/internal/adapters"
	"AzubiTool/internal/config"
	"AzubiTool/internal/web"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	ldapAuth := adapters.NewLdapAdapter(
		cfg.LDAP_HOST,
		cfg.LDAP_BASE_DN,
		cfg.LDAP_USER_LOGIN_ATTR,
	)

	web.InitSessionStore()
	handler := &web.Handler{
		AuthService: ldapAuth,
	}
	router := web.NewRouter(handler)

	port := ":8080"
	log.Printf("Starte Webserver auf http://localhost%s ...", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Fehler beim Starten des Servers: %v", err)
	}
}
