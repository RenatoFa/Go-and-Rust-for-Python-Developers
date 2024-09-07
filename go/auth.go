package main

type Role string

const (
	Viewer    Role = "viewer"
	Developer Role = "developer"
	Admin     Role = "admin"
)

type User struct {
	Login string
	Role  Role
}

func Promoter(user *User, role Role) {
	user.Role = role
}
