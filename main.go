package main

import (
	"github.com/itsorganic/farmacko-assignment/apploader"
	"github.com/itsorganic/farmacko-assignment/controller"
	"github.com/itsorganic/farmacko-assignment/database"
	"github.com/itsorganic/farmacko-assignment/service"
)

func main() {
	apploader.Init()
	database.InitDbConn()
	couponRepo := &database.CouponRepository{}
	couponService := service.NewCouponService(couponRepo)
	controller.InitServer(couponService)
}
