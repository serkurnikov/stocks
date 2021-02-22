-- +goose Up

CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE IF NOT EXISTS public.result_currenies
(
    result_currency_id      SERIAL PRIMARY KEY,
    result_currency_raw     jsonb,
    result_currency_display jsonb
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.