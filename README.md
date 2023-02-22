

# Save Words 

this is a simple crud using MYSQL and GO ,
the idea of this program is a aplication for save words
in english and spanish and examples and allow to show and 
delete words 

# configuration MYSQL 
i use docker for MYSQL here the configuration in docker 

The first thing is Dowload the image in docker :

docker pull mysql 

2 the next step is build a container whith root user and password password 

docker run --name mysql-database -e MYSQL_ROOT_PASSWORD=password -d mysql:latest

3 You can  then acces the MYSQL container command line using.

docker exec -it mysql-database bash

4 connect to mysql use :

mysql -u root -ppassword

# ready ... Now create the database for the use 

- Create a new database:
SQL Code :

CREATE DATABASE Words_App;

#Table for the users 

CREATE TABLE users (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
    );






