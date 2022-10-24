package models

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	ID     int `json: "id" gorm:"primaryKey"`
	UserID uint
	Produk []*Produk `gorm:"many2many:transaksi_produk;"`
}
