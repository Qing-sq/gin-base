package controllers

import (
	"awesomeProject3/models"
	"awesomeProject3/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)
type loginResponse struct {
	models.User
	Token string `json:"token"`
}
func UserRegister(c *gin.Context) {
	var form models.User
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err,"接受hi白")
		ResponseValid(c, 500)
	} else {
		if res, er := form.Register(); er != nil {
			fmt.Println(er,"接受hi白1")
			ResponseError(c, 500)
		} else {
			if res.ID == 0 {
				ResponseString(c, 500, "用户名已被注册")
			} else {
				ResponseSuccess(c, res)
			}

		}

	}

}

func UserLogin(c *gin.Context) {

	var form models.User
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)

		ResponseValid(c, 500)
		return
	}
	fmt.Println(form)

	if res := form.Login(); res == nil {
		if regires, er := form.Register(); er != nil {
			fmt.Println(regires)
			ResponseError(c, 500)
		} else {
			if regires.ID == 0 {
				ResponseString(c, 500, "用户名已被注册")
			} else {
				token,errLog := jwt.GenToken(regires.Phone,regires.Name)
				if errLog != nil{
					fmt.Println(errLog.Error())
					ResponseString(c, 500,"token解析失败")
					return
				}
				ResponseSuccess(c, loginResponse{
					*regires,
					token,
				})
			}

		}
	}else{
		token,errLog := jwt.GenToken(res.Phone,res.Name)
		if errLog != nil{
			fmt.Println(errLog.Error())
			ResponseString(c, 500,"token解析失败")
			return
		}
		ResponseSuccess(c,loginResponse{
			*res,
			token,
		})
	}


}
func UserChange(c *gin.Context) {
	fmt.Println(c.GetString("phone"),"获取信息")
	var form models.User
	if err := c.ShouldBind(&form); err != nil {
		ResponseValid(c, 500)
	} else {
		if res := form.Change(); res != nil {
			ResponseSuccess(c, res)
		} else {
			ResponseUserError(c, 500)
		}

	}
}

func GetUserList(c *gin.Context)  {
	fmt.Println(c.Query("phone"),"获取信息")
	if res,err := models.GetUserList(c.Query("phone"));err != nil{
		ResponseError(c,500)
	}else{

		ResponseSuccess(c,res)
	}

}