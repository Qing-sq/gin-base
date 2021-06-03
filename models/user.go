package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Phone    int64   `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" `
}

func (user *User) UserInsert() (*User, error) {
	err := DB.Create(&user).Error
	return user, err
}
func GetUserList(phone string) ([]*User, error) {
	var users []*User
	if phone == ""{
		err := DB.Find(&users).Error
		return users, err

	}

	err := DB.Find(&users,"phone = ?",phone).Error

	return users, err
}
func (user *User) GetUser(isCheck bool) (*User, error) {
	fmt.Println(user, "进入查询", user.Phone)
	var err error
	if isCheck {
		err = DB.First(&user, "phone = ?", user.Phone).Error
	} else {
		err = DB.First(&user, "phone = ? AND password = ?", user.Phone, user.Password).Error
	}
	return user, err
}

func (user *User) CheckPhone() *User {
	DB.Where("phone = ?", user.Phone).First(&user)
	return user
}
func (user *User) GetUserById() (*User, error) {
	err := DB.First(&user, "id = ?", user.ID).Error
	return user, err
}

func (user *User) Register()(*User,error) {
	fmt.Println("获取数据", user)
	count := 0
	err := DB.Model(&User{}).Where("phone = ?", user.Phone).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return user, err
	}
	if err = DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, err
}

func (user *User) Login()*User  {
	if bool := user.check();bool{
		return user
	}	else{
		return nil
	}

}

func (user *User) Change()*User  {
	if bool := user.check();bool{
		DB.Save(&user)
		return user
	}else{
		return nil
	}

}

func (user *User) check() bool{
	count := 0
	err := DB.Where("phone = ? AND password = ?",user.Phone,user.Password).First(&user).Count(&count).Error
	if err != nil  || count <= 0{
			return  false
	}
	return  true
}
