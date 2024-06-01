CREATE TABLE user (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    nickname VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL,
    gender VARCHAR(255) NOT NULL,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);