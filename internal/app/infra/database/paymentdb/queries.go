package paymentdb

var (
	InsertPaymentQuery             = "insert into payments (id, booking_id, external_id, coupon_id, gross_amount, net_amount, status, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
	SelectPaymentByExternalIDQuery = "select id, booking_id, external_id, coupon_id, gross_amount, net_amount, status, created_at, updated_at from payments where external_id = $1;"
	SelectPaymentByIDQuery         = "select id, booking_id, external_id, coupon_id, gross_amount, net_amount, status, created_at, updated_at from payments where id = $1;"
	InsertCouponQuery              = "insert into coupons (id, code, mode, value, expires_at, created_at) values ($1, $2, $3, $4, $5, $6)"
	SelectCouponByIDQuery          = "select id, code, mode, value, expires_at, created_at from coupons where id = $1;"
)
