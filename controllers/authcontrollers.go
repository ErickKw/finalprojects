package controllers

import (
	"net/http"
	"projectsrest/database"
	"projectsrest/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthControllers struct {
	Db    *gorm.DB
	store *session.Store
}

func InitAuthControllers(s *session.Store) *AuthControllers {
	db := database.ConnectDatabase()
	db.AutoMigrate(&models.Produk{})
	return &AuthControllers{Db: db, store: s}
}

func (controllers *AuthControllers) Register(c *fiber.Ctx) error {
	var user models.User
	var keranjang models.Keranjang

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(400) // Bad Request, RegisterForm is not complete
	}

	// Hash password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	sHash := string(bytes)

	// Simpan hashing, bukan plain passwordnya
	user.Password = sHash

	// save user
	err := models.TambahUser(controllers.Db, &user)
	if err != nil {
		return c.SendStatus(500) // Server error, gagal menyimpan user
	}

	// Find user
	errs := models.CariUsername(controllers.Db, &user, user.Username)
	if errs != nil {
		return c.SendStatus(500) // Server error, gagal menyimpan user
	}

	// also create cart
	errKeranjang := models.TambahKeranjang(controllers.Db, &keranjang, user.ID)
	if errKeranjang != nil {
		return c.SendStatus(500) // Server error, gagal menyimpan user
	}
	// if succeed
	return c.JSON(user)
}

func (controllers *AuthControllers) Login(c *fiber.Ctx) error {
	sess, err := controllers.store.Get(c)
	if err != nil {
		panic(err)
	}
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data User Tidak Ditemukan 1",
		})
	}
	// Find user
	errs := models.CariUsername(controllers.Db, &user, user.Username)
	if errs != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan 2",
		})
	}

	// Compare password
	compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
	if compare == nil { // compare == nil artinya hasil compare di atas true
		sess.Set("username", user.Username)
		sess.Set("userId", user.ID)
		sess.Save()

		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data User Tidak Ditemukan 3",
		})
	}
	return c.SendString("Success")
}

// /logout
func (controllers *AuthControllers) Logout(c *fiber.Ctx) error {

	sess, err := controllers.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Destroy()
	return c.JSON(fiber.Map{
		"message": "Berhasil Logout",
	})
}
