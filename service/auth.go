package service

import (
	"blog_post/contract"
	"blog_post/dto"
	"blog_post/entity"
	"blog_post/repository"
	"blog_post/utils"
)

type AuthServiceImpl struct {
	authRepo *repository.AuthRepository
}

func NewAuthService(authRepo *repository.AuthRepository) contract.AuthService {
	return &AuthServiceImpl{authRepo: authRepo}
}

func (s *AuthServiceImpl) Register(req dto.RegisterRequest) error {
	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	return s.authRepo.SaveUser(user)
}
func (s *AuthServiceImpl) Login(req dto.LoginRequest) (dto.Response, error) {
	user, err := s.authRepo.GetUserByEmail(req.Email, req.Password)
	if err != nil {
		return dto.Response{}, err
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Email)
	if err != nil {
		return dto.Response{}, err
	}

	return dto.Response{Token: token}, nil
}
