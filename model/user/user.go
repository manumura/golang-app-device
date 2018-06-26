package usermodel

// User type
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Roles     []Role `json:"roles"`
}

// Role type
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
