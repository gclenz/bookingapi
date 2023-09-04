package roomdb

var (
	InsertRoomQuery     = "insert into rooms (id, name, single_bed_count, double_bad_count, guests_limit, are_pets_allowed, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8);"
	SelectRoomByIDQuery = "select id, name, single_bed_count, double_bad_count, guests_limit, are_pets_allowed, created_at, updated_at from rooms where id = $1;"
)
