package ports

// Authenticator definiert die Schnittstelle zur Authentifizierung (z.B. LDAP).
type Authenticator interface {
    Authenticate(username, password string) (User, error)
}
