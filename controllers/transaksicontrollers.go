package controllers

import (
	"projectsrest/database"
	"projectsrest/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransaksiController struct {
	// Declare variables
	Db *gorm.DB
}

func InitTransaksiControllers() *TransaksiController {
	db := database.ConnectDatabase()
	// gorm sync
	db.AutoMigrate(&models.Transaksi{})

	return &TransaksiController{Db: db}
}

func (controllers *TransaksiController) TambahKeTransaksi(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	UserId, _ := strconv.Atoi(params["userid"])

	var transaksi models.Transaksi
	var keranjang models.Keranjang

	// Find the cart
	errNoCart := models.TampilKeranjangId(controllers.Db, &keranjang, UserId)
	if errNoCart != nil {
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Tidak dapat menemukan Keranjang dengan Id " + params["userid"] + ", Gagal Melakukan Checkout",
		})
	}

	// Find the product first,
	err := models.TampilProduksDiKeranjang(controllers.Db, &keranjang, UserId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// Jika Cart kosong
	if len(keranjang.Produks) == 0 {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "Keranjang kosong, silahkan isi Produk ke Ke Keranjanag",
		})
	}

	errs := models.BuatTransaksi(controllers.Db, &transaksi, uint(UserId), keranjang.Produks)
	if errs != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// Delete products in cart
	errss := models.EditKeranjang(controllers.Db, keranjang.Produks, &keranjang, uint(UserId))
	if errss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// if succeed
	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil Checkout",
	})
}

func (controllers *TransaksiController) TampilTransaksi(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	UserId, _ := strconv.Atoi(params["userid"])

	var transaksis []models.Transaksi
	err := models.TampilTransaksiId(controllers.Db, &transaksis, UserId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Message":    "History Transaksi",
		"Transaksis": transaksis,
	})

}

func (controllers *TransaksiController) DetailTransaksi(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	TransaksiId, _ := strconv.Atoi(params["transaksiid"])

	var transaksi models.Transaksi
	err := models.TampilProdukDiTransaksi(controllers.Db, &transaksi, TransaksiId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Message":  "Detail Product pada Transaksi",
		"Products": transaksi.Produks,
	})
}
