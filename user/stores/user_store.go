package stores

import (
	"context"
	"satellite/user/models"

	"gorm.io/gorm"
)

type UserStore interface {
	InsertUser(ctx context.Context, user *models.User) (*models.User, error)
}

type PostgresUserStore struct {
	Client *gorm.DB
}

func NewPostgresUserStore(client *gorm.DB) *PostgresUserStore {
	return &PostgresUserStore{
		Client: client,
	}
}

func (s *PostgresUserStore) InsertUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := s.Client.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
