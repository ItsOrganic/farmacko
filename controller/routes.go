package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/itsorganic/farmacko-assignment/docs"
	"github.com/itsorganic/farmacko-assignment/handler"
	"github.com/itsorganic/farmacko-assignment/service"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(couponService *service.CouponService) {
	router := gin.Default()

	handler.CouponService = couponService

	router.POST("/coupons", handler.CreateCoupon)
	router.GET("/coupon/cache", handler.FetchCompleteCache)
	router.GET("/coupons/applicable", handler.GetApplicableCoupons)
	router.POST("/coupons/validate", handler.ValidateCoupon)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	router.Run(":8080")
}
