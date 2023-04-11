package dao

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type User struct {
	ID          int64
	Username    string `gorm:"column:username"`
	Password    string `gorm:"column:password"`
	CreateTime  int64  `gorm:"column:createtime"`
	Admin       bool   `gorm:"-"` //如果不加-就会报错  Unknown column 'admin' 可以让他忽略当成字段来实现
	CreatedAt   time.Time
	UserProfile UserProfile
}

func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

// 动态绑定表名
func UserTable(user User) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if user.Admin {
			return tx.Table("admin_users")
		}
		return tx.Table("user")
	}

}

func Save(user *User) {
	user.Admin = true
	//err := DB.Scopes(UserTable(*user)).Create(user)//第一种定义动态表名
	//err := DB.Table("admin_users").Create(user).Error//第二种定义动态表名
	result := GetDB().Create(user)
	affected := result.RowsAffected //获取受影响的行数
	fmt.Println("受影响的行数", affected)
	err := result.Error
	if err != nil {
		log.Println("insert fail : ", err)
	}
}

// SaveH 指定行列来插入数据
func SaveH(user *User) {
	//1
	//result := GetDB().Select("username", "password").Create(user) //select函数是指定插入的字段
	//2
	//result := GetDB().Omit("username").Create(user)// Omit函数是去掉某个字段 然后插入剩下全部字段
	//3
	//var users = []User{{Username: "jinzhu1"}, {Username: "jinzhu2"}, {Username: "jinzhu3"}}//批量插入的实现
	//result := GetDB().Create(users)
	//4
	//var users = []User{{Username: "jinzhu1"}, {Username: "jinzhu2"}, {Username: "jinzhu3"}} //批量插入的实现
	//result := GetDB().CreateInBatches(users, 2)                                             //把多个数据列 按2个2个插入
	//5
	result := GetDB().Model(&User{}).Create(map[string]interface{}{
		"username": "jinzhu",
		"password": clause.Expr{SQL: "md5(?)", Vars: []interface{}{"123456"}}, //md5加密
	})
	//6
	//原生的sql语句插入
	GetDB().Exec("insert into users (username,password,createtime) values (?,?,?)", user.Username, user.Password, user.CreateTime)
	affected := result.RowsAffected //获取受影响的行数
	fmt.Println("受影响的行数", affected)
	err := result.Error
	if err != nil {
		log.Println("insert fail : ", err)
	}
}

// GetById 根据id来查询id
func GetById(id int64) User {
	var user User
	//first函数 等于查询第一个 == limit 1 sql语句
	err := GetDB().Where("id=?", id).First(&user).Error
	if err != nil {
		log.Println("get user by id fail : ", err)
	}
	return user
}

// GetAll 查询所有数据
func GetAll() []User {
	var users []User
	err := GetDB().Find(&users)
	if err != nil {
		log.Println("get users  fail : ", err)
	}
	return users
}

// UpdateUser 更新数据
func UpdateUser(id int64) {
	//model负责把表名传入 where是条件 Update是修改条件
	err := GetDB().Model(&User{}).Where("id =?", id).Update("username", "test")
	if err != nil {
		log.Println("UpdateUser  fail : ", err)
	}
}

// DeleteUser 删除用户根据id
func DeleteUser(id int64) {
	//Delete 要把删除的表给传入进去
	err := GetDB().Where("id =?", id).Delete(&User{})
	if err != nil {
		log.Println("UpdateUser  fail : ", err)
	}
}
