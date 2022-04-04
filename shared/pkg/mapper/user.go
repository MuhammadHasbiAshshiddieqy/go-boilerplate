package mapper

import (
	_domain "microservice/shared/domain"
	_dto "microservice/shared/dto"

	_helper "microservice/shared/pkg/helper"
)

func MapUserRequestCreateToUser(u _dto.UserRequestCreate) (_domain.User, error) {
	pwd, err := _helper.HashPassword(u.Password)
	if err != nil {
		return _domain.User{}, err
	}
	return _domain.User{
		Name:     u.Name,
		Password: pwd,
	}, nil
}

func MapUserToUserResponse(u _domain.User) _dto.UserResponse {
	return _dto.UserResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
