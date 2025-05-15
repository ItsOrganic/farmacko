package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/itsorganic/farmacko-assignment/cache"
	"github.com/itsorganic/farmacko-assignment/constants"
	"github.com/itsorganic/farmacko-assignment/globals"
	"github.com/itsorganic/farmacko-assignment/models"
	"github.com/lib/pq"
)

var Db = globals.DbConn

func InitDbConn() {
	loadDbConn()
	LoadAllCoupons()
}

func loadDbConn() {
	config := globals.Config.Db
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Pass, config.Database, config.SslMode)
	db, err := sql.Open(config.Driver, connStr)
	if err != nil {
		log.Fatal(constants.ERR_CONNECTING_DB, err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(constants.ERR_PINGING_DB, err)
	}
	Db = db
}

type CouponRepository struct{}

func (cr *CouponRepository) InsertCoupon(coupon models.Coupon) error {
	_, err := Db.Exec(constants.INSERT_COUPON, coupon.CouponCode,
		coupon.ExpDate, coupon.UsageType, pq.StringArray(coupon.MedicineIds), pq.StringArray(coupon.CategoryIds),
		coupon.MinOrderValue, coupon.ValidTimeWindow, coupon.TermsAndConditions,
		coupon.DiscountType, coupon.DiscountValue, coupon.MaxUsagePerUser)
	if err != nil {
		log.Print(constants.ERR_INSERTING_COUPON, err)
		return err
	}
	return nil
}

func LoadAllCoupons() {
	rows, err := Db.Query(constants.GET_ALL_COUPONS)
	if err != nil {
		log.Print(constants.ERR_LOADING_COUPONS, err)
	}
	defer rows.Close()
	var coupons []models.Coupon
	for rows.Next() {
		var coupon models.Coupon
		err := rows.Scan(&coupon.CouponCode, &coupon.ExpDate, &coupon.UsageType, pq.Array(&coupon.MedicineIds),
			pq.Array(&coupon.CategoryIds), &coupon.MinOrderValue, &coupon.ValidTimeWindow,
			&coupon.TermsAndConditions, &coupon.DiscountType, &coupon.DiscountValue, &coupon.MaxUsagePerUser)
		if err != nil {
			log.Print(constants.ERR_SCANNING_COUPONS, err)
		}
		coupons = append(coupons, coupon)
	}
	if err := rows.Err(); err != nil {
		log.Print(constants.ERR_ITERATING_COUPONS, err)
	}
	for _, coupon := range coupons {
		cache.SetCouponCode(coupon.CouponCode, coupon)
	}
	log.Print(constants.SUCCESS_CACHE_LOADED)
}
