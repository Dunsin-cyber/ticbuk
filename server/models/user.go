package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	Manager  UserRole = "manager"
	Attendee UserRole = "attendee"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex" json:"username"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Role      UserRole  `json:"role" gorm:"text;default:attendee"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID == 1 {
		db.Model(u).Update("role", Manager)
	}
	return nil
} 