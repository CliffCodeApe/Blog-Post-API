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

func (r *PostRepository) FindById(id uint64) (entity.Post, error) {
	var post entity.Post
	err := r.db.First(&post, id).Error
	return post, err
}

func (r *PostRepository) UpdatePost(id uint64, post entity.Post) error {
	return r.db.Model(&entity.Post{}).Where("post_id = ?", id).Updates(post).Error
}

func (r *PostRepository) DeletePost(id uint64) error {
	return r.db.Delete(&entity.Post{}, id).Error
}
