package main

import (
	"projectsrest/controllers"
	// "projectsrest/config"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	// store := session.New()
	app := fiber.New()

	authcontrollers := controllers.InitAuthControllers()
	prodcontrollers := controllers.InitProdukControllers()
	keranjangcontrollers := controllers.InitKeranjangControllers()
	transaksicontrollers := controllers.InitTransaksiControllers()

	app.Post("/register", authcontrollers.Register)
	app.Post("/login", authcontrollers.Login)
	app.Get("/logout", authcontrollers.Logout)

	//JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretpassword"),
	}))

	prod := app.Group("/produk")
	prod.Get("/", prodcontrollers.TampilProduk)
	prod.Post("/tambahproduk", prodcontrollers.TambahProduk)
	prod.Get("/detailproduk/:id", prodcontrollers.DetailProdukId)
	prod.Post("/editproduk/:id", prodcontrollers.EditProduk)
	prod.Get("/hapusproduk/:id", prodcontrollers.HapusProduk)
	app.Get("/logout", authcontrollers.Logout)

	t := app.Group("/checkout")
	t.Post("/userid", transaksicontrollers.TambahKeTransaksi)

	his := app.Group("/history")
	his.Get("/", transaksicontrollers.TampilTransaksi)
	his.Get("/:userid", transaksicontrollers.TampilTransaksi)
	his.Get("/detailtransaksi/:transaksiid", transaksicontrollers.DetailTransaksi)

	keranjang := app.Group("/keranjang")
	keranjang.Get("/:keranjangid", keranjangcontrollers.LihatKeranjang)

	keranjang.Get("/tambahkekeranjang/:keranjangid/produk/:produkid", keranjangcontrollers.TambahKeKeranjang)

	app.Listen(":3000")
}
