CREATE SCHEMA urlshortener;

CREATE TABLE urlshortener.urls(
    id SERIAL PRIMARY KEY ,
    short_code TEXT UNIQUE NOT NULL,
    url TEXT NOT NULL
);

CREATE INDEX idx_short_code ON urlshortener.urls(short_code);