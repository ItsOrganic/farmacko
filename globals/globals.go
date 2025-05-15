package globals

import (
	"database/sql"

	"github.com/itsorganic/farmacko-assignment/models"
)

var (
	Config *models.Config
	DbConn *sql.DB
	Cache  *models.CouponCache
)
