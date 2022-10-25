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

	// authcontrollers := controllers.InitAuthControllers(store)

	// auth := app.Group("/")
	// app.Post("/register", authcontrollers.Register)
	// app.Post("/login", authcontrollers.Login)
	// app.Get("/logout", authcontrollers.Logout)
	// app.Get("/produk", restricted)
	// prodcontrollers := controllers.InitProdukControllers(store)
	// prod := app.Group("/produk")
	// prod.Get("/", CheckLogin, prodcontrollers.TampilProduk)
	// prod.Get("/detail/:id", CheckLogin, prodcontrollers.TampilProdukId)
	// prod.Post("/tambahproduk", prodcontrollers.TambahProduk)
	// prod.Post("/editproduk", prodcontrollers.EditProduk)
	// prod.Get("/hapusproduk", prodcontrollers.HapusProduk)

	authcontrollers := controllers.InitAuthControllers(store)
	app.Post("/login", authcontrollers.Login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretpassword"),
	}))
	prodcontrollers := controllers.InitProdukControllers(store)
	prod := app.Group("/produk")
	prod.Get("/", prodcontrollers.TampilProduk)
	prod.Get("/:id", prodcontrollers.TampilProdukId)
	app.Get("/logout", authcontrollers.Logout)

	app.Listen(":3000")
}

// func restricted(c *fiber.Ctx) error {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	name := claims["name"].(string)
// 	return c.SendString("Welcome " + name)
// }
