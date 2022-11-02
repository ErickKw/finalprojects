package controllers

import (
	"projectsrest/database"
	"projectsrest/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	// "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"time"
	// "github.com/gofiber/jwt/v3"
)

type AuthControllers struct {
	DB *gorm.DB
	// store *session.Store
}

func InitAuthControllers() *AuthControllers {
	db := database.ConnectDatabase()
	// db.AutoMigrate(&models.User{})
	return &AuthControllers{DB: db}
}

func (controllers *AuthControllers) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(400) // Bad Request, RegisterForm is not complete
	}

	// Hash password
	hashpw, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	sHash := string(hashpw)

	// Simpan hashing, bukan plain passwordnya
	user.Password = sHash

	// save user
	err := models.TambahUser(controllers.DB, &user)
	if err != nil {
		return c.SendStatus(500) // Server error, gagal menyimpan user
	}

	var keranjang models.Keranjang
	errCart := models.TambahKeranjang(controllers.DB, &keranjang, user.ID)
	if errCart != nil {
		// Server error, gagal menyimpan user
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "Server error, gagal menyimpan user",
		})
	}

	return c.JSON(user)
}

type LoginForm struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func (controllers *AuthControllers) Login(c *fiber.Ctx) error {
	var user models.User
	var form LoginForm

	if err := c.BodyParser(&form); err != nil {
		return c.JSON("Login Failed!")
	}
	var checker = validator.New()
	errors := checker.Struct(form)
	if errors != nil {
		return c.JSON("Login Failed!")
	}

	er := models.CariUsername(controllers.DB, &user, form.Username)
	if er != nil {
		return c.JSON("Username Not Valid!")
	}
	// user := c.FormValue(users.Username)
	// pass := c.FormValue(users.Password)
	// Throws Unauthorized error

	// if username != users.Username || password != users.Password {
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }

	// Create the Claims
	compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if compare == nil {
		exp := time.Now().Add(time.Hour * 72)
		claims := jwt.MapClaims{
			"name":  user.Username,
			"admin": true,
			"exp":   exp.Unix(),
		}
		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("mysecretpassword"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"message": "Success Login",
			"token":   t,
			"expired": exp.Format("2006-01-02 15:04:05"),
		})
	}
	return c.SendString("Status Unauthorized")

}

// func (controllers *AuthControllers) Login(c *fiber.Ctx) error {
// 	sess, err := controllers.store.Get(c)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var user models.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(http.StatusNotFound).JSON(fiber.Map{
// 			"message": "Data User Tidak Ditemukan 1",
// 		})
// 	}
// 	// Find user
// 	errs := models.CariUsername(controllers.Db, &user, user.Username)
// 	if errs != nil {
// 		return c.Status(http.StatusNotFound).JSON(fiber.Map{
// 			"message": "Data tidak ditemukan 2",
// 		})
// 	}

// 	// Compare password
// 	compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
// 	if compare == nil { // compare == nil artinya hasil compare di atas true
// 		sess.Set("username", user.Username)
// 		sess.Set("userId", user.ID)
// 		sess.Save()

// 		return c.SendString("Data Tidak Ditemukan!")
// 	}
// 	return c.SendString("Success")
// }

// /logout
func (controllers *AuthControllers) Logout(c *fiber.Ctx) error {

	// sess, err := controllers.store.Get(c)
	// if err != nil {
	// 	panic(err)
	// }
	// sess.Destroy()
	return c.JSON(fiber.Map{
		"message": "Berhasil Logout",
	})
}
