-- name: CreatePessoaJuridica :one
INSERT INTO pessoa_juridica (
    faturamento,
    idade,
    nome_fantasia,
    celular,
    email_corporativo,
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


-- name: GetPessoaJuridica :one
SELECT *
FROM pessoa_juridica
WHERE id = $1
LIMIT 1;


-- name: GetPessoaJuridicaByEmail :one
SELECT *
FROM pessoa_juridica
WHERE email_corporativo = $1
LIMIT 1;


-- name: GetPessoaJuridicaByUserID :one
SELECT *
FROM pessoa_juridica
WHERE id = $1
  AND user_id = $2
LIMIT 1;


-- name: GetPessoaJuridicaByUserId :many
SELECT *
FROM pessoa_juridica
WHERE user_id = $1;


-- name: ListPessoasJuridicas :many
SELECT *
FROM pessoa_juridica
ORDER BY nome_fantasia
LIMIT $1
OFFSET $2;


-- name: ListPessoasJuridicasByUserID :many
SELECT *
FROM pessoa_juridica
WHERE user_id = $1
ORDER BY nome_fantasia
LIMIT $2
OFFSET $3;


-- name: UpdatePessoaJuridica :one
UPDATE pessoa_juridica
SET
    faturamento = $2,
    idade = $3,
    nome_fantasia = $4,
    celular = $5,
    email_corporativo = $6,
    categoria = $7,
    saldo = $8
WHERE id = $1
RETURNING *;


-- name: UpdateSaldoPessoaJuridica :one
UPDATE pessoa_juridica
SET saldo = $2
WHERE id = $1
RETURNING *;


-- name: AddSaldoPessoaJuridica :one
UPDATE pessoa_juridica
SET saldo = saldo + $2
WHERE id = $1
RETURNING *;


-- name: RemoveSaldoPessoaJuridica :one
UPDATE pessoa_juridica
SET saldo = saldo - $2
WHERE id = $1
  AND saldo >= $2
RETURNING *;


-- name: DeletePessoaJuridica :exec
DELETE FROM pessoa_juridica
WHERE id = $1;
