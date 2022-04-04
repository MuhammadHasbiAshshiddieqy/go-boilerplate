package usecase

import (
	"github.com/gofiber/fiber/v2"

	_domain "microservice/shared/domain"
	_dto "microservice/shared/dto"
	_mapper "microservice/shared/pkg/mapper"
)

type userUsecase struct {
	userRepo _domain.UserRepository
}

// NewUserUsecase will create new an userUsecase object representation of domain.UserUsecase interface
func NewUserUsecase(u _domain.UserRepository) _domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) Store(c *fiber.Ctx, ureq _dto.UserRequestCreate) (_dto.UserResponse, error) {
	us, err := _mapper.MapUserRequestCreateToUser(ureq)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	res, err := u.userRepo.Store(c, us)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	return _mapper.MapUserToUserResponse(res), nil
}
