package mysqlrepository

import (
	"context"
	_domain "microservice/shared/domain"
	"time"
)

func (u *userMysqlRepository) Store(c context.Context, us _domain.User) (_domain.User, error) {
	if err := u.Orm.Create(&us).Error; err != nil {
		return us, err
	}

	return us, nil
}

func (u *userMysqlRepository) Update(c context.Context, us _domain.User) (_domain.User, error) {
	if err := u.Orm.Updates(&us).Error; err != nil {
		return us, err
	}

	return us, nil
}

func (u *userMysqlRepository) Delete(c context.Context, id string) error {
	if err := u.Orm.Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
