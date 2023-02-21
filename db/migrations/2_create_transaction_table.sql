-- +migrate Up
CREATE TABLE IF NOT EXISTS transaction(
    id SERIAL NOT NULL,
    sender VARCHAR(256),
    beneficiary VARCHAR(256),
    remark VARCHAR(256),
    amount INT
);