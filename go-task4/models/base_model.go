package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	BaseModel
	Username string    `gorm:"unique;not null" json:"username"`
	Password string    `gorm:"not null" json:"-"`
	Email    string    `gorm:"unique" json:"email"`
	Posts    []Post    `gorm:"foreignKey:UserID" json:"-"`
	Comments []Comment `gorm:"foreignKey:UserID" json:"-"`
}

type Post struct {
	BaseModel
	Title    string    `gorm:"not null" json:"title"`
	Content  string    `gorm:"type:text;not null" json:"content"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"user"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments"`
}
type Comment struct {
	BaseModel
	Content string `gorm:"type:text;not null" json:"content"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	PostID  uint   `gorm:"not null" json:"post_id"`
	Post    Post   `gorm:"foreignKey:PostID" json:"-"`
}
