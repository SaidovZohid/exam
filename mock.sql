CREATE DATABASE exam;

DROP DATABASE exam;

CREATE TABLE avtomobile (
    id serial primary key,
    name varchar(50) NOT NULL,
    model varchar(50) NOT NULL,
    year DATE NOT NULL,
    color varchar(100) NOT NULL,
    horse_power INTEGER NOT NULL,
    km numeric(18, 2) DEFAULT 0
);

DROP TABLE avtomobile;

CREATE TABLE images (
    id serial NOT NULL,
    car_id INTEGER NOT REFERENCES avtomobile(id),
    sequence_number INTEGER NOT NULL,
    image_url VARCHAR NOT NULL
);

DROP Table images;

