package configs

import (
	"fmt"
	"go-fiber/app/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLConn() *gorm.DB {
	configDB := map[string]string{
		"DB_Username": os.Getenv("DB_USERNAME"),
		"DB_Password": os.Getenv("DB_PASSWORD"),
		"DB_Port":     os.Getenv("DB_PORT"),
		"DB_Host":     os.Getenv("DB_ADDRES"),
		"DB_Name":     os.Getenv("DB_NAME"),
	}
	fmt.Println(configDB)

	openConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])
	db, e := gorm.Open(mysql.Open(openConnection), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	db.AutoMigrate(&models.Diary{}, &models.User{})

	return db

}
