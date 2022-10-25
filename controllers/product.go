package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"rapidtech/shoppingcart-res/database"
	"rapidtech/shoppingcart-res/models"
)

// type ProductForm struct {
// 	Email string `form:"email" validate:"required"`
// 	Address string `form:"address" validate:"required"`
// }

type ProductAPIController struct {
	// declare variables
	Db *gorm.DB
}

func InitProductAPIController() *ProductAPIController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Products{})

	return &ProductAPIController{Db: db}
}

// routing
// GET /products
func (controller *ProductAPIController) IndexAPIProduct(c *fiber.Ctx) error {
	// load all products
	var products []models.Products
	err := models.ReadProducts(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(products)
}

// POST /products/create
func (controller *ProductAPIController) AddAPIProduct(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var myform models.Products

	file, errFile := c.FormFile("gambar")
	if errFile != nil {
		fmt.Println("Error File =", errFile)
	}
	var filename string = file.Filename
	if file != nil {

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/upload/%s", filename))
		if errSaveFile != nil {
			fmt.Println("Gambar berhasil di upload")
		}
	} else {
		fmt.Println("Gambar tidak di upload")
	}

	if err := c.BodyParser(&myform); err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	myform.Gambar = filename
	// save product
	errr := models.CreateProduct(controller.Db, &myform)
	if errr != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	// if succeed
	return c.JSON(myform)
}

// GET /products/detail/xxx
func (controller *ProductAPIController) GetAPIDetailProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Products
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(product)
}

// / PUT products/editproduct/xx
func (controller *ProductAPIController) EditAPIProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Products
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var myform models.Products

	if err := c.BodyParser(&myform); err != nil {
		return c.SendStatus(400)
	}

	file, errFile := c.FormFile("gambar")
	if errFile != nil {
		fmt.Println("Error File =", errFile)
	}
	var filename string = file.Filename
	if file != nil {

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/upload/%s", filename))
		if errSaveFile != nil {
			fmt.Println("Gambar tidak berhasil di upload")
		}
	} else {
		fmt.Println("Gambar tidak di upload")
	}
	myform.Gambar = filename
	product.Name = myform.Name
	product.Gambar = myform.Gambar
	product.Quantity = myform.Quantity
	product.Price = myform.Price
	// save product
	models.UpdateProduct(controller.Db, &product)

	return c.JSON(product)

}

// / GET /products/deleteproduct/xx
func (controller *ProductAPIController) DeleteAPIProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var product models.Products
	models.DeleteProductById(controller.Db, &product, idn)
	return c.JSON(fiber.Map{
		"message": "Produk di Hapus",
	})
}
