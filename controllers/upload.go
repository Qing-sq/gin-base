package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	time2 "time"
)

func Upload(c *gin.Context)  {
	file,_ := c.FormFile("file")
	time  := time2.Now().Unix()
	timeStr := fmt.Sprintf("%d",time)
	c.SaveUploadedFile(file,"static/"+file.Filename + timeStr)
	fmt.Println(file)
}
