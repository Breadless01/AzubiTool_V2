package ports

import "AzubiTool/internal/domain"

type AuthService interface {
	Authenticate(username, password string) (*domain.User, error)
}
