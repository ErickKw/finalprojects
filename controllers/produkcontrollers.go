package controllers

import (
	"fmt"
	"net/http"
	"projectsrest/database"
	"projectsrest/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type ProdukControllers struct {
	// Declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitProdukControllers(s *session.Store) *ProdukControllers {
	db := database.ConnectDatabase()
	db.AutoMigrate(&models.Produk{})
	return &ProdukControllers{Db: db, store: s}
}

func (controllers *ProdukControllers) TampilProduk(c *fiber.Ctx) error {
	var prod []models.Produk
	err := models.TampilProduk(controllers.Db, &prod)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan1",
		})
	}
	sess, err := controllers.store.Get(c)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan2",
		})
	}
	val := sess.Get("userId")
	return c.JSON(fiber.Map{
		"message": "Data berhasil Tampil",
		"UserId":  val,
		"Prod":    prod,
	})
}

// func (controllers *ProdukControllers) TampilProdukId(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var products models.Produk
// 	if err := database.DB.First(&products, id).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Status(http.StatusNotFound).JSON(fiber.Map{
// 				"message": "Data tidak ditemukan1",
// 			})
// 		}
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Data tidak ditemuka!",
// 		})
// 	}

// 	return c.JSON(products)
// }

func (controllers *ProdukControllers) TambahProduk(c *fiber.Ctx) error {
	var newproduk models.Produk
	if err := c.BodyParser(&newproduk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "image" key:
		files := form.File["image"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			newproduk.Image = fmt.Sprintf(file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("upload/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	if err := database.DB.Create(&newproduk).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(newproduk)
}

func (controllers *ProdukControllers) DetailProduk(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var produk models.Produk
	err := models.TampilProdukId(controllers.Db, &produk, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(produk)

}

func (controllers *ProdukControllers) EditProduk(c *fiber.Ctx) error {
	var updateproducts models.Produk
	Idproducts := c.Params("id")
	//ambil content body yang dikirimkan
	if err := c.BodyParser(&updateproducts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if database.DB.Where("id=?", Idproducts).Updates(&updateproducts).RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal update data!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate!",
	})
}

func (controllers *ProdukControllers) HapusProduk(c *fiber.Ctx) error {

	var delproduk models.Produk
	Idproduk := c.Params("id")

	if database.DB.Delete(&delproduk, Idproduk).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Gagal menghapus data!",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus!",
	})
}
