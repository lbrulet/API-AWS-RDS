package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type BindingUser struct {
	ID string `uri:"id" binding:"required"`
}
