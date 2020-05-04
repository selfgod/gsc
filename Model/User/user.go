package User

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gsc/Base"
)

type User struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(100);not null;unique_index"`
	Pass   string `json:"pass" gorm:"type:varchar(100)"`
	Token  string `json:"token" gorm:"type:varchar(100);not null;unique_index"`
	Expire int    `json:"expire" gorm:"type:int(10)"`
}

func InitUser() {
	Base.DB.AutoMigrate(new(User))
}

func (u *User) Add() (err error) {
	err = Base.DB.Debug().Create(u).Error
	return
}

func (u *User) NameExists() (ret bool, err error) {
	ret = false
	var c int8
	err = Base.DB.Model(u).Where("name = ?", u.Name).Count(&c).Error
	if c > 0 {
		ret = true
	}
	return
}
func (u *User) Login() (ret bool, err error) {
	ret = false
	var c int8
	err = Base.DB.Model(u).Debug().Where("name = ? and pass=?", u.Name, u.Pass).Count(&c).Error
	if c > 0 {
		ret = true
	}
	return
}

func (u *User) SetToken(token string, expire int64) (err error) {
	/*updata := &User{
		Token:token,
		Expire:expire,
	}*/
	err = Base.DB.Model(u).Debug().Updates(gin.H{"token": token, "expire": expire}).Error
	return
}

func GetUsers() (ret bool) {
	return true
}
