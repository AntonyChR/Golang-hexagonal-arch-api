CREATE DATABASE IF NOT EXISTS test;
USE test;
CREATE TABLE IF NOT EXISTS todos (
    id VARCHAR(225), 
    title VARCHAR(225), 
    description VARCHAR(225), 
    date VARCHAR(225), 
    done boolean,
    PRIMARY KEY(id) 
    );
DELETE FROM todos;
INSERT INTO todos (id, title, description, date, done) VALUES("id-1","title-1","desc-1","11/12/2021", 1);
INSERT INTO todos (id, title, description, date, done) VALUES("id-2","title-1","desc-1","11/12/2021", 0);
INSERT INTO todos (id, title, description, date, done) VALUES("id-3","title-1","desc-1","11/12/2021", 0);
INSERT INTO todos (id, title, description, date, done) VALUES("id-4","title-1","desc-1","11/12/2021", 0);