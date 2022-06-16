package mysql

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_config "microservice/shared/config"
	_domain "microservice/shared/domain"
)

var db *gorm.DB
var err error

// Init - mysql init
func Init(DbUser, DbPassword string) error {
	conf := _config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPassword,
		conf.Sql.Host,
		conf.Sql.Port,
		conf.Sql.DB,
	)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name
	}), &gorm.Config{})

	if err != nil {
		return errors.New("MySQL Connection Error")
	}
	err = db.AutoMigrate(
		&_domain.User{},
		&_domain.Role{},
	)

	return err
}

// DbManager - return db connection
func MySQLManager() *gorm.DB {
	return db
}
