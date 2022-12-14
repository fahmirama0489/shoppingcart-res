package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int    `form:"id" json:"id" validate:"required" gorm:"PrimaryKey"`
	Name      string `form:"name" json:"name" validate:"required"`
	Email     string `form:"email" json:"email" validate:"required"`
	Username  string `form:"username" json:"username" validate:"required"`
	Password  string `form:"password" json:"password" validate:"required"`
	Cart      []*Cart
	Transaksi []Transaksi
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CRUD
func CreateUser(db *gorm.DB, newUser *User) (err error) {
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func FindByUsername(db *gorm.DB, user *User, username string) (err error) {
	err = db.Where("username=?", username).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadUser(db *gorm.DB, users *[]User) (err error) {
	err = db.Find(users).Error
	if err != nil {
		return err
	}
	return nil
}
