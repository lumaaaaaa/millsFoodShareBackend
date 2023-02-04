

CREATE DATABASE hackathon;

CREATE TABLE IF NOT EXISTS AUTH_TABLE(
    user VARCHAR(15), 
    stat VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS DONATIONS_TABLE(
    item VARCHAR(25),
    needed int,
    current int
);