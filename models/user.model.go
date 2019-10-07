package models

// User model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

// BindingUser used to bind the user from a query
type BindingUser struct {
	ID string `uri:"id" binding:"required"`
}
