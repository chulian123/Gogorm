package api

import (
	"github.com/gin-gonic/gin"
	"gormtest/dao"
	"time"
)

func SaveUser(c *gin.Context) {
	user := &dao.User{
		Username:   "zhangsan",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	dao.Save(user)
	c.JSON(200, user)
}

func saveh(c *gin.Context) {
	user := &dao.User{
		Username:   "zhangsan",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	dao.SaveH(user)
	c.JSON(200, user)
}

func GetUser(c *gin.Context) {
	user := dao.GetById(1)
	c.JSON(200, user)
}

func GetUserSecond(c *gin.Context) {
	user := dao.GetAll()
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	dao.UpdateUser(3)
	user := dao.GetById(3)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	dao.DeleteUser(3)
	user := dao.GetById(3)
	c.JSON(200, user)
}
