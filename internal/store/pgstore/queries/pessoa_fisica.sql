-- name: CreatePessoaFisica :one
INSERT INTO pessoa_fisica (
    renda_mensal,
    idade,
    nome_completo,
    celular,
    email,
    categoria,
    saldo,
    user_id
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;


-- name: GetPessoaFisica :one
SELECT *
FROM pessoa_fisica
WHERE id = $1
LIMIT 1;


-- name: ListPessoasFisicas :many
SELECT *
FROM pessoa_fisica
ORDER BY nome_completo
LIMIT $1
OFFSET $2;


-- name: UpdatePessoaFisica :one
UPDATE pessoa_fisica
SET
    renda_mensal = $2,
    idade = $3,
    nome_completo = $4,
    celular = $5,
    email = $6,
    categoria = $7,
    saldo = $8
WHERE id = $1
RETURNING *;


-- name: UpdateSaldoPessoaFisica :one
UPDATE pessoa_fisica
SET saldo = $2
WHERE id = $1
RETURNING *;


-- name: DeletePessoaFisica :exec
DELETE FROM pessoa_fisica
WHERE id = $1;


-- name: GetPessoaFisicaByUserIdAndEmail :one
SELECT *
FROM pessoa_fisica
WHERE user_id = $1 AND email = $2
LIMIT 1;


-- name: GetPessoaFisicaByUserId :many
SELECT *
FROM pessoa_fisica
WHERE user_id = $1;
