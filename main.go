package main

import (
	"projectsrest/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	store := session.New()
	app := fiber.New()

	// Middleware to check login
	CheckLogin := func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		val := sess.Get("username")
		if val != nil {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	authcontrollers := controllers.InitAuthControllers(store)
	prodcontrollers := controllers.InitProdukControllers(store)

	prod := app.Group("/produk")
	prod.Get("/", CheckLogin, prodcontrollers.TampilProduk)
	prod.Get("/detail/:id", CheckLogin, prodcontrollers.TampilProdukId)
	prod.Post("/tambahproduk", CheckLogin, prodcontrollers.TambahProduk)
	prod.Post("/editproduk", CheckLogin, prodcontrollers.EditProduk)
	prod.Get("/hapusproduk", CheckLogin, prodcontrollers.HapusProduk)

	// auth := app.Group("/")
	app.Post("/register", authcontrollers.Register)
	app.Post("/login", authcontrollers.Login)
	app.Get("/logout", authcontrollers.Logout)

	app.Listen(":3000")
}
