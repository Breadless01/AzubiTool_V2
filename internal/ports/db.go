package ports

import "AzubiTool/internal/domain"

type DbService interface {
	Initialize() error
	GetUserByUsername(username string) (*domain.User, error)
	SetUser(user *domain.User) error
}
