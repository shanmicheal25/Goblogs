## prerequisite

Docker Desktop need to install in your machine.


## Download the application from the Git hub using below url,

    Here you can find, 3 application in the git folder, 

        https://github.com/shanmicheal25/AngularClient.git

            AngularClient - frontend

        https://github.com/shanmicheal25/Goblogs.git

            GoBlogs       - backend
            Mysql         - Database

    combine both application into same folder looks below.

        AngularClient
        GoBlogs
        MySql
        docker-compose.yml
        GoBlogs.postman_collection.json
        projectFolderStructure.png
        Readme.md

    Each folder have the own dockerfile to download their libraries.

    Inside the folder, contains docker-compose.yml this file will help to spin the application.

## Run the application, execute the following commend step by step,

    Step 1:
        
        docker-compose build --no-cache
        docker-compose up --build

    Step 2:

        Once docker compose build the application done. create record for only user table using below API     

        POST   http://localhost:9090/api/v1/user              

        Request Body: 

        {
            "user_id": 101,
            "name": "shankar",
            "email": "shanmicheal25@gmail.com",
            "mobile": "85993401",
            "password": "abcd1234",
            "address": "AMK"
        }

    Step 3:
        Open the browser and run the below url,

        http://localhost:4200


## Folder structure should be like below,

    see documentation [here](./projectFolderStructure.png)

    see output looks [here](./outputlookandfeel.png)

## API Specification.

    Postman collection list GoBlogs.postman_collection.json

## User Details:

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

## Post Blogs

POST   http://localhost:9090//api/v1/blogs   

Request Body: 

{
    "user_id": 101,
    "text": "At this point you might be thinking “Why create all the packages, separate files, layers of functions and what not?” — well, the answer "
}

## Read a Blogs 

GET   http://localhost:9090//api/v1/blogs?user_id=101&nextpage=0     


## Post a comment for particular blog post with post_id

POST   http://localhost:9090//api/v1/comments     

Request Body:

{
    "user_id": 101,
    "post_id": 1001,
    "comment_text": "Created a comment for the post.",
    "comment_level": 1,
    "parent_comm_id": 0
}

## Get a comment list using a particular post_id

GET    http://localhost:9090//api/v1/comments?post_id=1001&nextpage=0          

GET    http://localhost:9090//api/v1/subcomments?comment_id=1&nextpage=0   


## Below for debugging
## check the database configuration throw below commands.

docker exec -it golang_db mysql -uroot -p
Enter password : root

mysql> show databases;

mysql> use blogsdb;

mysql> show tables;

## container shell access if need can debug Mysql Db
    docker exec -it golang_db bash

    docker stop golang_db
    docker start golang_db
    docker restart golang_db


## container shell access if need can debug or manually run Golang application
    docker exec -it goblogs-cli bash
    cd src/GoBlogs/
    go run main.go

## Delete the all local images use below commands.

Delete all docker containers (must be run before trying to delete images)

    $ docker rm $(docker ps -aq)

Delete all docker images:

    $ docker rmi $(docker images -q)




