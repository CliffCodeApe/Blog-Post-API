package contract

import "blog_post/dto"

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (dto.Response, error)
}
