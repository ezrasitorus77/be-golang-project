package db

import (
	"time"
)

type (
	User struct {
		ID          int       `gorm:"column:id;primary_key" json:"user_id"`
		UserName    string    `gorm:"column:user_name" json:"user_name"`
		Name        string    `gorm:"column:name" json:"name"`
		Password    string    `gorm:"column:password" json:"password"`
		IDNumber    string    `gorm:"column:id_number" json:"id_number"`
		UserPhone   string    `gorm:"column:user_phone" json:"user_phone"`
		UserAddress string    `gorm:"column:user_address" json:"user_address"`
		IsNew       int       `gorm:"column:is_new" json:"is_new"`
		CompanyID   int       `gorm:"column:company_id" json:"company_id"`
		UserRole    int       `gorm:"column:user_role" json:"user_role"`
		CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	Vendor struct {
		ID            int       `gorm:"column:id;primary_key" json:"vendor_id"`
		VendorName    string    `gorm:"column:vendor_name" json:"vendor_name"`
		VendorField   int       `gorm:"column:vendor_field" json:"vendor_field"`
		VendorType    int       `gorm:"column:vendor_type" json:"vendor_type"`
		VendorAddress string    `gorm:"column:vendor_address" json:"vendor_address"`
		VendorPhone   string    `gorm:"column:vendor_phone" json:"vendor_phone"`
		VendorWebsite string    `gorm:"column:vendor_website" json:"vendor_website"`
		Email         string    `gorm:"column:vendor_email" json:"vendor_email"`
		NPWP          string    `gorm:"column:npwp" json:"npwp"`
		SocialMedia   string    `gorm:"column:social_media" json:"social_media"`
		Province      string    `gorm:"column:province" json:"province"`
		City          string    `gorm:"column:city" json:"city"`
		District      string    `gorm:"column:district" json:"district"`
		Avatar        string    `gorm:"column:avatar" json:"avatar"`
		IsNew         int       `gorm:"column:is_new" json:"is_new"`
		CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	Client struct {
		ID            int       `gorm:"column:id;primary_key" json:"client_id"`
		ClientName    string    `gorm:"column:client_name" json:"client_name"`
		ClientParent  string    `gorm:"column:client_parent" json:"client_parent"`
		ClientField   int       `gorm:"column:client_field" json:"client_field"`
		ClientAddress string    `gorm:"column:client_address" json:"client_address"`
		ClientPhone   string    `gorm:"column:client_phone" json:"client_phone"`
		ClientWebsite string    `gorm:"column:client_website" json:"vendor_website"`
		Email         string    `gorm:"column:client_email" json:"client_email"`
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

func (tbl *User) TableName() string {
	return "user"
}

func (tbl *Vendor) TableName() string {
	return "vendor"
}

func (tbl *Client) TableName() string {
	return "client"
}
