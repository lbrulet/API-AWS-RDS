BEGIN;

CREATE TABLE User
(
	id int auto_increment
		primary key,
	username varchar(255) NOT NULL,
	password varchar(255) NOT NULL
);



COMMIT;