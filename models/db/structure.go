package db

import (
	"time"
)

type (
	User struct {
		ID          int       `gorm:"column:id;primary_key" json:"user_id"`
		UserName    string    `gorm:"column:user_name" json:"user_name"` // unique
		Name        string    `gorm:"column:name" json:"name"`
		Password    string    `gorm:"column:password" json:"password"`
		IDNumber    string    `gorm:"column:id_number" json:"id_number"`   // unique
		UserPhone   string    `gorm:"column:user_phone" json:"user_phone"` // unique
		UserAddress string    `gorm:"column:user_address" json:"user_address"`
		IsNew       int       `gorm:"column:is_new" json:"is_new"`
		CompanyID   int       `gorm:"column:company_id" json:"company_id"`
		UserRole    int       `gorm:"column:user_role" json:"user_role"`
		CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

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

	Client struct {
		ID            int       `gorm:"column:id;primary_key" json:"client_id"`
		ClientName    string    `gorm:"column:client_name" json:"client_name"` // unique
		ClientParent  string    `gorm:"column:client_parent" json:"client_parent"`
		ClientField   int       `gorm:"column:client_field" json:"client_field"`
		ClientAddress string    `gorm:"column:client_address" json:"client_address"`
		ClientPhone   string    `gorm:"column:client_phone" json:"client_phone"`     // unique
		ClientWebsite string    `gorm:"column:client_website" json:"vendor_website"` // unique
		Email         string    `gorm:"column:client_email" json:"client_email"`     // unique
		SocialMedia   string    `gorm:"column:social_media" json:"social_media"`
		Province      string    `gorm:"column:province" json:"province"`
		City          string    `gorm:"column:city" json:"city"`
		District      string    `gorm:"column:district" json:"district"`
		Avatar        string    `gorm:"column:avatar" json:"avatar"`
		IsNew         int       `gorm:"column:is_new" json:"is_new"`
		CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
		UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	}

	Procurement struct {
		ID            int       `gorm:"column:id;primary_key" json:"procurement_id"`
		VendorID      int       `gorm:"column:vendor_id" json:"vendor_id"`
		UserID        int       `gorm:"column:user_id" json:"user_id"`
		ClientID      int       `gorm:"column:client_id" json:"client_id"`
		Type          int       `gorm:"column:type" json:"type"`
		Category      int       `gorm:"column:category" json:"category"`
		Title         string    `gorm:"column:title" json:"title"`
		Body          string    `gorm:"column:body" json:"body"`
		PriceStart    int       `gorm:"column:price_start" json:"price_start"`
		PriceEnd      int       `gorm:"column:price_end" json:"price_end"`
		PaymentMethod int       `gorm:"column:payment_method" json:"payment_method"`
		TermStart     int       `gorm:"column:term_start" json:"term_start"`
		TermEnd       int       `gorm:"column:term_end" json:"term_end"`
		Image         string    `gorm:"column:image" json:"image"`
		EditedBySuper int       `gorm:"column:edited_by_super" json:"edited_by_super"`
		Status        int       `gorm:"column:status" json:"status"`
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
