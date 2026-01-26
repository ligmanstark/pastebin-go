package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RoleID   int    `json:"role_id"`
	Password string `json:"password,omitempty"`
}
