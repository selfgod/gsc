package Base

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDb(config Db) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", config.Username, config.Password, config.Host, config.Name)
	fmt.Printf("dsn:%#v\n", dsn)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	err = DB.DB().Ping()
	if err != nil {
		panic(err.Error())
	}
	return
}
