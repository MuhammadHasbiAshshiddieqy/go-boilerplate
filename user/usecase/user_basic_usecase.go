package usecase

import (
	"context"

	_dto "microservice/shared/dto"
	_mapper "microservice/shared/pkg/mapper"
)

func (u *userUsecase) Store(c context.Context, ureq _dto.UserRequestCreate) (_dto.UserResponse, error) {
	us, err := _mapper.MapUserRequestCreateToUser(ureq)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	res, err := u.userMysqlRepo.Store(c, us)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	return _mapper.MapUserToUserResponse(res), nil
}

func (u *userUsecase) GetByID(c context.Context, id string) (_dto.UserResponse, error) {
	res, err := u.userMysqlRepo.GetByID(c, id)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	return _mapper.MapUserToUserResponse(res), nil
}

func (u *userUsecase) Fetch(c context.Context, pagination _dto.Pagination) (_dto.Pagination, error) {
	res, err := u.userMysqlRepo.Fetch(c, &pagination)
	if err != nil {
		return _dto.Pagination{}, err
	}

	if len(res) != 0 {
		pagination.Rows = _mapper.MapUsersToUserResponses(res)
	}

	return pagination, nil
}

func (u *userUsecase) Update(c context.Context, ureq _dto.UserRequestUpdate) (_dto.UserResponse, error) {
	us, err := u.userMysqlRepo.GetByID(c, ureq.ID)
	if err != nil {
		return _dto.UserResponse{}, err
	}
	usr, err := _mapper.MapUserRequestUpdateToUser(ureq, us)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	res, err := u.userMysqlRepo.Update(c, usr)
	if err != nil {
		return _dto.UserResponse{}, err
	}

	return _mapper.MapUserToUserResponse(res), nil
}

func (u *userUsecase) Delete(c context.Context, id string) error {
	err := u.userMysqlRepo.Delete(c, id)
	if err != nil {
		return err
	}

	return nil
}
