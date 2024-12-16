package contract

import (
	"blog_post/dto"
	"blog_post/entity"
)

type PostService interface {
	CreatePost(postDTO dto.PostRequest) error
	GetAllPosts() ([]dto.GetPost, error)
	GetPostByID(id uint64) (entity.Post, error)
	UpdatePost(id uint64, postDTO dto.EditRequest) error
	DeletePost(id uint64) error
}
