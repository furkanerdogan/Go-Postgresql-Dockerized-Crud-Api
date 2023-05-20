package models


type User struct {
ID uint64 `json:"id"`
Name string `json: "name"`
Surname string `json:"surname"`
Email string `json:"email"`
RoleID int `json: "role_id"`
}