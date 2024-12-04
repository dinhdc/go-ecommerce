package service

import (
	"github.com/dinhdc/go-ecommerce/internal/repo"
	"github.com/dinhdc/go-ecommerce/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

// INTERFACE
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasAlreadyExist
	}
	return response.ErrCodeSuccess
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
