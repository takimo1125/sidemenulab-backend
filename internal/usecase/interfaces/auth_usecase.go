package interfaces

import "sidemenulab-backend/internal/domain/entity"

type AuthUseCase interface {
	SignUp(req *entity.SignUpRequest) (*entity.AuthResponse, error)
	SignIn(req *entity.SignInRequest) (*entity.AuthResponse, error)
}
