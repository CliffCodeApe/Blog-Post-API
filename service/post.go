package service

import (
	"blog_post/contract"
	"blog_post/dto"
	"blog_post/entity"
	"blog_post/repository"
)

type PostServiceImpl struct {
	postRepo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) contract.PostService {
	return &PostServiceImpl{postRepo: repo}
}

func (s *PostServiceImpl) CreatePost(postDTO dto.PostRequest) error {
	post := entity.Post{
		Title:   postDTO.Title,
		Content: postDTO.Content,
		Author:  postDTO.Author,
	}
	return s.postRepo.InsertPost(&post)
}

func (s *PostServiceImpl) GetPosts() ([]dto.GetPost, error) {
	posts, err := s.postRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []dto.GetPost
	for _, post := range posts {
		result = append(result, dto.GetPost{
			ID:        int(post.ID),
			Title:     post.Title,
			Content:   post.Content,
			Author:    post.Author,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}
	return result, nil
}

func (s *PostServiceImpl) GetPostByID(id uint64) (entity.Post, error) {
	return s.postRepo.FindById(id)
}

func (s *PostServiceImpl) UpdatePost(id uint64, req dto.EditRequest) error {
	post := entity.Post{}
	if req.Title != nil {
		post.Title = *req.Title
	}
	if req.Content != nil {
		post.Content = *req.Content
	}
	return s.postRepo.UpdatePost(id, post)
}

func (s *PostServiceImpl) DeletePost(id uint64) error {
	return s.postRepo.DeletePost(id)
}
