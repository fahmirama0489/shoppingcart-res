package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id        int          `form:"id" json:"id" validate:"required" gorm:"PrimaryKey"`
	Name      string       `form:"name" json:"name" validate:"required"`
	Gambar    string       `form:"gambar" json:"gambar" validate:"required"`
	Quantity  int          `form:"quantity" json:"quantity" validate:"required"`
	Price     float32      `form:"price" json:"price" validate:"required"`
	Carts     []*Cart      `gorm:"many2many:cart_products;"`
	Transaksi []*Transaksi `gorm:"many2many:transaksi_products;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CRUD
func CreateProduct(db *gorm.DB, newProduct *Product) (err error) {
	err = db.Create(newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProducts(db *gorm.DB, product *[]Product) (err error) {
	err = db.Find(product).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProductById(db *gorm.DB, product *Product, id int) (err error) {
	err = db.Where("id=?", id).First(product).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *gorm.DB, product *Product) (err error) {
	db.Save(product)

	return nil
}

func DeleteProductById(db *gorm.DB, product *Product, id int) (err error) {
	db.Where("id=?", id).Delete(product)

	return nil
}
