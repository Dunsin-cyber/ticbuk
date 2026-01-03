package repositories

import (
	"context"

	"github.com/Dunsin-cyber/ticbuk/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) models.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) RegisterUser(ctx context.Context, registerData *models.AuthCredentials) (*models.User, error) {
	user := &models.User{
		Email:    registerData.Email,
		Password: registerData.Password,
	}
	res := r.db.Model(&models.User{}).WithContext(ctx).Create(user)

	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error) {
	user := &models.User{}
	if res := r.db.Model(user).WithContext(ctx).Where(query, args).First(user); res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
