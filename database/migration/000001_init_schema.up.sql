CREATE TABLE IF NOT EXISTS coupon_codes (
		coupon_code TEXT PRIMARY KEY NOT NULL,
		expiry_date timestamp NOT NULL,
		usage_type TEXT NOT NULL,
		applicable_medicine_ids TEXT[] NOT NULL,
		applicable_categories TEXT[] NOT NULL,
		min_order_value REAL NOT NULL,
		valid_time_window TEXT NOT NULL,
		terms_and_conditions TEXT NOT NULL,
		discount_type TEXT NOT NULL,
		discount_value REAL NOT NULL,
		max_usage_per_user INTEGER NOT NULL
	);