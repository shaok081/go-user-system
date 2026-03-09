package main

import (
	"fmt"

	"go-user-system/config"
	"go-user-system/model"
	"go-user-system/router"
)

func main() {
	config.InitDB()

	err := config.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic("数据表迁移失败")
	}

	r := router.SetupRouter()

	fmt.Println("服务器启动成功：http://localhost:8080")
	err = r.Run(":8080")
	if err != nil {
		panic("服务器启动失败")
	}
}
