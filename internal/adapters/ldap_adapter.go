package adapters

import "yourproject/internal/ports"

// LdapAdapter implementiert das Authenticator-Interface mit LDAP-Logik.
type LdapAdapter struct{}

func (l *LdapAdapter) Authenticate(username, password string) (ports.User, error) {
    // TODO: Implementieren
    return ports.User{}, nil
}
