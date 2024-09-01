package user

type IUserRepository interface {
	FindByID(id string) (*User, error)
	FindByGitHubID(gitHubID int64) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
}
