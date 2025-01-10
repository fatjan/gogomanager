package user

import (
	"context"
	"fmt"
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/repositories/user"
	"log"
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
func (u *useCase) UpdateUser(ctx context.Context, userID int, request *dto.UserPatchRequest) (*dto.User, error) {
	// Get existing user
	user, err := u.userRepository.GetUser(userID)
	if err != nil {
		log.Println(fmt.Errorf("failed to get user: %w", err))
		return nil, err
	}

	// Create update request with current values as defaults
	updateRequest := &dto.UserPatchRequest{}

	// Update fields if provided in request
	if request != nil {

		// check key request email
		if request.Email != nil {
			updateRequest.Email = request.Email
			user.Email = *request.Email
		}

		// check key request name
		if request.Name != nil {
			updateRequest.Name = request.Name
			user.Name = *request.Name
		}

		// check key request UserImageUri
		if request.UserImageUri != nil {
			updateRequest.UserImageUri = request.UserImageUri
			user.UserImageUri = *request.UserImageUri
		}

		// check key request CompanyName
		if request.CompanyName != nil {
			updateRequest.CompanyName = request.CompanyName
			user.CompanyName = *request.CompanyName
		}

		// check key request CompanyImageUri
		if request.CompanyImageUri != nil {
			updateRequest.CompanyImageUri = request.CompanyImageUri
			user.CompanyImageUri = *request.CompanyImageUri
		}

	}

	// Update user in repository
	if err = u.userRepository.Update(ctx, userID, updateRequest); err != nil {
		log.Println(fmt.Errorf("failed to update user: %w", err))
		return nil, err
	}

	// Return new user data
	return &dto.User{
		Email:           user.Email,
		Name:            user.Name,
		UserImageUri:    user.UserImageUri,
		CompanyName:     user.CompanyName,
		CompanyImageUri: user.CompanyImageUri,
	}, nil
}
