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
		RoleID:   u.RoleID,
		Password: pwd,
	}, nil
}

func MapUserRequestUpdateToUser(u _dto.UserRequestUpdate, us _domain.User) (_domain.User, error) {
	if u.Name != "" {
		us.Name = u.Name
	}

	if u.RoleID != "" {
		us.RoleID = u.RoleID
	}

	return us, nil
}

func MapUserRequestPasswordUpdateToUser(u _dto.UserRequestPasswordUpdate, us _domain.User) (_domain.User, error) {
	pwd, err := _helper.HashPassword(u.NewPassword)
	if err != nil {
		return _domain.User{}, err
	}
	us.Password = pwd

	return us, nil
}

func MapUserToUserResponse(u _domain.User) _dto.UserResponse {
	return _dto.UserResponse{
		ID:     u.ID,
		Name:   u.Name,
		RoleID: u.RoleID,
	}
}

func MapUsersToUserResponses(u []*_domain.User) []_dto.UserResponse {
	res := []_dto.UserResponse{}
	for _, usr := range u {
		res = append(res, MapUserToUserResponse(*usr))
	}

	return res
}
