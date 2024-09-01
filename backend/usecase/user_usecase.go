package usecase

import (
	"fmt"
	userDomain "repocon/domain/user"

	"gorm.io/gorm"
)

type IUserUsecase interface {
	GetUser(id string) (*userDomain.User, error)
	UpsertUser(user *userDomain.User) (*userDomain.User, error)
}

type userUsecase struct {
	ur userDomain.IUserRepository
}

func NewUserUsecase(ur userDomain.IUserRepository) IUserUsecase {
	return &userUsecase{
		ur: ur,
	}
}

func (uu *userUsecase) GetUser(id string) (*userDomain.User, error) {
	return uu.ur.FindByID(id)
}

func (uu *userUsecase) UpsertUser(user *userDomain.User) (*userDomain.User, error) {
	existingUser, err := uu.ur.FindByGitHubID(user.GitHubID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("error not found user: %w", err)
	}
	if existingUser == nil {
		err = uu.ur.Create(user)
		if err != nil {
			return nil, fmt.Errorf("error creating user: %w", err)
		}
		return user, nil
	}
	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.AvatarURL = user.AvatarURL

	err = uu.ur.Update(existingUser)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}
	return existingUser, nil
}
