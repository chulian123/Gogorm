package main

import (
	"github.com/gin-gonic/gin"
	"gormtest/router"
	"log"
)

// 一般流程：
// 1.新建连接数据库初始化文件 /dao/gorm.go 主要用作连接数据等参数。
// 2.定义数据模型操作文件，如操作用户相关 /dao/user.go 包含模型，常用数据 新建/编辑/删除/查询 等相关定义。
// 3.定义数据查询接口 ,如用户相关查询 /api/user.go
// 4.定义对外访问路由，/api/router.go
func main() {
	//启动了一个gin的框架
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
