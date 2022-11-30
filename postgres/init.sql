CREATE DATABASE urlshort;
\c urlshort;
CREATE TABLE urls (
                      long_url VARCHAR(200),
                      short_url VARCHAR(10)
);