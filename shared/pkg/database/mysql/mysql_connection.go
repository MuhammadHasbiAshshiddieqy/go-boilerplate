package mysql

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db       = make(map[string]*gorm.DB)
	connList []string
	err      error
)

type MySqlOption struct {
	ConnName string
	Username string
	Password string
	Host     string
	DbName   string
	Port     int
	Role     int
}

type mySqlOption interface {
	Init() error
}

// New - mysql constructor
func New(role, port int, host, db, uname, pwd, connname string) mySqlOption {
	return &MySqlOption{
		ConnName: connname,
		Username: uname,
		Password: pwd,
		Host:     host,
		DbName:   db,
		Port:     port,
		Role:     role,
	}
}

// Init - mysql init
func (s *MySqlOption) Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.Username,
		s.Password,
		s.Host,
		s.Port,
		s.DbName,
	)
	db[s.ConnName], err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name
	}), &gorm.Config{})

	if err != nil {
		errMsg := fmt.Sprintf("MySQL ERROR : error to create %s connection", s.ConnName)
		return errors.New(errMsg)
	}

	connList = append(connList, s.ConnName)

	return err
}

// GetConnection - return mysql connection
func GetConnection(connname string) *gorm.DB {
	return db[connname]
}

// GetConnectionList - return mysql connection names
func GetConnectionList() []string {
	return connList
}
