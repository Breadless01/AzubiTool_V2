package adapters

import (
	"AzubiTool/internal/domain"
	"fmt"
	"os"

	"github.com/go-ldap/ldap/v3"
)

type LdapAdapter struct {
	Host       string // z.B. "ldap://ldap.firma.de:389"
	BaseDN     string // z.B. "OU=Users,DC=example,DC=de"
	UserFilter string // z.B. "(sAMAccountName=%s)" oder "(uid=%s)"
}

func NewLdapAdapter(addr, baseDN, userFilter string) *LdapAdapter {
	return &LdapAdapter{
		Host:       addr,
		BaseDN:     baseDN,
		UserFilter: userFilter,
	}
}

func (a *LdapAdapter) Authenticate(username, password string) (*domain.User, error) {
	l, err := ldap.DialURL(a.Host)
	if err != nil {
		return nil, fmt.Errorf("LDAP-Verbindung fehlgeschlagen: %w", err)
	}
	defer l.Close()

	user_bind := fmt.Sprintf("%s%s", username, os.Getenv("LDAP_BIND_DIRECT_SUFFIX"))

	err = l.Bind(user_bind, password)
	if err != nil {
		return nil, fmt.Errorf("LDAP Bind failed: %w", err)
	}

	searchRequest := ldap.NewSearchRequest(
		a.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(%s=%s)", a.UserFilter, username),
		[]string{"sAMAccountName", "title", "name", "proxyAddresses"},
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil || len(sr.Entries) != 1 {
		return nil, fmt.Errorf("user nicht gefunden oder kein eindeutiger Treffer")
	}

	entry := sr.Entries[0]
	return &domain.User{
		LoggedIn:    true,
		Username:    entry.GetAttributeValue("sAMAccountName"),
		DisplayName: entry.GetAttributeValue("name"),
		Email:       entry.GetAttributeValue("proxyAddresses"),
		Role:        entry.GetAttributeValue("title"),
	}, nil
}
