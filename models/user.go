package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" form:"name" binding:"required,min=3,max=100" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" form:"email" binding:"required,email,min=8,max=100" gorm:"unique;not null"`
	Password  string    `json:"-" binding:"required,min=6" gorm:"not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
