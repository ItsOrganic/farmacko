package models

import (
	"sync"
)

type Config struct {
	Db       DBConfig `yaml:"db"`
	UserMode string   `yaml:"user-mode"`
}

type DBConfig struct {
	User     string `yaml:"user-name"`
	Pass     string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
	SslMode  string `yaml:"ssl-mode"`
}

type Coupon struct {
	CouponCode         string   `json:"coupon_code" sql:"PRIMARY KEY"`
	ExpDate            string   `json:"expiry_date"` // "YYYY-MM-DD HH:MM:SS"
	UsageType          string   `json:"usage_type"`  // "one_time" , "multi_use" , or "time_based"
	MedicineIds        []string `json:"applicable_medicine_ids"`
	CategoryIds        []string `json:"applicable_categories"`
	MinOrderValue      float64  `json:"min_order_value"`
	ValidTimeWindow    string   `json:"valid_time_window"`
	TermsAndConditions string   `json:"terms_and_conditions"`
	DiscountType       string   `json:"discount_type"` // "percentage" or "flat"
	DiscountValue      float64  `json:"discount_value"`
	MaxUsagePerUser    int      `json:"max_usage_per_user"`
}

type CouponCache struct {
	MuRW        sync.RWMutex
	CouponCache map[string]*Coupon `json:"coupon_cache"`
}

type Product struct {
	Id       string
	Category string
}

type Order struct {
	CouponCode string    `json:"coupon_code"`
	CartItems  []Product `json:"cart_items"`
	OrderTotal float64   `json:"order_total"`
	Timestamp  string    `json:"timestamp"`
}
