package repo

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "Aino"
// }

// INTERFACE
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	return false
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
