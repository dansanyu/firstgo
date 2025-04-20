package main

import (
	db2 "firstgo/internal/db"
	"firstgo/internal/router"
)

func main() {
	// 数据库
	db := db2.ConnectDB()
	r := router.SetupRouter(db)
	err := r.Run(":8080")
	if err != nil {
		panic("启动失败")
	}
}
