package usecase

import (
	_domain "microservice/shared/domain"
)

type userUsecase struct {
	userMysqlRepo _domain.UserMysqlRepository
	userRedisRepo _domain.UserRedisRepository
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(umsqlr _domain.UserMysqlRepository, urdsr _domain.UserRedisRepository) _domain.UserUsecase {
	return &userUsecase{
		userMysqlRepo: umsqlr,
		userRedisRepo: urdsr,
	}
}
