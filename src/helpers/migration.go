package helpers

import (
	"server/src/configs"
	"server/src/models"
)

func Migration() {
	configs.DB.AutoMigrate(
		&models.User{},
		&models.UserVerification{},
		&models.Address{},
		&models.Store{},
		&models.Cart{},
		&models.CartDetail{},
		&models.Category{},
		&models.Product{},
		&models.ProductImage{},
		&models.ProductColor{},
		&models.ProductSize{},
		&models.Transaction{},
		&models.TransactionDetail{},
	)
}
