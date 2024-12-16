package entity

import "time"

type Post struct {
	ID        uint64    `gorm:"column:post_id;primaryKey;autoIncrement;not null;<-create"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint64    `gorm:"column:user_id;not null;uniqueIndex"`
}
