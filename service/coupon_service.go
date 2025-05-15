package service

import (
	"errors"

	"github.com/itsorganic/farmacko-assignment/cache"
	"github.com/itsorganic/farmacko-assignment/constants"
	"github.com/itsorganic/farmacko-assignment/database"
	"github.com/itsorganic/farmacko-assignment/globals"
	"github.com/itsorganic/farmacko-assignment/models"
	"github.com/itsorganic/farmacko-assignment/utils"
)

type CouponService struct {
	Repo *database.CouponRepository
}

func NewCouponService(repo *database.CouponRepository) *CouponService {
	return &CouponService{Repo: repo}
}

func (cs *CouponService) CreateCoupon(coupon models.Coupon) error {
	if globals.Config.UserMode != "admin" {
		return errors.New(constants.ERR_UNAUTHORIZED)
	}

	if utils.VerifyExpTime(coupon.CouponCode, coupon.ExpDate) {
		return errors.New(constants.ERR_INVALID_EXP_DATE)
	}

	if err := cs.Repo.InsertCoupon(coupon); err != nil {
		return err
	}

	cache.SetCouponCode(coupon.CouponCode, coupon)
	return nil
}
