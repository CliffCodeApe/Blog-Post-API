package entity

import "time"

type Post struct {
	ID        uint64    `gorm:"column:post_id;primaryKey;autoIncrement;not null;<-create"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	Author    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
