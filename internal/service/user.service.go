package service

import (
	"fmt"

	"github.com/dinhdc/go-ecommerce/internal/repo"
	"github.com/dinhdc/go-ecommerce/internal/utils/random"
	"github.com/dinhdc/go-ecommerce/pkg/response"
)

// INTERFACE
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 0. hash email

	// 5. check OTP is avaibale

	// 6. user spam otp

	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasAlreadyExist
	}
	// 2, new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Sprintf("Otp is :::%d\n", otp)
	// 3. save OTP in Redis with exp
	
	// 4. send email OTP
	return response.ErrCodeSuccess
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
