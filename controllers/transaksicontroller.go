package controllers

import (
	"rapidtech/shoppingcart-res/database"
	"rapidtech/shoppingcart-res/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransactionAPIController struct {
	// Declare variables
	Db *gorm.DB
}

func InitTransactionAPIController() *TransactionAPIController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.Transaksi{})

	return &TransactionAPIController{Db: db}
}

// GET /checkout/:userid
func (controller *TransactionAPIController) PayTransaction(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intUserId, _ := strconv.Atoi(params["userid"])

	var transaction models.Transaksi
	var cart models.Cart

	// Find the product first,
	err := models.FindCart(controller.Db, &cart, intUserId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	errs := models.CreateTransaction(controller.Db, &transaction, int(intUserId), cart.Products)
	if errs != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// Delete products in cart
	errss := models.DeleteProductInChart(controller.Db, cart.Products, &cart, int(intUserId))
	if errss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.JSON(fiber.Map{
		"message": "Transaksi Sukses",
	})
}

// GET /historytransaksi/:userid
func (controller *TransactionAPIController) GetTransaction(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intUserId, _ := strconv.Atoi(params["userid"])

	var transactions []models.Transaksi
	err := models.ViewTransactionById(controller.Db, &transactions, intUserId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Title":      "Riwayat Pembelian",
		"Transaksis": transactions,
	})

}

// GET /history/detail/:transaksiid
func (controller *TransactionAPIController) DetailTransaction(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intTransactionId, _ := strconv.Atoi(params["transactionid"])

	var transaction models.Transaksi
	err := models.ViewTransaction(controller.Db, &transaction, intTransactionId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(fiber.Map{
		"Title":    "History Transaksi",
		"Products": transaction.Products,
	})
}
