package models

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	Id      int `json: "id" gorm:"primaryKey"`
	UserID  uint
	Produks []*Produk `gorm:"many2many:transaksi_produk;"`
}

func BuatTransaksi(db *gorm.DB, newTransaksi *Transaksi, userId uint, produks []*Produk) (err error) {
	newTransaksi.UserID = userId
	newTransaksi.Produks = produks
	err = db.Create(newTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}

func TambahProdukKeTransaksi(db *gorm.DB, insertedTransaksi *Keranjang, produks *Produk) (err error) {
	insertedTransaksi.Produks = append(insertedTransaksi.Produks, produks)
	err = db.Save(insertedTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}
func TampilProdukDiTransaksi(db *gorm.DB, transaksi *Transaksi, id int) (err error) {
	err = db.Where("id=?", id).Preload("Products").Find(transaksi).Error
	if err != nil {
		return err
	}
	return nil
}

func TampilTransaksiId(db *gorm.DB, transaksis *[]Transaksi, id int) (err error) {
	err = db.Where("user_id=?", id).Find(transaksis).Error
	if err != nil {
		return err
	}
	return nil
}
