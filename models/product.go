package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	Id        int     `form:"id" json:"id" validate:"required" gorm:"PrimaryKey"`
	Name      string  `form:"name" json:"name" validate:"required"`
	Gambar    string  `form:"gambar" json:"gambar" validate:"required"`
	Quantity  int     `form:"quantity" json:"quantity" validate:"required"`
	Price     float32 `form:"price" json:"price" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CRUD
func CreateProduct(db *gorm.DB, newProduct *Products) (err error) {
	err = db.Create(newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProducts(db *gorm.DB, products *[]Products) (err error) {
	err = db.Find(products).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProductById(db *gorm.DB, product *Products, id int) (err error) {
	err = db.Where("id=?", id).First(product).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *gorm.DB, product *Products) (err error) {
	db.Save(product)

	return nil
}

func DeleteProductById(db *gorm.DB, product *Products, id int) (err error) {
	db.Where("id=?", id).Delete(product)

	return nil
}
