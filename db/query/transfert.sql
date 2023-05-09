-- name: CreateTransfert :one
INSERT INTO transfert (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;


-- name: GetTransfert :one
SELECT * FROM transfert
WHERE id = $1 LIMIT 1;

-- name: ListTransferts :many
SELECT * FROM transfert
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfert :one
UPDATE transfert
  set amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfert :exec
DELETE FROM transfert
WHERE id = $1;