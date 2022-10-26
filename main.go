package main

import (
	"rapidtech/shoppingcart-res/controllers"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v3"
)

func main() {

	app := fiber.New()

	// controllers
	prodAPIController := controllers.InitProductAPIController()
	authAPIController := controllers.InitAuthAPIController()
	cartAPIController := controllers.InitCartAPIController()
	transAPIController := controllers.InitTransactionAPIController()

	prod := app.Group("/products")
	prod.Get("/", prodAPIController.IndexAPIProduct)
	prod.Post("/create", prodAPIController.AddAPIProduct)
	prod.Get("/detailproduct/:id", prodAPIController.GetAPIDetailProduct)
	prod.Put("/editproduct/:id", prodAPIController.EditAPIProduct)
	prod.Delete("/deleteproduct/:id", prodAPIController.DeleteAPIProduct)

	// Login/Registrasi
	app.Post("/login", authAPIController.LoginPosted)
	app.Post("/register", authAPIController.AddPostedRegister)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("test123456"),
	}))

	// Restricted Routes
	app.Get("/restricted", authAPIController.Restricted)

	// Add to Cart
	prod.Get("/addtocart/:cartid/products/:productid", cartAPIController.AddPostedCart)

	// Cart Detail
	cart := app.Group("/cart")
	cart.Get("/:userid", cartAPIController.GetDetailCart)

	// Pay
	trans := app.Group("/transaksi")
	trans.Get("/:userid", transAPIController.PayTransaction)

	app.Listen(":3000")
}
