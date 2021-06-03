package main

import (
	"log"
	"os/exec"

)

func main()  {
	checkExe2()
}

func checkExe2() {
	cmd := exec.Command("cmd.exe", "/c", "start http://baidu.com && code .")
	err := cmd.Run()
	if err != nil {
		log.Println("启动失败:", err)
	} else {
		log.Println("启动成功!")
	}
}