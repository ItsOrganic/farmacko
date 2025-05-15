package constants

var (
	INSERT_COUPON = `INSERT INTO "coupon_codes" (
	coupon_code,
	expiry_date,
	usage_type, 
	applicable_medicine_ids, 
	applicable_categories,
	min_order_value,
	valid_time_window,
	terms_and_conditions, 
	discount_type,
	discount_value,
	max_usage_per_user) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	GET_ALL_COUPONS = `SELECT * FROM coupon_codes;`
)
var (
	STANDARD_TIME_FORMAT = "2006-01-02 15:04:05"
	PERCENTAGE           = "percentage"
	FLAT                 = "flat"
)

var (
	ERR_CONFIG_PATH_NOT_FOUND = "Config file path cannot be empty "
	ERR_CONNECTING_DB         = "Error connecting to database "
	ERR_PINGING_DB            = "Error pinging database "
	ERR_INSERTING_COUPON      = "Error inserting coupon "
	ERR_LOADING_COUPONS       = "Error loading coupons from database "
	ERR_SCANNING_COUPONS      = "Error scanning coupons from database "
	ERR_ITERATING_COUPONS     = "Error iterating over coupons "

	ERR_ORDER_TOTAL_LESS_THAN_MIN_VALUE = "Order total is less than minimum order value"
	ERR_COUPON_NOT_APPLICABLE           = "Coupon not applicable to cart items"
	ERR_UNAUTHORIZED                    = "Unauthorized access "
	ERR_INVALID_PAYLOAD                 = "Invalid payload"
	ERR_COUPON_NOT_VALID                = "Invalid coupon code"
	ERR_INVALID_EXP_DATE                = "Invalid expiry date"
)

var (
	SUCCESS_COUPON_ADDED   = "Coupon added successfully"
	SUCCESS_COUPON_CREATED = "Coupon created successfully"

	SUCCESS_CACHE_LOADED   = "Cache loaded successfully"
	SUCCESS_COUPON_APPLIED = "Coupon applied successfully"
)
