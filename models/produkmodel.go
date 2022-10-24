package models

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	Id        int          `json: "id" validate:"required"`
	Name      string       `json: "name" validate:"required"`
	Image     string       `json: "image" validate:"required"`
	Quantity  int          `json: "quantity" validate:"required"`
	Price     float32      `json: "price" validate:"required"`
	Keranjang []*Keranjang `gorm:"many2many:keranjang_produk;"`
	Transaksi []*Transaksi `gorm:"many2many:transaksi_produk;"`
}

func TambahProduk(db *gorm.DB, newProduct *Produk) (err error) {
	err = db.Create(newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func TampilProduk(db *gorm.DB, produk *[]Produk) (err error) {
	err = db.Find(produk).Error
	if err != nil {
		return err
	}
	return nil
}

func TampilProdukId(db *gorm.DB, produk *Produk, id int) (err error) {
	err = db.Where("id=?", id).First(produk).Error
	if err != nil {
		return err
	}
	return nil
}

func EditProduk(db *gorm.DB, produk *Produk) (err error) {
	db.Save(produk)

	return nil
}

func HapusProdukId(db *gorm.DB, produk *Produk, id int) (err error) {
	db.Where("id=?", id).Delete(produk)

	return nil
}
