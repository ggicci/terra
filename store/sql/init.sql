CREATE DATABASE terra WITH OWNER = 'root' ENCODING = 'utf8';
SET search_path TO terra;

-- https://dba.stackexchange.com/questions/76655/postgresql-multi-column-gin-index
CREATE EXTENSION IF NOT EXISTS btree_gin;
-- https://www.postgresql.org/docs/13/uuid-ossp.html
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Global id sequence for all my tables.
-- https://ggicci.me/posts/postgres-id-generator-with-safe-json-numbers/
-- Let's use a "SAFE" range. For JOSN compatible, we use only 53 bits.
-- 32: timestamp in seconds
--  5: shard bits, max 2^5=32 shards
-- 16: seq bits, max ops 2^16=65536, means max write 65k ops/s
CREATE SEQUENCE IF NOT EXISTS public.global_id_sequence;
CREATE OR REPLACE FUNCTION public.id_generator(OUT result bigint) AS $$
DECLARE
    seq_id bigint;
    now_seconds bigint;
    shard_id int := 1;
BEGIN
    SELECT nextval('public.global_id_sequence') % 65536 INTO seq_id;

    SELECT FLOOR(EXTRACT(EPOCH FROM CLOCK_TIMESTAMP())) INTO now_seconds;
    result := now_seconds << 21; -- Shard(5) + Seq(16)
    result := result | (shard_id << 16); -- Seq(16)
    result := result | (seq_id);
END;
$$ LANGUAGE PLPGSQL;

-- NOTE(ggicci): DONT forget to grant privileges on `public.global_id_sequence` to roles.

CREATE TABLE IF NOT EXISTS users
(
    id            BIGINT            PRIMARY KEY DEFAULT id_generator(),
    created_at    TIMESTAMPTZ       NOT NULL DEFAULT '0001-01-01T00:00:00Z',
    updated_at    TIMESTAMPTZ       NOT NULL DEFAULT '0001-01-01T00:00:00Z',
    github_id     BIGINT            NOT NULL DEFAULT 0,
    login         TEXT              NOT NULL DEFAULT '',
    display       TEXT              NOT NULL DEFAULT '',
    email         TEXT              NOT NULL DEFAULT '',
    location      TEXT              NOT NULL DEFAULT '',
    company       TEXT              NOT NULL DEFAULT '',
    avatar        TEXT              NOT NULL DEFAULT ''
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users__github_id ON users (github_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users__username ON users (login);