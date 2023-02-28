

# Save Words 

This is a simple crud using MYSQL and GO ,
the idea of this program is a aplication for save words
in english and spanish and examples and allow to show and 
delete words 

# configuration MYSQL 
# you can use docker compose 


  1. Comfigure a file for docker compose
  for example: we need a file compose.yml in the directory proyect 
```shell
# version of docker compose  v2.3.3

version: '3.7'

services:
  db:
    container_name: mysql
    image: mysql:8.0
    cap_add:
      - SYS_NICE
    restart: always
    environment:
      - MYSQL_DATABASE=Test
      - MYSQL_ROOT_PASSWORD=password
    ports:
      - '3306:3306'
    volumes:
      - db:/var/lib/mysql
volumes:
  db:
    driver: local
```
  2. Run docker :
```shell
    $ docker-compose up -d
```
  3. Connect to docker container and Connect to mysql use your password: password 
```shell
  # Enter in the docker container
  $ docker exec -it mysql bash
  # Connect to mysql service, and enter in password 12345678
  $ mysql -h 127.0.0.1 -u root -P 3306 -p
```
  4. Other commands maybe needed:
  
```shell
   
  # you can stop service by the command
  $ docker stop mysql
  # you can start service by the command
  $ docker start mysql
  # you can restart service by the command
  $ docker restart mysql

```

5. Ready ... Now Use the database

```shell
  use  Words_App;   
```

6. Create the Table for the users 
```shell
CREATE TABLE users (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
    );
```

7. Insert Data  inside of the database 

```shell   
INSERT INTO users (username,password)
VALUES("caleb","12345");
```
8. Finish :)








