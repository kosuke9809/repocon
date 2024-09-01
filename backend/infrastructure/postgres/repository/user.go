package repository

import (
	userDomain "repocon/domain/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userDomain.IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindByID(id string) (*userDomain.User, error) {
	user := userDomain.User{}
	err := ur.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) FindByGitHubID(githubID int64) (*userDomain.User, error) {
	user := userDomain.User{}
	err := ur.db.First(&user, "github_id = ?", githubID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) Create(user *userDomain.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) Update(user *userDomain.User) error {
	return ur.db.Save(user).Error
}

func (ur *userRepository) Delete(id string) error {
	return ur.db.Delete(&userDomain.User{}, "id = ?", id).Error
}
