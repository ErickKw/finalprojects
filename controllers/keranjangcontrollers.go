package controllers

import (
	"projectsrest/database"
	"projectsrest/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type KeranjangControllers struct {
	// Declare variables
	Db *gorm.DB
}

func InitKeranjangControllers() *KeranjangControllers {
	db := database.ConnectDatabase()
	// gorm sync
	db.AutoMigrate(&models.Keranjang{})

	return &KeranjangControllers{Db: db}
}

// GET /addtocart/:cartid/products/:productid
func (controllers *KeranjangControllers) TambahKeKeranjang(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intKeranjangId, _ := strconv.Atoi(params["keranjangid"])
	intProdukId, _ := strconv.Atoi(params["produkid"])

	var keranjang models.Keranjang
	var produk models.Produk

	// Find the product first,
	err := models.TampilProdukId(controllers.Db, &produk, intProdukId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	ertmplkrj := models.TampilKeranjangId(controllers.Db, &keranjang, intKeranjangId)
	if ertmplkrj != nil {
		return c.JSON(fiber.Map{
			"message": "Tidak Dapat Memuat Keranjang Status Server 500",
		})
	}

	ertmbhpro := models.TambahProdukKeKeranjang(controllers.Db, &keranjang, &produk)
	if ertmbhpro != nil {
		return c.JSON(fiber.Map{
			"message": "Gagal Menambahkan Ke Keranjang",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil Menambahkan Produk Ke Keranjang",
		"data":    params["produkid"],
	})
}

// GET /shoppingcart/:cartid
func (controllers *KeranjangControllers) LihatKeranjang(c *fiber.Ctx) error {
	params := c.AllParams()

	KeranjangId, _ := strconv.Atoi(params["keranjangid"])

	var keranjang models.Keranjang
	err := models.TampilKeranjangId(controllers.Db, &keranjang, KeranjangId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	errss := models.TampilProduksDiKeranjang(controllers.Db, &keranjang, KeranjangId)
	if errss != nil {
		return c.JSON(fiber.Map{
			"message": "Gagal Memuat Keranjang",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Keranjang" + params["cartid"],
		// "UserId":  val,
		"Produk": keranjang.Produks,
	})
}

// 	// Then find the cart
// 	errs := models.TampilKeranjangId(controllers.Db, &keranjang, intKeranajangId)
// 	if errs != nil {
// 		return c.SendStatus(500) // http 500 internal server error
// 	}

// 	// Finally, insert the product to cart
// 	errss := models.TambahKeranjang(controllers.Db, &keranjang, &produk)
// 	if errss != nil {
// 		return c.SendStatus(500) // http 500 internal server error
// 	}

// 	return c.Status(fiber.StatusOK).JSON(keranjang)
// }
