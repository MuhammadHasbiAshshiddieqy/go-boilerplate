package mysqlrepository

import (
	"context"
	_domain "microservice/shared/domain"
	_dto "microservice/shared/dto"
	_helper "microservice/shared/pkg/helper"
)

func (u *userMysqlRepository) GetByID(c context.Context, id string) (_domain.User, error) {
	us := _domain.User{}
	if err := u.Orm.First(&us, "id = ?", id).Error; err != nil {
		return us, err
	}

	return us, nil
}

func (u *userMysqlRepository) GetByCondition(c context.Context, us _domain.User) (_domain.User, error) {
	usr := _domain.User{}
	if err := u.Orm.Where(&us).First(&usr).Error; err != nil {
		return usr, err
	}

	return usr, nil
}

func (u *userMysqlRepository) Fetch(c context.Context, pagination *_dto.Pagination) ([]*_domain.User, error) {
	var usrs []*_domain.User
	if err := u.Orm.Scopes(_helper.Paginate(usrs, pagination, u.Orm)).Find(&usrs); err != nil {
		return usrs, nil
	}

	return usrs, nil
}
