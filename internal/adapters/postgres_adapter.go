package adapters

import (
	"context"
	"fmt"
	"log"

	"AzubiTool/internal/domain"

	"github.com/jackc/pgx/v5"
)

type PostgresAdapter struct {
	conn   *pgx.Conn
	dbData map[string]any
}

func NewPostgresAdapter(conn *pgx.Conn, dbData map[string]any) *PostgresAdapter {
	println("PostgresAdapter initialized with connection data:", dbData["dbName"])
	return &PostgresAdapter{conn: conn, dbData: dbData}
}

func (r *PostgresAdapter) Initialize() error {
	if targetDB := r.dbData["dbName"]; targetDB != "" {
		// Check if the database exists, if not create it
		adminConn := r.dbData["adminDB"].(*pgx.Conn)
		var exists bool
		err := adminConn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", targetDB).Scan(&exists)
		if err != nil {
			fmt.Errorf("Fehler beim Check der Datenbank: %v", err)
		}
		if !exists {
			_, err = adminConn.Exec(context.Background(), "CREATE DATABASE "+targetDB.(string))
			if err != nil {
				log.Fatalf("Fehler beim Erstellen der Datenbank: %v", err)
			}
			log.Printf("Datenbank '%s' wurde erstellt.", targetDB)
		} else {
			log.Printf("Datenbank '%s' existiert bereits.", targetDB)
		}

		// Reconnect to the target database
		targetConnStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", r.dbData["dbUser"], r.dbData["dbPass"], r.dbData["dbHost"], r.dbData["dbPort"], targetDB)
		userConn, err := pgx.Connect(context.Background(), targetConnStr)
		if err != nil {
			log.Fatalf("Verbindung zur Ziel-DB fehlgeschlagen: %v", err)
		}
		defer userConn.Close(context.Background())

		// Create the users table if it doesn't exist
		createUserTableSQL := `
			CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				username TEXT UNIQUE NOT NULL,
				display_name TEXT,
				email TEXT,
				role TEXT
			);`

		_, err = userConn.Exec(context.Background(), createUserTableSQL)
		if err != nil {
			log.Fatalf("Fehler beim Anlegen der User Tabelle: %v", err)
		}
		log.Printf("Tabelle 'users' ist bereit!")
		r.conn = userConn // Update the connection to the target database

		return err
	}

	log.Println("Keine Ziel-Datenbank angegeben, Initialisierung übersprungen.")
	return nil

}

func (r *PostgresAdapter) GetUserByUsername(username string) (*domain.User, error) {
	var u domain.User
	err := r.conn.QueryRow(context.Background(),
		"SELECT username, display_name, email, role FROM users WHERE username = $1", username,
	).Scan(&u.Username, &u.DisplayName, &u.Email, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *PostgresAdapter) SetUser(user *domain.User) error {
	_, err := r.conn.Exec(context.Background(),
		`INSERT INTO users (username, display_name, email, role)
         VALUES ($1, $2, $3, $4)
         ON CONFLICT (username) DO UPDATE
         SET display_name=$2, email=$3, role=$4`,
		user.Username, user.DisplayName, user.Email, user.Role,
	)
	return err
}
