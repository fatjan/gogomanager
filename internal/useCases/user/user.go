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
	updateRequest := &dto.UserPatchRequest{
		Email:           user.Email,
		Name:            user.Name,
		UserImageUri:    user.UserImageUri,
		CompanyName:     user.CompanyName,
		CompanyImageUri: user.CompanyImageUri,
	}

	// Update fields if provided in request
	if request != nil {
		if request.Email != "" {
			updateRequest.Email = request.Email
		}
		if request.Name != "" {
			updateRequest.Name = request.Name
		}
		if request.UserImageUri != "" {
			updateRequest.UserImageUri = request.UserImageUri
		}
		if request.CompanyName != "" {
			updateRequest.CompanyName = request.CompanyName
		}
		if request.CompanyImageUri != "" {
			updateRequest.CompanyImageUri = request.CompanyImageUri
		}
	}

	// Update user in repository
	if err = u.userRepository.Update(ctx, userID, updateRequest); err != nil {
		log.Println(fmt.Errorf("failed to update user: %w", err))
		return nil, err
	}

	// Return new user data
	return &dto.User{
		Email:           updateRequest.Email,
		Name:            updateRequest.Name,
		UserImageUri:    updateRequest.UserImageUri,
		CompanyName:     updateRequest.CompanyName,
		CompanyImageUri: updateRequest.CompanyImageUri,
	}, nil
}
