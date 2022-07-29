package db

import (
	"time"
)

type (
	Vendor struct {
		ID             int       `gorm:"column:id;primary_key" json:"user_id"`
		UserName       string    `gorm:"column:user_name" json:"user_name"`
		Password       string    `gorm:"column:password" json:"password"`
		CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
		Email          string    `gorm:"column:email" json:"vendor_email"`
		IDNumber       string    `gorm:"column:id_number" json:"id_number"`
		NPWP           string    `gorm:"column:npwp" json:"npwp"`
		CompanyName    string    `gorm:"column:company_name" json:"company_name"`
		UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
		CompanyField   string    `gorm:"column:company_field" json:"company_field"`
		CompanyType    string    `gorm:"column:company_type" json:"company_type"`
		CompanyAddress string    `gorm:"column:company_address" json:"company_address"`
		CompanyPhone   string    `gorm:"column:company_phone" json:"company_phone"`
		CompanyWebsite string    `gorm:"column:company_website" json:"company_website"`
		SocialMedia    string    `gorm:"column:social_media" json:"social_media"`
		Province       string    `gorm:"column:province" json:"province"`
		City           string    `gorm:"column:city" json:"city"`
		District       string    `gorm:"column:district" json:"district"`
		Avatar         string    `gorm:"column:avatar" json:"avatar"`
	}

	Client struct {
		ID        int       `gorm:"column:id;primary_key" json:"client_id"`
		UserName  string    `gorm:"column:user_name" json:"user_name"`
		Password  string    `gorm:"column:password" json:"password"`
		CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
		Email     string    `gorm:"column:email" json:"client_email"`
	}
)

func (tbl *Vendor) TableName() string {
	return "vendor"
}

func (tbl *Client) TableName() string {
	return "client"
}
