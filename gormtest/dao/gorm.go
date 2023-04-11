package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB //_db不希望被其他人使用

func init() {
	//连接数据库的操作
	//配置MySQL连接参数
	username := "root"            //账号
	password := "Chulianpzj12345" //密码
	host := "172.0.0.1"           //数据库地址，可以是Ip或者域名
	port := 5000                  //数据库端口
	Dbname := "GormTest"          //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	//倒入sql配置 连接数据库
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //gorm的配置 设置日记级别
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	db, err := _db.DB()
	db.SetMaxIdleConns(100)
	db.SetConnMaxIdleTime(20)
}

func GetDB() *gorm.DB {
	return _db
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//	q := r.URL.Query()
		//	page, _ := strconv.Atoi(q.Get("page"))
		//	if page == 0 {
		//		page = 1
		//	}
		//
		//	pageSize, _ := strconv.Atoi(q.Get("page_size"))
		//	switch {
		//	case pageSize > 100:
		//		pageSize = 100
		//	case pageSize <= 0:
		//		pageSize = 10
		//	}

		offset := (page - 1) * pageSize
		return db.Limit(pageSize).Offset(offset)
	}
}
