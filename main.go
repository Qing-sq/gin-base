package main

import (
	"awesomeProject3/conf"
	"awesomeProject3/models"
	"awesomeProject3/router"
	"fmt"
)
func main() {
	//加载配置
	if err := conf.Init();err != nil{
		fmt.Println("setting fail")
	}
	models.InitDB()
 	defer 	models.DB.Close()
	//if err := redis.Main();err != nil{
	//	fmt.Println("redis fail")
	//}

	r := router.NewRouter()
	r.Run(":3001")

}
