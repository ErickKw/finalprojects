package models

import "gorm.io/gorm"

type Keranjang struct {
	gorm.Model
	UserID uint
	Produk []*Produk `gorm:"many2many:keranjang_produk;"`
}

func TambahKeranjang(db *gorm.DB, newKeranjang *Keranjang, userId uint) (err error) {
	newKeranjang.UserID = userId
	err = db.Create(newKeranjang).Error
	if err != nil {
		return err
	}
	return nil
}

func TambahProdukKeKeranjang(db *gorm.DB, tambahProdukKeranjang *Keranjang, produk *Produk) (err error) {
	tambahProdukKeranjang.Produk = append(tambahProdukKeranjang.Produk, produk)
	err = db.Save(tambahProdukKeranjang).Error
	if err != nil {
		return err
	}
	return nil
}

func TampilProdukDiKeranjang(db *gorm.DB, keranjang *Keranjang, id int) (err error) {
	err = db.Where("user_id=?", id).Preload("Produk").Find(keranjang).Error
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

func UpdateCart(db *gorm.DB, produk []*Produk, editKeranjang *Keranjang, userId uint) (err error) {
	db.Model(&editKeranjang).Association("Products").Delete(produk)

	return nil
}
