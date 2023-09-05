package userdb

var (
	InsertUserQuery     = "insert into users (id, first_name, last_name, email, phone, document, role, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
	SelectUserByIDQuery = "select id, first_name, last_name, email, phone, document, role, created_at, updated_at from users where id = $1;"
)