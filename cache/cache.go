package cache

import (
	"github.com/itsorganic/farmacko-assignment/globals"
	"github.com/itsorganic/farmacko-assignment/models"
)

func Init() {
	globals.Cache = &models.CouponCache{
		CouponCache: make(map[string]*models.Coupon),
	}
}
func SetCouponCode(couponCode string, coupon models.Coupon) {
	globals.Cache.MuRW.Lock()
	globals.Cache.CouponCache[couponCode] = &coupon
	globals.Cache.MuRW.Unlock()
}

func GetCouponCacheById(couponCode string) *models.Coupon {
	globals.Cache.MuRW.RLock()
	coupon, ok := globals.Cache.CouponCache[couponCode]
	defer globals.Cache.MuRW.RUnlock()
	if !ok {
		return nil
	}
	return coupon
}
