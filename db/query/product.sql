-- name: CreateProduct :one
INSERT INTO products (
  user_id,
  name,
  description,
  start_price,
  images,
  watchers
) VALUES (
  $1, $2, $3, $4, $5, $6
)RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
LIMIT $1
OFFSET $2;


-- name: ListUserProducts :many
SELECT * from products
WHERE user_id = $1
ORDER BY id;


