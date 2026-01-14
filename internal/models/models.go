package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string //bcrypt hash
	RoleID   int
	VendorID uint
}

type Vendor struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	PICName      string
	Email        string
	Address      string
	PhoneNumber  string
	PaymentTerms string
	DeliveryTime int
}

type Tender struct {
	ID               uint `gorm:"primaryKey"`
	Name             string
	Date             string
	RequesterName    string
	Description      string
	TotalParticipant string
	TotalProduct     string
	Status           int
	CreatedBy        string
	CreatedOn        time.Time
	ModifiedBy       string
	ModifiedOn       time.Time
}

type TenderVendor struct {
	ID       uint `gorm:"primaryKey"`
	TenderID uint
	VendorID uint
}

type TenderProduct struct {
	ID            uint `gorm:"primaryKey"`
	TenderID      uint
	ProductName   string
	Brand         string
	Specification string
	UOM           string
	Quantity      float64
	TermOfPayment string
	LastPrice     float64
}
