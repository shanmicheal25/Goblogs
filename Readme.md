## prerequisite

Docker Desktop need to install in your machine.


## Download the application from the Git hub using this url, (https://github.com/shanmicheal25/Goblogs.git)

    In addition, please copy Angularclient frontend application from this url (https://github.com/shanmicheal25/AngularClient.git) and place into the AngularClient folder (https://github.com/shanmicheal25/Goblogs.git)

    ALl together this application should have three folders (frontend, backend, database) as like below,

           1. AngularClient - frontend
           2. GoBlogs       - backend
           3. Mysql         - Database

    Finally, Folder structure should look like the below.

        AngularClient
        GoBlogs
        MySql
        docker-compose.yml
        GoBlogs.postman_collection.json
        projectFolderStructure.png
        Readme.md

    Each folder have own dockerfile to download their libraries.

    Inside the application, it contains "docker-compose.yml" this file will help to spin the application.

## Run the application, execute the following commend step by step,

    Step 1:
        
        docker-compose build --no-cache
        docker-compose up --build

    Step 2: (please exec once step 1 completed)

        To create user details, use the below API from postman.    

        URL:   http://localhost:9090/api/v1/user      
        Method: POST        
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


## Folder structure should look like below,

    https://github.com/shanmicheal25/Goblogs/blob/master/projectFolderStructure.png

## blogs post page look like below,

    https://github.com/shanmicheal25/Goblogs/blob/master/outputlookandfeel.png
    

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




