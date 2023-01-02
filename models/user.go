package models

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}
