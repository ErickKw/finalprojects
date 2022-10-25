package models

import "gorm.io/gorm"

type Keranjang struct {
	gorm.Model
	UserID  uint
	Produks []*Produk `gorm:"many2many:keranjang_produk;"`
}

func TambahKeranjang(db *gorm.DB, newKeranjang *Keranjang, userId uint) (err error) {
	newKeranjang.UserID = userId
	err = db.Create(newKeranjang).Error
	if err != nil {
		return err
	}
	return nil
}

func TambahProdukKeKeranjang(db *gorm.DB, tambahProdukKeKeranjang *Keranjang, produk *Produk) (err error) {
	tambahProdukKeKeranjang.Produks = append(tambahProdukKeKeranjang.Produks, produk)
	err = db.Save(tambahProdukKeKeranjang).Error
	if err != nil {
		return err
	}
	return nil
}

func TampilProduksDiKeranjang(db *gorm.DB, keranjang *Keranjang, id int) (err error) {
	err = db.Where("user_id=?", id).Preload("Produks").Find(keranjang).Error
	if err != nil {
		return err
	}
	return nil
}

func TampilKeranjangId(db *gorm.DB, keranjang *Keranjang, id int) (err error) {
	err = db.Where("user_id=?", id).First(keranjang).Error
	if err != nil {
		return err
	}
	return nil
}

func EditKeranjang(db *gorm.DB, produk []*Produk, editKeranjang *Keranjang, userId uint) (err error) {
	db.Model(&editKeranjang).Association("Produks").Delete(produk)

	return nil
}
