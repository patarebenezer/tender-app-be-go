package seed

import (
	"tender-app-be-go/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Ensure(db *gorm.DB) error {
	var userCount int64
	if err := db.Model(&models.User{}).Count(&userCount).Error; err != nil {
		return err
	}

	if userCount == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		admin := models.User{
			Username: "admin",
			Password: string(hash),
			RoleID:   1,
			VendorID: 0,
		}
		if err := db.Create(&admin).Error; err != nil {
			return err
		}
	}

	var vendorCount int64
	if err := db.Model(&models.Vendor{}).Count(&vendorCount).Error; err != nil {
		return err
	}

	if vendorCount == 0 {
		vendors := []models.Vendor{
			{
				Name:         "VENDOR 001",
				PICName:      "PIC 001",
				Email:        "vendor_001@test.com",
				Address:      "Jakarta",
				PhoneNumber:  "0812345689",
				PaymentTerms: "d30",
				DeliveryTime: 7,
			},
			{
				Name:         "VENDOR 002",
				PICName:      "PIC 002",
				Email:        "vendor_002@test.com",
				Address:      "Solo",
				PhoneNumber:  "08123456890",
				PaymentTerms: "d14",
				DeliveryTime: 14,
			},
			{
				Name:         "VENDOR 003",
				PICName:      "PIC 003",
				Email:        "vendor_003@test.com",
				Address:      "Surabaya",
				PhoneNumber:  "0813345649",
				PaymentTerms: "d7",
				DeliveryTime: 3,
			},
		}
		if err := db.Create(&vendors).Error; err != nil {
			return err
		}
	}
	return nil
}
