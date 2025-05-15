package utils

import (
	"time"

	"github.com/itsorganic/farmacko-assignment/constants"
	"github.com/itsorganic/farmacko-assignment/globals"
)

func VerifyExpTime(couponCode, t string) bool {
	start := time.Now().Format(constants.STANDARD_TIME_FORMAT)
	if t > start {
		return false
	}
	if couponData, exists := globals.Cache.CouponCache[couponCode]; exists && couponData != nil {
		delete(globals.Cache.CouponCache, t)
	}
	return true
}
