-- +migrate Up
CREATE TABLE IF NOT EXISTS account(
    id SERIAL NOT NULL,
    name VARCHAR(256),
    balance INT
);