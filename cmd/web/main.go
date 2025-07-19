package main

import (
	"AzubiTool/internal/adapters"
	"AzubiTool/internal/config"
	"AzubiTool/internal/domain"
	"AzubiTool/internal/web"
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Register domain types for session storage
	gob.Register(domain.User{})

	// Connect to PostgreSQL database
	connData := map[string]any{
		"dbName": "azubitool",
		"dbUser": "azubi",
		"dbPass": "azubi",
		"dbHost": "localhost",
		"dbPort": 5432,
	}

	ConnStr := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", connData["dbUser"], connData["dbPass"], connData["dbHost"], connData["dbPort"])
	conn, err := pgx.Connect(context.Background(), ConnStr)
	if err != nil {
		log.Fatalf("Fehler beim Verbinden mit der Datenbank: %v", err)
	}
	connData["adminDB"] = conn
	defer conn.Close(context.Background())

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize adapters
	ldapAuth := adapters.NewLdapAdapter(
		cfg.LDAP_HOST,
		cfg.LDAP_BASE_DN,
		cfg.LDAP_USER_LOGIN_ATTR,
	)

	bcsApi := adapters.NewBcsAdapter(cfg.BCS_ENDPOINT_URL, true)

	db := adapters.NewPostgresAdapter(conn, connData)
	db.Initialize()

	web.InitSessionStore()
	handler := &web.Handler{
		AuthService: ldapAuth,
		BcsService:  bcsApi,
		DbService:   db,
	}
	router := web.NewRouter(handler)

	port := ":8080"
	log.Printf("Starte Webserver auf http://localhost%s ...", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Fehler beim Starten des Servers: %v", err)
	}
}
