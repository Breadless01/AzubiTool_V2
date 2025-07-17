package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SESSION_KEY             string
	BCS_ENDPOINT_URL        string
	TENANT_ID               string
	SITE_ID                 string
	DRIVE_ID                string
	CLIENT_ID               string
	CLIENT_SECRET           string
	HEADER_PREFIX           string
	GRANT_TYPE              string
	SCOPE                   string
	AUTH_URL                string
	DOCUMENT_URL            string
	LDAP_HOST               string
	LDAP_BIND_DIRECT_SUFFIX string
	LDAP_BASE_DN            string
	LDAP_USER_LOGIN_ATTR    string
	APP_EMAIL_PASSWORD      string
	APP_EMAIL               string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		SESSION_KEY:             os.Getenv("SESSION_KEY"),
		BCS_ENDPOINT_URL:        os.Getenv("BCS_ENDPOINT_URL"),
		TENANT_ID:               os.Getenv("TENANT_ID"),
		SITE_ID:                 os.Getenv("SITE_ID"),
		DRIVE_ID:                os.Getenv("DRIVE_ID"),
		CLIENT_ID:               os.Getenv("CLIENT_ID"),
		CLIENT_SECRET:           os.Getenv("CLIENT_SECRET"),
		HEADER_PREFIX:           os.Getenv("HEADER_PREFIX"),
		GRANT_TYPE:              os.Getenv("GRANT_TYPE"),
		SCOPE:                   os.Getenv("SCOPE"),
		AUTH_URL:                os.Getenv("AUTH_URL"),
		DOCUMENT_URL:            os.Getenv("DOCUMENT_URL"),
		LDAP_HOST:               os.Getenv("LDAP_HOST"),
		LDAP_BIND_DIRECT_SUFFIX: os.Getenv("LDAP_BIND_DIRECT_SUFFIX"),
		LDAP_BASE_DN:            os.Getenv("LDAP_BASE_DN"),
		LDAP_USER_LOGIN_ATTR:    os.Getenv("LDAP_USER_LOGIN_ATTR"),
		APP_EMAIL_PASSWORD:      os.Getenv("APP_EMAIL_PASSWORD"),
		APP_EMAIL:               os.Getenv("APP_EMAIL"),
	}

	return cfg, nil
}
