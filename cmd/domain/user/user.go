package user

// Request represents user request type
type Request struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// Response represents user response type
type Response struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// User represent user type
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// IsEmptyName check empty username
func (u *User) IsEmptyName() bool {
	return len(u.Name) == 0
}

// Model type
type Model struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}
