package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primary_key;"`
	GitHubID  int64          `json:"github_id" gorm:"uniqueIndex;not null;"`
	Email     string         `json:"email" `
	Username  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
