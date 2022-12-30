package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var (
	db * gorm.DB
)




func Connect() *gorm.DB{
	dsn:="root:Ak@sh2000@tcp(localhost)/shoppingcart?charset=utf8&parseTime=True&loc=Local"
	d, err:= gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!= nil{
		panic(err)		
	}
	db=d
	return d
}

func TestDBConnect() *gorm.DB{
	dsn:="root:Ak@sh2000@tcp(localhost)/testdb?charset=utf8&parseTime=True&loc=Local"
	d, err:= gorm.Open(mysql.Open(dsn),&gorm.Config{})
	// d.Logger.LogMode(true)
	if err!= nil{
		panic(err)		
	}
	db=d
	return d
}

func TestDBFree(test_db *gorm.DB) error {
	sqlDB,_:=test_db.DB()
	err:=sqlDB.Close()
	return err
}

func GetDB() *gorm.DB{
	return db
}