

# Save Words 

This is a simple crud using MYSQL and GO    ,
the idea of this program is a aplication for save words
in english and spanish and examples and allow to show and 
delete words 

# configuration MYSQL 
# you can use docker compose 


  1. Configure a file for docker compose
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
  # Connect to mysql service, and enter in password = password
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


7.1 need  Create a Tabla words

```shell   
CREATE TABLE palabras (
  id INT PRIMARY KEY AUTO_INCREMENT,
  palabra VARCHAR(50) NOT NULL,
  traduccion VARCHAR(50) NOT NULL,
  example  VARCHAR(150) NOT NULL
);


```


7.2  You can insert data inside  of table  words 

```shell  

INSERT INTO palabras (palabra, traduccion, example) VALUES 
('hello', 'hola', 'Hello, how are you?'),
('goodbye', 'adiós', 'Goodbye, see you soon!'),
('thank you', 'gracias', 'Thank you for your help.'),
('please', 'por favor', 'Can you pass me the salt, please?'),
('friend', 'amigo', 'He is my best friend.'),
('family', 'familia', 'My family is very important to me.'),
('love', 'amor', 'I love spending time with you.'),
('hate', 'odio', 'I hate cleaning the house.'),
('work', 'trabajo', 'I have a lot of work to do.'),
('study', 'estudiar', 'I need to study for my exam.'),
('eat', 'comer', 'I usually eat lunch at 12 PM.'),
('drink', 'beber', 'Do you want to drink something?'),
('sleep', 'dormir', 'I need to sleep more.'),
('watch', 'ver', 'I like to watch movies on weekends.'),
('listen', 'escuchar', 'I love listening to music.'),
('read', 'leer', 'I read a lot of books last year.'),
('write', 'escribir', 'She likes to write stories.'),
('play', 'jugar', 'My son likes to play with his toys.'),
('travel', 'viajar', 'I want to travel around the world.');

```





8. Finish :)





