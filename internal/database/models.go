package database

type UserModel struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
