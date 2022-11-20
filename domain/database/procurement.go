package database

import "time"

type (
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
