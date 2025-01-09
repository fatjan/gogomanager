package user

import (
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/repositories/user"
)

type useCase struct {
	userRepository user.Repository
}

func NewUseCase(userRepository user.Repository) UseCase {
	return &useCase{userRepository: userRepository}
}

func (u *useCase) GetUser(userRequest *dto.UserRequest) (*dto.User, error) {
	user, err := u.userRepository.GetUser(userRequest.UserID)
	if err != nil {
		return nil, err
	}

	return &dto.User{
		Email:           user.Email,
		Name:            user.Name,
		UserImageUri:    user.UserImageUri,
		CompanyName:     user.CompanyName,
		CompanyImageUri: user.CompanyImageUri,
	}, nil
}
