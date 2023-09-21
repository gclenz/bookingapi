package bookingdb

var (
	InsertBookingQuery              = "insert into bookings (id, room_id, customer_id, start_on, end_on, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7);"
	SelectOverlappingBookingQuery   = "select id, customer_id, room_id, start_on, end_on, created_at, updated_at from bookings where room_id = $1 and (start_on between $2 and $3) and (end_on between $2 and $3);"
	SelectBookingsByRoomIDQuery     = "select id, customer_id, room_id, start_on, end_on, created_at, updated_at from bookings where room_id = $1;"
	SelectBookingsByCustomerIDQuery = "select id, customer_id, room_id, start_on, end_on, created_at, updated_at from bookings where customer_id = $1;"
)
