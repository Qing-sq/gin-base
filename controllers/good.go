package controllers

import (
	"awesomeProject3/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddGood(c *gin.Context) {
	var good *models.Good
	if err := c.ShouldBind(&good); err != nil {
		ResponseError(c, 500)
		fmt.Println(err.Error())
		return
	}
	res, err := good.GoodInsert()
	if err != nil {
		ResponseError(c, 500)
		return
	}
	ResponseSuccess(c, res)
}
func GetGood(c *gin.Context)  {

}
