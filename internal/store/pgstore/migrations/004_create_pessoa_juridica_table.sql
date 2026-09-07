-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS pessoa_juridica(
 id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
 faturamento FLOAT NOT NULL,
 idade INT NOT NULL,
 nome_fantasia  TEXT UNIQUE NOT NULL,
 celular VARCHAR(20) NOT NULL,
 email_corporativo TEXT UNIQUE NOT NULL,
 categoria VARCHAR(50) NOT NULL,
 saldo FLOAT NOT NULL,
 user_id UUID NOT NULL REFERENCES users (id)
);
---- create above / drop below ----
DROP TABLE IF EXISTS pessoa_juridica;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
