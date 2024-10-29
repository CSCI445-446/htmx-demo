DROP DATABASE htmx_demo;

CREATE USER IF NOT EXISTS 'htmx'@'localhost' IDENTIFIED BY 'password';
CREATE DATABASE IF NOT EXISTS htmx_demo;
GRANT ALL ON `htmx_demo`.* TO 'htmx'@'localhost';

use htmx_demo;

CREATE TABLE customer(
	id INT NOT NULL AUTO_INCREMENT,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	CONSTRAINT PRIMARY KEY(id)
);

INSERT INTO customer (first_name, last_name, email) VALUES ("Mickey", "Mouse", "mmouse@mines.edu"),
	("Donald", "Duck", "dduck@mines.edu"),
	("Eric", "Dattore", "edattore@mines.edu"),
	("John", "Doe", "jdoe@gmail.com"),
	("Jane", "Doe", "jdoe@gmail.com"),
	("Joe", "Schmoe", "jschmoe@gmail.com");
