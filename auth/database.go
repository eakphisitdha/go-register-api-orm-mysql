package auth

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {
	//---conect to mysql-server
	dsn := "root:P@ssw0rd@tcp(192.168.1.89:3306)/member?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&User{})
	//---

}
