package http

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	_domain "microservice/shared/domain"
// 	_mocks "microservice/shared/domain/mocks"
// 	_dto "microservice/shared/dto"

// 	"github.com/stretchr/testify/assert"
// )

// func TestUserHttpHandler_Store(t *testing.T) {
// 	type args struct {
// 		user    _domain.User
// 		httpReq _dto.UserRequestCreate
// 	}
// 	type wants struct {
// 		user    _domain.User
// 		httpRes _dto.UserResponse
// 	}
// 	pwd := "password_test"
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wants   wants
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "success_store_user",
// 			args: args{
// 				user: _domain.User{
// 					Name:     "Hasbi",
// 					Password: pwd,
// 					Email:    "hasbi@gmail.com",
// 					RoleID:   "nanoid_role",
// 				},
// 				httpReq: _dto.UserRequestCreate{
// 					Name:     "Hasbi",
// 					Email:    "hasbi@gmail.com",
// 					Password: pwd,
// 					RoleID:   "nanoid_role",
// 				},
// 			},
// 			wants: wants{
// 				user: _domain.User{
// 					Base: _domain.Base{
// 						ID:        "nanoid",
// 						CreatedAt: time.Now(),
// 						UpdatedAt: time.Now(),
// 					},
// 					Name:     "Hasbi",
// 					Password: pwd,
// 					Email:    "hasbi@gmail.com",
// 					RoleID:   "nanoid_role",
// 				},
// 				httpRes: _dto.UserResponse{
// 					ID:     "nanoid",
// 					Name:   "Hasbi",
// 					RoleID: "nanoid_role",
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "failed_store_user",
// 			args: args{
// 				user:    _domain.User{},
// 				httpReq: _dto.UserRequestCreate{},
// 			},
// 			wants: wants{
// 				user:    _domain.User{},
// 				httpRes: _dto.UserResponse{},
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			assert := assert.New(t)
// 			uu := &_mocks.UserUsecase{}
// 			NewUserHttpHandler()
// 			if !tt.wantErr {
// 				uu.On("Store", context.Background(), tt.args.user).Return(tt.wants.user, nil)
// 			} else {
// 				uu.On("Store", context.Background(), tt.args.user).Return(tt.wants.user, errors.New("failed to store user"))
// 			}
// 			res, err := uu.Store(context.Background(), tt.args.httpReq)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UserUsecase.Store() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			assert.Equal(tt.wants.httpRes, res)
// 		})
// 	}
// }
