package configs

import (
	"fmt"
	"go-fiber/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLConn() *gorm.DB {

	configDB := map[string]string{
		"DB_Username": Config("DB_USERNAME"),
		"DB_Password": Config("DB_PASSWORD"),
		"DB_Port":     Config("DB_PORT"),
		"DB_Host":     Config("DB_ADDRES"),
		"DB_Name":     Config("DB_NAME"),
	}
	fmt.Println(configDB)
	fmt.Println("Server running well...")

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

	role := []models.Role{{Name: "Administrator"}, {Name: "User"}}

	var value = make([]models.Role, 0)
	for _, v := range role {
		value = append(value, v)
	}

	db.AutoMigrate(&models.Diary{}, &models.User{}, value)
	return db

}
