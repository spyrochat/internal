package models

type UserModel struct {
   Id uint `json:"id" db:"id"`
   Name string `json:"name" db:"name"`
   Email string `json:"email" db:"email"`
}
