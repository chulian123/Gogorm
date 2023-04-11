package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

// Goods 对应数据库结构
type Goods struct {
	Id         int
	Title      string
	Price      float64
	Stock      int
	Type       int
	CreateTime time.Time
}

//比如有一个用户属性表，查询用户的时候需要将其的性别和年龄查询出来：

type UserProfile struct {
	ID     int64
	UserId int64
	Sex    int
	Age    int
}

func (u UserProfile) TableName() string {
	return "user_profiles"
}

// TableName 数据库表
func (v Goods) TableName() string {
	return "goods"
}

func (*Goods) BeforeCreate(tx *gorm.DB) error {
	log.Println("before create ...")
	return nil
}
func (*Goods) BeforeSave(tx *gorm.DB) error {
	log.Println("BeforeSave create ...")
	return nil
}
func (*Goods) AfterCreate(tx *gorm.DB) error {
	log.Println("After create ...")
	return nil
}
func (*Goods) AfterSave(tx *gorm.DB) error {
	log.Println("AfterSave create ...")
	return nil
}
func (*Goods) BeforeUpdate(tx *gorm.DB) error {
	log.Println("BeforeUpdate create ...")
	return nil
}
func (*Goods) AfterUpdate(tx *gorm.DB) error {
	log.Println("AfterUpdate create ...")
	return nil
}
func (*Goods) BeforeDelete(tx *gorm.DB) error {
	log.Println("BeforeDelete create ...")
	return nil
}
func (*Goods) AfterDelete(tx *gorm.DB) error {
	log.Println("AfterDelete create ...")
	return nil
}

// 保存
func SaveGoods(goods Goods) {
	GetDB().Save(&goods)
}

// 更新全部
func UpdateGoods() {
	var goods Goods
	var id int = 1
	//GetDB().Where("id =?", id).Take(&goods)
	//相关的
	goods.Price = 9.3
	//1
	//err := GetDB().Model(&Goods{}).Where("id", id).Update("title", "111").Error//更新单列
	//2
	//更新多列可以用map或者结构体
	err := GetDB().Model(&Goods{}).Where("id", id).Updates(Goods{
		Title: "橘子",
		Price: 20.5,
	}).Error
	if err != nil {
		log.Fatalln(err)
	}
}

// SearchSon 子查询更新
func SearchSon() {
	goods := Goods{}
	GetDB().Where("id = ?", 3).Take(&goods)
	GetDB().Model(&goods).Update("title", GetDB().Model(&User{}).Select("username").Where("id=?", 2))
}

// DeleteData 删除操作
func DeleteData() {
	//为了避免共用db导致的一些问题，gorm提供了会话模式，通过新建session的形式，将db的操作分离，互不影响。
	db := GetDB().Session(&gorm.Session{})
	//1
	//db.Delete(&Goods{},2)//默认根据id来删除
	db.Where("id=?", 3).Delete(&Goods{})
}

func FindData() {
	db := GetDB()
	var goods Goods
	var titles []string

	//1
	db.Where("id =?", 4).Take(&goods)
	//2
	db.Where("id =?", 4).First(&goods) //排序后第一个数据
	//3
	db.Where("id =?", 4).Last(&goods) //倒叙后第一个数据
	//4
	err := db.Where("id =?", 4).Take(&goods)
	fmt.Println(err)
	db.Model(&Goods{}).Pluck("title", &titles)
	fmt.Println(titles)
	//5
	//设置指定字段返回数据
	db.Select("id", "title").Find(&goods)

	//6
	//聚合函数
	var total int64
	db.Model(&Goods{}).Select("count(*) as total").Pluck("total", &total)
	fmt.Println(total)
	//7
	// limit 1,1 分页功能
	db.Order("create_time desc").Limit(2).Offset(2).Find(&goods)
	fmt.Println(goods)
	//8
	//返回查询匹配的行数
	db.Model(Goods{}).Count(&total)
	fmt.Println(total)
	//9
	//统计每个商品分类下面有多少个商品
	//定一个Result结构体类型，用来保存查询结果
	type Result struct {
		Type  int
		Total int
	}
	var results []Result
	//等价于: SELECT type, count(*) as  total FROM `goods` GROUP BY type HAVING (total > 0)
	db.Model(Goods{}).Select("type, count(*) as  total").Group("type").Having("total > 0").Scan(&results)
	fmt.Println(results)
	//scan类似Find都是用于执行查询语句，然后把查询结果赋值给结构体变量，区别在于scan不会从传递进来的结构体变量提取表名.
	//这里因为我们重新定义了一个结构体用于保存结果，但是这个结构体并没有绑定goods表，所以这里只能使用scan查询函数。

	//10
	session := db.Where("type=?", 0).Session(&gorm.Session{NewDB: true})
	session.Take(&goods)

}

func SessionContext() {
	db := GetDB()
	timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second)
	tx := db.Session(&gorm.Session{Context: timeoutCtx})
	var user User
	tx.First(&user) // 带有 context timeoutCtx 的查询操作
	//	tx.Model(&user).Update("role", "admin") // 带有 context timeoutCtx 的更新操作
	go handler(timeoutCtx)
}

func handler(ctx context.Context) {
	select {

	case <-ctx.Done():
		fmt.Println("超时了")
		return
	default:
		fmt.Println("default")
	}
}

// Transaction 事物操作
func Transaction() {
	//自动事物操作
	//db := GetDB().Session(&gorm.Session{})
	//// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	//err := db.Transaction(func(tx *gorm.DB) error {
	//	//在这里进行事物的操作
	//	goods := &Goods{
	//		Title:      "苹果",
	//		Price:      2.3,
	//		Stock:      200,
	//		Type:       1,
	//		CreateTime: time.Now(),
	//	}
	//	if err := tx.Create(&goods).Error; err != nil {
	//		return err
	//	}
	//	//在这里进行事物的操作
	//	goods2 := &Goods{
	//		Title:      "苹果2",
	//		Price:      2.3,
	//		Stock:      200,
	//		Type:       1,
	//		CreateTime: time.Now(),
	//	}
	//	if err := tx.Create(&goods2).Error; err != nil {
	//		return err
	//	}
	//	return nil
	//})
	//log.Println("Transaction err", err)
	//
	//手动开启事物
	db := GetDB().Session(&gorm.Session{})
	tx := db.Begin()
	//在这里进行事物的操作
	goods := &Goods{
		Title:      "苹果",
		Price:      2.3,
		Stock:      200,
		Type:       1,
		CreateTime: time.Now(),
	}
	if err := tx.Create(&goods).Error; err != nil {
		tx.Rollback() //报错回滚事物
		return
	}
	//在这里进行事物的操作
	goods2 := &Goods{
		Title:      "苹果2",
		Price:      2.3,
		Stock:      200,
		Type:       1,
		CreateTime: time.Now(),
	}
	if err := tx.Create(&goods2).Error; err != nil {
		tx.Rollback() //报错回滚事物
		return
	}
	tx.Commit()
}

func c() {
	//// 开启事务
	//tx := db.Begin()
	//
	////在事务中执行数据库操作，使用的是tx变量，不是db。
	////库存减一
	////等价于: UPDATE `goods` SET `stock` = stock - 1  WHERE `goods`.`id` = '2' and stock > 0
	////RowsAffected用于返回sql执行后影响的行数
	//rowsAffected := tx.Model(&goods).Where("stock > 0").Update("stock", gorm.Expr("stock - 1")).RowsAffected
	//if rowsAffected == 0 {
	//	//如果更新库存操作，返回影响行数为0，说明没有库存了，结束下单流程
	//	//这里回滚作用不大，因为前面没成功执行什么数据库更新操作，也没什么数据需要回滚。
	//	//这里就是举个例子，事务中可以执行多个sql语句，错误了可以回滚事务
	//	tx.Rollback()
	//	return
	//}
	//err := tx.Create(保存订单).Error
	//
	////保存订单失败，则回滚事务
	//if err != nil {
	//	tx.Rollback()
	//} else {
	//	tx.Commit()
	//}
}

// 嵌套事物操作
func SecondTransaction() {
	//GetDB().Transaction(func(tx *gorm.DB) error {
	//	tx.Create(&user1)
	//
	//	tx.Transaction(func(tx2 *gorm.DB) error {
	//		tx2.Create(&user2)
	//		return errors.New("rollback user2") // Rollback user2
	//	})
	//
	//	tx.Transaction(func(tx2 *gorm.DB) error {
	//		tx2.Create(&user3)
	//		return nil
	//	})
	//
	//	return nil
	//})

	// Commit user1, user3
}

// 事物保存点
func TransactionSave() {
	//GORM 提供了 SavePoint、Rollbackto 方法，来提供保存点以及回滚至保存点功能，例如：

	db := GetDB().Session(&gorm.Session{})
	tx := db.Begin()
	//在这里进行事物的操作
	goods := &Goods{
		Title:      "苹果",
		Price:      2.3,
		Stock:      200,
		Type:       1,
		CreateTime: time.Now(),
	}
	if err := tx.Create(&goods).Error; err != nil {
		tx.Rollback() //报错回滚事物
		return
	}
	tx.SavePoint("sp1") //设置保存点1

	//在这里进行事物的操作
	goods2 := &Goods{
		Title:      "苹果2",
		Price:      2.3,
		Stock:      200,
		Type:       1,
		CreateTime: time.Now(),
	}
	if err := tx.Create(&goods2).Error; err != nil {
		tx.Rollback() //报错回滚事物
		return
	}
	tx.RollbackTo("sp1")
	tx.Commit()
}

// FindUserANDGoods 1. scope
// 作用域允许你复用通用的逻辑，这种共享逻辑需要定义为类型func(*gorm.DB) *gorm.DB。
//
// 例子：
func FindUserANDGoods() {
	db := GetDB().Session(&gorm.Session{})
	var user []User
	db.Scopes(Paginate(2, 3)).Find(&user)
	var goods []Goods
	db.Scopes(Paginate(2, 3)).Find(&goods)
	fmt.Println(user)
}

// 智能选择字段
type APIUser struct {
	ID   uint
	Name string
}

func SmartChoose() {
	// 查询时会自动选择 `id`, `name` 字段
	GetDB().Model(&User{}).Limit(10).Scan(&APIUser{})
	// SELECT `id`, `name` FROM `users` LIMIT 10
}

// 子查询
func SonSearch() {
	//GetDB().Where("amount > (?)", GetDB().Table("orders").Select("AVG(amount)")).Find(&orders)
	// SELECT * FROM "orders" WHERE amount > (SELECT AVG(amount) FROM "orders");
	//	from子查询

	GetDB().Table("(?) as u", GetDB().Model(&User{}).Select("name", "age")).Where("age = ?", 18).Find(&User{})
	// SELECT * FROM (SELECT `name`,`age` FROM `users`) as u WHERE `age` = 18

	subQuery1 := GetDB().Model(&User{}).Select("name")
	subQuery2 := GetDB().Model(&User{}).Select("name")
	GetDB().Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&User{})
}

// GuanLianCheck 关联操作
func GuanLianCheck() {
	db := GetDB().Session(&gorm.Session{})
	var user = User{
		Username:   "ms",
		Password:   "ms",
		CreateTime: time.Now().UnixMilli(),
		UserProfile: UserProfile{
			Sex: 0,
			Age: 20,
		},
	}
	db.Create(&user)
}
