package database

var Usuarios string = `CREATE TABLE IF NOT EXISTS usuarios ( 
	id serial, 
	username varchar(255) not null,
	lastname varchar(255) not null,
	primary key (id)
	);`

var index_users string = `CREATE UNIQUE INDEX username_table_users ON usuarios (username);`	

var Friends string = `CREATE TABLE IF NOT EXISTS friends ( 
	id serial, 
	primary key (id),
	IDuser1 varchar(255) not null,
	IDuser2 varchar(255) not null,
	owner int not null,
	foreign key (owner) references usuarios(id)
	ON DELETE CASCADE
	);`
