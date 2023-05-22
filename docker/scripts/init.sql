CREATE DATABASE IF NOT EXISTS chat;

USE chat;

CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(225), 
    name VARCHAR(225), 
    email VARCHAR(225), 
    password VARCHAR(225), 
    PRIMARY KEY(id) 
    );

CREATE TABLE IF  NOT EXISTS messages(
    `id` VARCHAR(225), 
    `date` VARCHAR(225), 
    `content` text, 
    `from` VARCHAR(225), 
    `to` VARCHAR(225), 
    PRIMARY KEY(id) 
);
