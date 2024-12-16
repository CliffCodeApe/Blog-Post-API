package contract

import (
	"blog_post/dto"
)

type PostService interface {
	CreatePost(postDTO dto.PostRequest, userID uint64, author string) error
	GetPosts() ([]dto.GetPost, error)
	// UpdatePost(id int, postDTO dto.EditRequest) error
	// DeletePost(id int) error
}
