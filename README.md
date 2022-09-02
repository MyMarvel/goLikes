# My pet project on Golang
#### Allows users to add likes to other users.

I use Redis for NoSQL database, gin for web framework, gomock for mocking framework.

## Where to look at?
- Check swagger docs to see the public API of the project (run http://localhost:8080/swagger/index.html after the installation)
- [routes/routes.go] - handlers for all routes
- [controllers/database.go] - Redis database implementation
- [controllers/account_test.go] - an example of a unit test
- Don't look at ./mocks subfolder - it is generated automatically, as well as ./docs

## Installation
You should install and run a local redis server first (preferred port: 6379). Then:
```sh
git clone git@github.com:gin-gonic/gin.git
cd likes_handler
```
Change [settings.go] in case you want to use different ports for redis or a web server. Then execute:
```sh
go run main.go
```

## How to use
Send a POST request to http://localhost:8080/api/v1.0/likes/doLike with a json body like
```json
{
    "targetAccount": "Margulan Seissembai",
    "currentAccount": "Nikita Petrov"
}
```
You'll get an updated amount of likes on this user, like
```json
{
    "likes": 1
}
```

Then send it again to observe the warning.
```json
{
    "warning": "Already liked"
}
```

To fetch amount of likes for a user without incrementing it, use http://localhost:8080/api/v1.0/likes/getLikeCount with a json body like:
```json
{
    "targetAccount": "Margulan Seissembai",
}
```

## Troubleshooting
I am still new in Go world, so any improvements or suggestions are appreciated. Just submit a new issue [here][submitIssue].

[submitIssue]: <https://github.com/gin-gonic/gin/issues/new>
[routes/routes.go]: <https://github.com/gin-gonic/gin/blob/master/.golangci.yml>
[controllers/database.go]: <https://github.com/gin-gonic/gin/blob/master/.golangci.yml>
[controllers/account_test.go]: <https://github.com/gin-gonic/gin/blob/master/.golangci.yml>
[settings.go]: <https://github.com/joemccann/dillinger>