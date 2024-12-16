package repository

import (
	"blog_post/entity"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

// NewPostRepository initializes a new PostRepository with a database connection
func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (c *PostRepository) InsertPost(cover *entity.Post) error {
	return c.db.Create(&cover).Error
}

func (r *PostRepository) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	err := r.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
