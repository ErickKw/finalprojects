package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `form:"name" json:"name" validate:"required"`
	Address   string `form:"address" json:"address" validate:"required"`
	Nohp      int    `form:"nohp" json:"nohp" validate:"required"`
	Email     string `form:"email" json:"email" validate:"required"`
	Username  string `form:"username" json:"username" validate:"required"`
	Password  string `form:"password" json:"password" validate:"required"`
	Keranjang Keranjang
	Transaksi []Transaksi
}

func TambahUser(db *gorm.DB, newUser *User) (err error) {
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func CariUsername(db *gorm.DB, user *User, username string) (err error) {
	err = db.Where("username=?", username).First(user).Error
	if err != nil {
		return err
	}
	return nil
}
