package database

import (
	// "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	DB = ConnectDatabase()
	return DB
}

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/finalprojects?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Produk{})
	// db.AutoMigrate(&models.Transaksi{})
	// db.AutoMigrate(&models.Keranjang{})
	// DB = db
	// return DB
	DB = db
	return DB

}
