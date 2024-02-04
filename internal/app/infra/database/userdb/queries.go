package userdb

var (
	InsertUserQuery        = "insert into users (id, first_name, last_name, email, phone, document, date_of_birth, role, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);"
	SelectUserByIDQuery    = "select id, first_name, last_name, email, phone, document, date_of_birth, role, code, code_expiration, created_at, updated_at from users where id = $1;"
	SelectUserByEmailQuery = "select id, first_name, last_name, email, phone, document, date_of_birth, role, code, code_expiration, created_at, updated_at from users where email = $1;"
	UpdateUserCodeQuery    = "update users set code = $1, code_expiration = $2 where email = $3"
)
