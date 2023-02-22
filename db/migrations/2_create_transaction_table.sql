-- +migrate Up
CREATE TABLE IF NOT EXISTS transaction(
    id SERIAL NOT NULL PRIMARY KEY ,
    sender INT,
    beneficiary INT,
    remark VARCHAR(256),
    amount INT
);