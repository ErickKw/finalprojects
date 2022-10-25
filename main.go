package main

import (
	"projectsrest/controllers"
	// "projectsrest/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	store := session.New()
	app := fiber.New()

	authcontrollers := controllers.InitAuthControllers(store)
	app.Post("/register", authcontrollers.Register)
	app.Post("/login", authcontrollers.Login)
	app.Get("/logout", authcontrollers.Logout)

	//JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretpassword"),
	}))

	prodcontrollers := controllers.InitProdukControllers(store)
	prod := app.Group("/produk")
	prod.Get("/", prodcontrollers.TampilProduk)
	// prod.Get("/:id", prodcontrollers.TampilProdukId)
	prod.Post("/tambahproduk", prodcontrollers.TambahProduk)
	app.Get("/logout", authcontrollers.Logout)
	prod.Get("/detailproduk/:id", prodcontrollers.DetailProduk)
	prod.Get("/hapusproduk/:id", prodcontrollers.HapusProduk)

	app.Listen(":3000")
}
