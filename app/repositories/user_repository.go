package repositories

import (
	"context"
	"github.com/nfangxu/core-skeleton/app/entities"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *gorm.DB
}

func (r UserRepository) Find(ctx context.Context, id string) (*entities.User, error) {
	var u entities.User
	err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user")
	}
	return &u, nil
}
