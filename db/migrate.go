package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yuuis/PersonalDataRepository/models/basic"
	"github.com/yuuis/PersonalDataRepository/models/health"
	"github.com/yuuis/PersonalDataRepository/models/location"
	"github.com/yuuis/PersonalDataRepository/models/registration_information"
	"log"
	"os"
)

func main() {
	// mysql
	mysqlDbName := os.Getenv("MYSQL_DATABASE")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlRootPassword := os.Getenv("MYSQL_ROOT_PASSWORD")

	mysql, err := gorm.Open("mysql", "root:"+mysqlRootPassword+"@tcp("+mysqlHost+":3306)/"+mysqlDbName+"?charset=utf8mb4&parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	mysql.CreateTable(&basic.Basic{}, &location.Location{}, &health.Health{}, &registration_information.RegistrationInformation{})

	// mongo
	// TODO: mongoのmigration
}
