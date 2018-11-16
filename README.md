
# postwitter-api

A postwitter-api [Go](https://golang.org/) (version 1.11.2) app, developed [echo](github.com/labstack/echo) framework (version v3.1.0). 

This app package include REST API's for user and their post management and it can be easily deployed to Heroku as it includes all requried configurations.

## Features

1. This repo demonstrates the integeration of Go, Echo framwork with MySQL
2. Its using JWT token verification for all internal calls issued on login
3. Its fetching lists Post from MySQL database and display into pagination list
4. It has proper echo framwork middleware and routs defined

## Live Demo

REST API's has been deployed on this link:
https://postwitter-api.herokuapp.com/

And its running interfaces are available at:
https://postwitter-portal.herokuapp.com

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) and the [govendor](https://github.com/kardianos/govendor) installed. It help to add any dependencies to application.

```sh
$ cd $GOPATH/src
$ mkdir postwitter-api
$ git clone git@github.com:abdulbasitmughal/postwitter-api.git
$ go run server.go
```
Output logs will show port number on which server is listening. 

If your see port 5000 then it will available on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku

### Install the Heroku CLI
Download and install the Heroku CLI. If you haven't already, log in to your Heroku account and follow the prompts to create a new SSH public key.

```sh
$ heroku login
```
### Create a new Git repository
Initialize a git repository in a new or existing directory

```sh
$ cd postwitter-api/
$ git init
$ heroku git:remote -a {remote-project-name}
```
### Deploy your application
Commit your code to the repository and deploy it to Heroku using Git.

```sh
$ git add .
$ git commit -am "make it better"
$ git push heroku master
```

## API's Documentation

There are total six API's incorporated to complete this assignment. Here are details:
1. Signup (POST: https://postwitter-api.herokuapp.com/v1/signup)

```bash

curl -H "Content-Type: application/json" -X POST -d "{\"email\":\"basit10@gmail.com\",\"password\":\"asdf\"}" "https://postwitter-api.herokuapp.com/v1/signup"

# => {"id":15,"email":"basit10@gmail.com"}

```

2. Login (POST: https://postwitter-api.herokuapp.com/v1/login)

```bash

curl -H "Content-Type: application/json" -X POST -d "{\"email\":\"basit@gmail.com\",\"password\":\"123\"}" "https://postwitter-api.herokuapp.com/v1/login"

# => {"id":2,"email":"basit@gmail.com","token":"eyJhbGciOiJIUzI1NiIsInRasdfCJ9.eyJlbWFpbCI6ImJhc2l0QGdtYWlsLmNvbSIsImV4cCI6MTU0MjYwNDAwOH0.uGRPdwxn4-7NzL1f9XOCr-v5sQySzSvwN78M9jGh6ZY","timetag":"2018-11-13 20:55:20"}

```

3. Get User List (GET: https://postwitter-api.herokuapp.com/v1/users)

```bash

curl -H "Content-Type: application/json" -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJhc2l0MUBnbWFpbC5jb20iLCJleHAiOjE1NDI1MTMwMjB9.TnxeVIhuVfpaD6d46tAVBsqVs3vx9PVmL6ExMvXiqug"  -X GET "https://postwitter-api.herokuapp.com/v1/users"

# => {"user":[{"id":1,"email":"abdulbasitmughal@gmail.com","timetag":"2018-11-13 20:54:40"}]}

```

4. Get User's Post (GET: https://postwitter-api.herokuapp.com/v1/users/:email/posts)

```bash

curl -H "Content-Type: application/json" -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJhc2l0MUBnbWFpbC5jb20iLCJleHAiOjE1NDI1MTMwMjB9.TnxeVIhuVfpaD6d46tAVBsqVs3vx9PVmL6ExMvXiqug"  -X GET "https://postwitter-api.herokuapp.com/v1/users/:email/posts?limit=5&page=1"

# => {"post":[{"id":0,"email":"basit1@gmail.com","Message":"test message1","TimeTag":"2018-11-16 09:33:34"},{"id":0,"email":"basit1@gmail.com","Message":"test message1","TimeTag":"2018-11-16 09:23:00"}]}

```

5. Get User's Post Feed (GET: https://postwitter-api.herokuapp.com/v1/posts)

```bash

curl -H "Content-Type: application/json" -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJhc2l0MUBnbWFpbC5jb20iLCJleHAiOjE1NDI1MTMwMjB9.TnxeVIhuVfpaD6d46tAVBsqVs3vx9PVmL6ExMvXiqug"  -X GET "https://postwitter-api.herokuapp.com/v1/posts?page=1&limit=2"

# => {"post":[{"id":0,"email":"basit1@gmail.com","Message":"test message1","TimeTag":"2018-11-16 09:33:34"},{"id":0,"email":"basit1@gmail.com","Message":"test message1","TimeTag":"2018-11-16 09:23:00"}]}

```

6. Create Post (POST: https://postwitter-api.herokuapp.com/v1/posts)

```bash

curl -H "Content-Type: application/json" -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJhc2l0MUBnbWFpbC5jb20iLCJleHAiOjE1NDI1MTMwMjB9.TnxeVIhuVfpaD6d46tAVBsqVs3vx9PVmL6ExMvXiqug" -X POST -d "{\"message\":\"test message\"}" "https://postwitter-api.herokuapp.com/v1/posts"

# => {"id":25,"email":"basit1@gmail.com","Message":"test message","TimeTag":"2018-11-16 10:21:36"}

```


