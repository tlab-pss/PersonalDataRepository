package infrastructures

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

func OpenMysql() (*gorm.DB, error) {
	mysqlDbName := os.Getenv("MYSQL_DATABASE")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlRootPassword := os.Getenv("MYSQL_ROOT_PASSWORD")

	return gorm.Open("mysql", "root:"+mysqlRootPassword+"@tcp("+mysqlHost+":3306)/"+mysqlDbName+"?charset=utf8mb4&parseTime=true")
}
