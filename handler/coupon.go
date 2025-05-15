package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsorganic/farmacko-assignment/cache"
	"github.com/itsorganic/farmacko-assignment/constants"
	"github.com/itsorganic/farmacko-assignment/globals"
	"github.com/itsorganic/farmacko-assignment/models"
	"github.com/itsorganic/farmacko-assignment/service"
	"github.com/itsorganic/farmacko-assignment/utils"
)

var CouponService *service.CouponService

// @Summary		Create a new coupon
// @Description	Create a new coupon with the provided details
// @Tags			Coupons
// @Accept			json
// @Produce		json
// @Param			coupon	body		models.Coupon		true	"Coupon details"
// @Success		200		{object}	map[string]string	"message: Coupon created successfully"
// @Failure		400		{object}	map[string]string	"error: Bad Request"
// @Failure		401		{object}	map[string]string	"error: Unauthorized"
// @Router			/coupons [post]
func CreateCoupon(c *gin.Context) {
	var coupon models.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := CouponService.CreateCoupon(coupon); err != nil {
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ERR_UNAUTHORIZED})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message": constants.SUCCESS_COUPON_CREATED})
}

// @Summary		Get Applicable Coupons
// @Description	Retrieve a list of applicable coupons for the given order
// @Tags			Coupons
// @Accept			json
// @Produce		json
// @Param			order	body		models.Order			true	"Order details"
// @Success		200		{object}	map[string]interface{}	"applicable_coupons: List of applicable coupons"
// @Failure		400		{object}	map[string]interface{}	"error: Bad Request"
// @Router			/coupons/applicable [post]
func GetApplicableCoupons(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ERR_INVALID_PAYLOAD})
		return
	}

	applicableCoupons := []models.Coupon{}
	for _, coupon := range globals.Cache.CouponCache {
		if coupon.ExpDate < time.Now().Format(constants.STANDARD_TIME_FORMAT) {
			continue
		}

		isApplicable := false
		for _, item := range order.CartItems {
			for _, medicineID := range coupon.MedicineIds {
				if item.Id == medicineID {
					isApplicable = true
					break
				}
			}
			for _, categoryID := range coupon.CategoryIds {
				if item.Category == categoryID {
					isApplicable = true
					break
				}
			}
		}

		if isApplicable && order.OrderTotal >= coupon.MinOrderValue {
			applicableCoupons = append(applicableCoupons, *coupon)
		}
	}

	c.JSON(http.StatusOK, gin.H{"applicable_coupons": applicableCoupons})
}

// @Summary		Validate Coupon
// @Description	Validate a coupon code for the given order
// @Tags			Coupons
// @Accept			json
// @Produce		json
// @Param			order	body		models.Order			true	"Order details"
// @Success		200		{object}	map[string]interface{}	"is_valid: true, discount: {items_discount: float64}, message: Coupon applied successfully"
// @Failure		400		{object}	map[string]interface{}	"is_valid: false, reason: string"
// @Router			/coupons/validate [post]
func ValidateCoupon(c *gin.Context) {
	var request models.Order

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ERR_INVALID_PAYLOAD})
		return
	}

	coupon := cache.GetCouponCacheById(request.CouponCode)
	if coupon == nil || coupon.CouponCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"is_valid": false, "reason": constants.ERR_COUPON_NOT_VALID})
		return
	}
	if utils.VerifyExpTime(coupon.CouponCode, coupon.ExpDate) {
		return
	}

	isApplicable := false
	for _, item := range request.CartItems {
		for _, medicineID := range coupon.MedicineIds {
			if item.Id == medicineID {
				isApplicable = true
				break
			}
		}
		for _, categoryID := range coupon.CategoryIds {
			if item.Category == categoryID {
				isApplicable = true
				break
			}
		}
	}

	if !isApplicable {
		c.JSON(http.StatusBadRequest, gin.H{"is_valid": false, "reason": constants.ERR_COUPON_NOT_APPLICABLE})
		return
	}

	if request.OrderTotal < coupon.MinOrderValue {
		c.JSON(http.StatusBadRequest, gin.H{"is_valid": false, "reason": constants.ERR_ORDER_TOTAL_LESS_THAN_MIN_VALUE})
		return
	}

	itemsDiscount := coupon.DiscountValue
	if coupon.DiscountType == constants.PERCENTAGE {
		itemsDiscount = request.OrderTotal * (coupon.DiscountValue / 100)
	} else if coupon.DiscountType == constants.FLAT {
		itemsDiscount = coupon.DiscountValue
	}

	c.JSON(http.StatusOK, gin.H{
		"is_valid": true,
		"discount": gin.H{
			"items_discount": itemsDiscount,
		},
		"message": constants.SUCCESS_COUPON_APPLIED,
	})
}
