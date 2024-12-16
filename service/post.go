package service

import (
	"blog_post/contract"
	"blog_post/dto"
	"blog_post/entity"
	"blog_post/repository"
)

type PostServiceImpl struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) contract.PostService {
	return &PostServiceImpl{repo: repo}
}

func (s *PostServiceImpl) CreatePost(postDTO dto.PostRequest, userID uint64, author string) error {
	post := entity.Post{
		Title:   postDTO.Title,
		Content: postDTO.Content,
		Author:  author,
		UserID:  userID,
	}
	return s.repo.InsertPost(&post) // Ensure to return the error if any
}

func (s *PostServiceImpl) GetPosts() ([]dto.GetPost, error) {
	posts, err := s.repo.FindAll() // Call the new GetAllPosts method
	if err != nil {
		return nil, err // Return the error if there is one
	}

	var result []dto.GetPost
	for _, post := range posts {
		result = append(result, dto.GetPost{
			Title:     post.Title,
			Content:   post.Content,
			Author:    post.Author,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}
	return result, nil
}
