-- name: CreateUser :one
INSERT INTO users (
  username,
  full_name,
  gender,
  email,
  phone_number,
  date_of_birth,
  address
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;