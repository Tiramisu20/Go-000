package data

import (
	"github.com/jinzhu/gorm"
	"project/internal/kit/sql"
)

type User struct {
	gorm.Model
	Id string
	Name string
}

func (User) FindAll() (user []User) {
	// TODO Generate Project level DB connection
	db := sql.DbConn("root", "root", "127.0.0.1", "User", 3306)
	defer db.Close()
	db.Find(&user)
	return
}