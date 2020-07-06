```

docker-compose build --no-cache 

docker-compose up --build
docker exec -it goblogs-cli bash

cd src/GoBlogs/

go run main.go
```

'''
Check the MySql database details,

$ docker exec -it golang_db mysql -udocker -p
Enter password : docker

$ docker exec -it golang_db mysql -uroot -p
Enter password : root

$ show databases;

$ use 


## check the logs in mysql database comments.
$ docker logs golang_db

## container shell access 
$ docker exec -it golang_db bash

$ docker stop golang_db
$ docker start golang_db
$ docker restart golang_db

$ docker rm golang_db 
'''


Delete all the container and images using below command .

Delete all docker containers (must be run before trying to delete images)

$ docker rm $(docker ps -aq)
Delete all docker images:

$ docker rmi $(docker images -q)


''''
## Pain points i faced when I develop the application .
1. Find the valuable query the level of record and proper database design for comments as tree structure.

## one level.
    SELECT * FROM article.comment where postid = 101 and level = 1  order by parentCommentId asc; 

'''


## prerequisite

Docker Desktop need to install in your machine.


## Download the application from the Git hub DockerBlogs,

Here you can find, 3 application in the git folder, 

AngularClient - frontend
GoBlogs       - backend
Mysql         - Database

Each folder have the own dockerfile to download their libraries.

Inside the DockerBlogs, contains docker-compose.yml this file will help to spin the application.

## Run the application use below command.

    docker-compose up --build


## API Specification.

User Details:


POST   http://localhost:9090//api/v1/user              

Request Body: 

{
    "user_id": 101,
    "name": "shankar",
    "email": "shanmicheal25@gmail.com",
    "mobile": "85993401",
    "password": "abcd1234",
    "address": "AMK"
}


POST   http://localhost:9090//api/v1/blogs   



GET    http://localhost:9090//api/v1/blogs            



GET    http://localhost:9090//api/v1/comments          

POST   http://localhost:9090//api/v1/comments      

GET    http://localhost:9090//api/v1/subcomments     

http://localhost:9090/api/v1/user


## Below for debugging
## check the database configuration throw below commands.

docker exec -it golang_db mysql -uroot -p
Enter password : root

mysql> show databases;

mysql> use blogsdb;

mysql> show tables;

## container shell access if need can debug.
$ docker exec -it golang_db bash

$ docker stop golang_db
$ docker start golang_db
$ docker restart golang_db

## Delete the all local images use below commands.

Delete all docker containers (must be run before trying to delete images)

    $ docker rm $(docker ps -aq)

Delete all docker images:

    $ docker rmi $(docker images -q)




