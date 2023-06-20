package model

import "time"

type User struct {
    ID        uint      `json:"id" grom:"primaryKey"`
    Email     string    `json:"email" grom:"unique"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
    ID    uint   `json:"id" grom:"primaryKey"`
    Email string `json:"email" grom:"unique"`	
}
