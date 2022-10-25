package controllers

// import (
// 	"projectsrest/database"
// 	"projectsrest/models"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/session"
// 	"gorm.io/gorm"
// )

// type KeranjangControllers struct {
// 	// Declare variables
// 	Db    *gorm.DB
// 	store *session.Store
// }

// func InitKeranjangControllers(s *session.Store) *KeranjangControllers {
// 	db := database.ConnectDatabase()
// 	// gorm sync
// 	db.AutoMigrate(&models.Keranjang{})

// 	return &KeranjangControllers{Db: db, store: s}
// // }

// // GET /addtocart/:cartid/products/:productid
// func (controllers *KeranjangControllers) TambahKeKeranjang(c *fiber.Ctx) error {
// 	params := c.AllParams() // "{"id": "1"}"

// 	intKeranjangId, _ := strconv.Atoi(params["keranjangid"])
// 	intProdukId, _ := strconv.Atoi(params["produkid"])

// 	var keranjang models.Keranjang
// 	var produk models.Produk

// 	// Find the product first,
// 	err := models.TampilProdukId(controllers.Db, &keranjang, intProdukId)
// 	if err != nil {
// 		return c.SendStatus(500) // http 500 internal server error
// 	}

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

// // GET /shoppingcart/:cartid
// func (controllers *KeranjangControllers) AmbilKeranjang(c *fiber.Ctx) error {
// 	params := c.AllParams() // "{"id": "1"}"

// 	intCartId, _ := strconv.Atoi(params["cartid"])

// 	var cart models.Keranjang
// 	err := models.TampilProdukDiKeranjang(controllers.Db, &cart, intCartId)
// 	if err != nil {
// 		return c.SendStatus(500) // http 500 internal server error
// 	}

// 	sess, err := controllers.store.Get(c)
// 	if err != nil {
// 		panic(err)
// 	}
// 	val := sess.Get("userId")

// 	return c.JSON(fiber.Map{
// 		"message": "Keranjang",
// 		"UserId":  val,
// 	})
// }
