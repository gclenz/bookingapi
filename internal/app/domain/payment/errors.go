package payment

import "errors"

var (
	ErrPaymentNotFound = errors.New("payment not found")
	ErrCreatePayment   = errors.New(
		"something went wrong while trying to create a new payment. Please, try again. If the error persists, contact the development team",
	)
	ErrDuplicatedKey = errors.New(
		"there is already a payment with some of the provided data",
	)
	ErrMissingData = errors.New(
		"you should provide the following data to create a payment: booking id, external (gateway) id, gross amount, net amount",
	)
	ErrCouponNotFound = errors.New("coupon not found")
)
