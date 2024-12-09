package internal

type User struct {
	ID       int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password_hash"`
}

func (User) TableName() string {
	return `"livecode2".Users`
}
