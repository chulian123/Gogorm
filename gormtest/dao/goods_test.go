package dao

import (
	"testing"
	"time"
)

func TestSaveGoods(t *testing.T) {
	goods := Goods{
		Title:      "毛巾",
		Price:      25.5,
		Stock:      100,
		Type:       0,
		CreateTime: time.Now(),
	}
	SaveGoods(goods)
}

func TestUpdateGoods(t *testing.T) {
	//goods := Goods{
	//	Title:      "橘子",
	//	Price:      25.5,
	//	Stock:      100,
	//	Type:       0,
	//	CreateTime: time.Now(),
	//}
	//GetDB().Save(&goods)
	UpdateGoods()
}

func TestSearch(t *testing.T) {
	SearchSon()
}

func TestDelete(t *testing.T) {
	DeleteData()
}
func TestFind(t *testing.T) {
	FindData()
}

func TestSessionContext(t *testing.T) {
	SessionContext()
}

func TestTransaction(t *testing.T) {
	Transaction()
}

func TestTransactionSave(t *testing.T) {
	//	TransactionSave()
	//FindUserANDGoods()
	//SmartChoose()
	GuanLianCheck()
}
