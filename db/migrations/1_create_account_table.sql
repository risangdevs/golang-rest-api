-- +migrate Up
CREATE TABLE IF NOT EXISTS account(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(256),
    balance INT
);