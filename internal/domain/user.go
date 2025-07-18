package domain

type User struct {
	LoggedIn    bool
	Username    string
	DisplayName string
	Email       string
	Role        string
	Oid         string
}
