package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/itsorganic/farmacko-assignment/globals"
)

// @Summary		Get all cache
// @Description	Retrieve all cached coupon data
// @Tags			Cache
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"cache data"
// @Failure		500	{object}	map[string]string		"error message"
// @Router			/coupon/cache [get]
func FetchCompleteCache(c *gin.Context) {
	data := globals.Cache.CouponCache
	if data == nil {
		c.JSON(500, gin.H{"error": "Cache is empty"})
		return
	}
	c.JSON(200, gin.H{"cache": data})
}
