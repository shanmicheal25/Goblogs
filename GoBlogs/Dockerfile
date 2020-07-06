
FROM golang:latest

RUN apt-get update
RUN apt-get upgrade -y

ENV GOBIN /go/bin

RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/rs/cors"
RUN go get "github.com/joho/godotenv"
RUN go get "github.com/gin-gonic/gin"
RUN go get "github.com/jinzhu/gorm"
RUN go get "github.com/gin-contrib/cors"
RUN go get "github.com/biezhi/gorm-paginator/pagination"
RUN go get "github.com/dgrijalva/jwt-go"
RUN go get "github.com/auth0/go-jwt-middleware"

COPY ./src /go/src/GoBlogs



