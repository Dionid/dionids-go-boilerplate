-- name: SignInGetUser :one
SELECT id, "password", "role" FROM "user" WHERE email = $1 LIMIT 1;