package models

import (
	u "authen-api/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
}

func (user *User) Validate() (map[string] interface{}, bool) {
		if user.Name == "" {
			return u.Message(false, "User name is null"), false
		}

		if user.Email == "" {
			return u.Message(false, "Email is null"), false
		} else {
			tmp_user := &User{}
			if !db.Where("email = ?", user.Email).First(&tmp_user).RecordNotFound(){
				return u.Message(false, "User email exists"), false
			}
		}

		return u.Message(true, "OK"), true
}


func (user *User) Create() (map[string] interface{}) {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	DB().Create(user)

	run := u.Message(true, "success")
	run["user"] = user
	return run
}

func GetUser(id uint) (*User){
	user := &User{}
	err := DB().Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return user
}

func GetListUsers() ([]*User){
	users := make([]*User, 0)
	err := DB().Table("users").Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return users
}
