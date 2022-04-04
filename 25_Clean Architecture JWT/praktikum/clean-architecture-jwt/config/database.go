package config

import (
	"fmt"
	"go-ca/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	var Config = struct {
		DB_Username string
		DB_Password string
		DB_Port     string
		DB_Host     string
		DB_Name     string
	}{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_test",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Config.DB_Username,
		Config.DB_Password,
		Config.DB_Host,
		Config.DB_Port,
		Config.DB_Name,
	)

	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
	)
}
