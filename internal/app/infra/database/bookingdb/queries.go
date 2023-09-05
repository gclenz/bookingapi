package bookingdb

var (
	InsertBookingQuery              = "insert into bookings (id, room_id, customer_id, start, end, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7);"
	SelectBookingsByRoomIDQuery     = "select id, customer_id, room_id, start, end, created_at, updated_at from bookings where room_id = $1;"
	SelectBookingsByCustomerIDQuery = "select id, customer_id, room_id, start, end, created_at, updated_at from bookings where customer_id = $1;"
)
