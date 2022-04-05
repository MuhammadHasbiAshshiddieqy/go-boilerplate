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

func (u *userUsecase) GetByID(c *fiber.Ctx, id string) (_dto.UserResponse, error) {
	res, err := u.userRepo.GetByID(c, id)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	return _mapper.MapUserToUserResponse(res), nil
}

func (u *userUsecase) Fetch(c *fiber.Ctx, pagination _dto.Pagination) (_dto.Pagination, error) {
	res, err := u.userRepo.Fetch(c, &pagination)
	if err != nil {
		return _dto.Pagination{}, err
	}

	if len(res) != 0 {
		pagination.Rows = _mapper.MapUsersToUserResponses(res)
	}

	return pagination, nil
}

func (u *userUsecase) Update(c *fiber.Ctx, ureq _dto.UserRequestUpdate) (_dto.UserResponse, error) {
	us, err := u.userRepo.GetByID(c, ureq.ID)
	if err != nil {
		return _dto.UserResponse{}, err
	}
	usr, err := _mapper.MapUserRequestUpdateToUser(ureq, us)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	res, err := u.userRepo.Update(c, usr)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	return _mapper.MapUserToUserResponse(res), nil
}

func (u *userUsecase) Delete(c *fiber.Ctx, id string) error {
	err := u.userRepo.Delete(c, id)
	if err != nil {
		return err
	}

	return nil
}
