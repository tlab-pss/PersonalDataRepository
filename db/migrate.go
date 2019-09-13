package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yuuis/PersonalDataRepository/model"
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

	mysql.CreateTable(&model.Basic{}, &model.Location{})

  // mongo
  // TODO: mongoのmigration
}
