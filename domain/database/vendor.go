package database

import "time"

type (
	Vendor struct {
		ID            int       `gorm:"column:id;primary_key" json:"vendor_id"`
		VendorName    string    `gorm:"column:vendor_name" json:"vendor_name"` // unique
		VendorField   int       `gorm:"column:vendor_field" json:"vendor_field"`
		VendorType    int       `gorm:"column:vendor_type" json:"vendor_type"`
		VendorAddress string    `gorm:"column:vendor_address" json:"vendor_address"`
		VendorPhone   string    `gorm:"column:vendor_phone" json:"vendor_phone"`     // unique
		VendorWebsite string    `gorm:"column:vendor_website" json:"vendor_website"` // unique
		Email         string    `gorm:"column:vendor_email" json:"vendor_email"`     // unique
		NPWP          string    `gorm:"column:npwp" json:"npwp"`                     // unique
		AccountName   string    `gorm:"column:account_name" json:"account_name"`     // unique
		AccountNumber string    `gorm:"column:account_number" json:"account_number"` // unique
		SocialMedia   string    `gorm:"column:social_media" json:"social_media"`
		Province      string    `gorm:"column:province" json:"province"`
		City          string    `gorm:"column:city" json:"city"`
		District      string    `gorm:"column:district" json:"district"`
		Avatar        string    `gorm:"column:avatar" json:"avatar"`
		IsNew         int       `gorm:"column:is_new" json:"is_new"`
		CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	}
)

func (tbl *Vendor) TableName() string {
	return "vendor"
}
