CREATE DATABASE IF NOT EXISTS chat;
USE chat;
CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(225), 
    name VARCHAR(225), 
    email VARCHAR(225), 
    password VARCHAR(225), 
    PRIMARY KEY(id) 
    );
INSERT INTO users (id, name, email, password) VALUES("id-1","pedro","pedro@email.com","user1");
INSERT INTO users (id, name, email, password) VALUES("id-2","juan","juan@email.com","user2");
INSERT INTO users (id, name, email, password) VALUES("id-3","maria","maria@email.com","user3");
INSERT INTO users (id, name, email, password) VALUES("id-4","luis","luis@email.com","user4");
INSERT INTO users (id, name, email, password) VALUES("id-5","ana","ana@email.com","user5");
INSERT INTO users (id, name, email, password) VALUES("id-6","carlos","carlos@email.com","user6");
INSERT INTO users (id, name, email, password) VALUES("id-7","jose","jose@email.com","user7");
INSERT INTO users (id, name, email, password) VALUES("id-8","diana","diana@email.com","user8");
INSERT INTO users (id, name, email, password) VALUES("id-9","raul","raul@email.com","user9");
INSERT INTO users (id, name, email, password) VALUES("id-10","laura","laura@email.com","user10");
INSERT INTO users (id, name, email, password) VALUES("060be2da-b0a9-48b2-baea-e00b8151c0e5","laura","laura@email.com","user10");
USE chat;

CREATE TABLE IF  NOT EXISTS messages(
    `id` VARCHAR(225), 
    `date` VARCHAR(225), 
    `content` text, 
    `from` VARCHAR(225), 
    `to` VARCHAR(225), 
    PRIMARY KEY(id) 
);

INSERT INTO messages (id, `date`, content, `from`, `to`) VALUES
('1aqaq', '2023-02-01', 'Hey', '060be2da-b0a9-48b2-baea-e00b8151c0e5', 'id-5'),
('qqqq1', '2023-02-01', 'Hola, ¿cómo estás?', '060be2da-b0a9-48b2-baea-e00b8151c0e5', 'id-5'),
('2qqqqq', '2023-02-01', 'Hola Juan, bien gracias ¿y tú?', 'id-5', '060be2da-b0a9-48b2-baea-e00b8151c0e5'),
('3wwwwwwwww', '2023-02-01', 'Estoy bien también, ¿qué has estado haciendo?', '060be2da-b0a9-48b2-baea-e00b8151c0e5', 'id-5'),
('4qqqqqqqqqqqqqq', '2023-02-01', 'He estado ocupada con el trabajo, pero todo bien. ¿Tú?', 'id-5', '060be2da-b0a9-48b2-baea-e00b8151c0e5'),
('5wqwqwq', '2023-02-01', 'Sí, también he estado ocupado con el trabajo. Pero estoy contento de hablar contigo.', '060be2da-b0a9-48b2-baea-e00b8151c0e5', 'id-5'),
('6ddddddddqq', '2023-02-01', 'Igualmente Juan, siempre es agradable hablar contigo.', 'id-5', '060be2da-b0a9-48b2-baea-e00b8151c0e5'),
('7dfdvc', '2023-02-01', 'Oye Ana, ¿quieres tomar un café un día de estos?', '060be2da-b0a9-48b2-baea-e00b8151c0e5', 'id-5'),
('8333sddd', '2023-02-01', '¡Claro! ¿Qué te parece si quedamos el próximo martes a las 10am?', 'id-5', '060be2da-b0a9-48b2-baea-e00b8151c0e5'),
('9aaaaaccccc', '2023-02-01', 'Perfecto, el martes a las 10am entonces. ¡Nos vemos!', '060be2da-b0a9-48b2-baea-e00b8151c0e5', 'id-5');
