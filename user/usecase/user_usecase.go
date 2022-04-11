package usecase

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	_domain "microservice/shared/domain"
	_dto "microservice/shared/dto"
	_helper "microservice/shared/pkg/helper"
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

func (u *userUsecase) Login(c *fiber.Ctx, ureq _dto.UserRequestLogin) (_dto.UserResponseToken, error) {
	us, err := u.userRepo.GetByCondition(c, _domain.User{
		Name: ureq.Name,
	})
	if err != nil {
		return _dto.UserResponseToken{}, _domain.ErrUnauthorized
	}
	ok := _helper.CheckPasswordHash(ureq.Password, us.Password)
	if !ok {
		return _dto.UserResponseToken{}, _domain.ErrUnauthorized
	}

	ts, err := _helper.CreateToken(us.ID)
	if err != nil {
		return _dto.UserResponseToken{}, _domain.ErrUnauthorized
	}

	saveErr := _helper.CreateAuth(us.ID, ts)
	if saveErr != nil {
		return _dto.UserResponseToken{}, _domain.ErrUnauthorized
	}

	return _dto.UserResponseToken{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}, nil
}

func (u *userUsecase) Refresh(c *fiber.Ctx, ureq _dto.UserRequestRefresh) (_dto.UserResponseToken, error) {
	//verify the token
	// os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	token, err := jwt.Parse(ureq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errMsg := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(errMsg)
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		return _dto.UserResponseToken{}, errors.New("refresh token expired")
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return _dto.UserResponseToken{}, err
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			return _dto.UserResponseToken{}, errors.New("failed to get refresh_uuid")
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			return _dto.UserResponseToken{}, errors.New("failed to get user_id")
		}
		//Delete the previous Refresh Token
		deleted, delErr := _helper.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			return _dto.UserResponseToken{}, _domain.ErrUnauthorized
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := _helper.CreateToken(userId)
		if createErr != nil {
			return _dto.UserResponseToken{}, createErr
		}
		//save the tokens metadata to redis
		saveErr := _helper.CreateAuth(userId, ts)
		if saveErr != nil {
			return _dto.UserResponseToken{}, saveErr
		}

		return _dto.UserResponseToken{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}, nil

	} else {
		return _dto.UserResponseToken{}, errors.New("refresh token expired")
	}
}

func (u *userUsecase) Logout(c *fiber.Ctx, metadata *_dto.AccessDetails) error {
	delErr := _helper.DeleteTokens(metadata)
	if delErr != nil {
		return errors.New("failed to delete token")
	}
	return nil
}

func (u *userUsecase) ResetPassword(c *fiber.Ctx, metadata *_dto.AccessDetails, ureq _dto.UserRequestPasswordUpdate) error {
	us, err := u.userRepo.GetByID(c, metadata.UserId)
	if err != nil {
		return err
	}

	ok := _helper.CheckPasswordHash(ureq.Password, us.Password)
	if !ok {
		return _domain.ErrUnauthorized
	}

	usr, err := _mapper.MapUserRequestPasswordUpdateToUser(ureq, us)
	if err != nil {
		return err
	}

	_, err = u.userRepo.Update(c, usr)
	if err != nil {
		return err
	}

	return nil
}
