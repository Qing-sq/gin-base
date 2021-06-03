package models

import "github.com/jinzhu/gorm"

type Good struct {
	gorm.Model
	Name string `json:"name"`
	SubName string `json:"subName"`
	Price int `json:"price`
	ImgUrl string `json:"imgUrl"`
	Socket int `json:"socket"`
	Selled int `json:"selled"`
}

func (good *Good) GoodInsert()(*Good,error)  {
	err := DB.Create(&good).Error
	return good, err
}

func (good *Good) GoodGet()()  {

}