package redisrepository

import (
	"context"
	_domain "microservice/shared/domain"

	"github.com/go-redis/redis/v8"
)

type userRedisRepository struct {
	Redis *redis.Client
}

// NewRedisUserRepository will create an object that represent the userRepository interface
func NewUserRedisRepository(Redis *redis.Client) _domain.UserRedisRepository {
	return &userRedisRepository{Redis: Redis}
}

func (u *userRedisRepository) GetByID(c context.Context, id string) (_domain.User, error) {
	us := _domain.User{}
	return us, nil
}
