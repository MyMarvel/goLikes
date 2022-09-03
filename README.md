# My pet project on Golang
#### Allows users to add likes to other users.

I use Redis for NoSQL database, gin for web framework, gomock for mocking framework.

## Anything interesting?
- Check swagger docs to see the public API of the project (run http://localhost:8080/swagger/index.html after the installation)
- [routes/routes.go] - handlers for all routes
- [controllers/database.go] - Redis database implementation
- [tests/account_test.go] - an example of a unit test
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

## API Documentation
### /likes/doLike

#### POST
##### Summary:

Adds a like to an account

##### Description:

Increments likes counter for a person in case the user didn't like them yet

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Body | body | Params:<ul><li><i>TargetAccount</i>: The account name we like.</li><li><i>CurrentAccount</i>: Our account (the person who likes)</li></ul> | Yes | [routes.doLikeParams](#routes.doLikeParams) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Example: <code>{'likes': likesCount}</code> | string |
| 208 | Example: <code>{'warning': 'Already liked'}</code> | string |
| 400 | Example: <code>{'error': 'Error description'}</code> | string |

### /likes/getLikeCount

#### POST
##### Summary:

Get likes for an account

##### Description:

Returns amount of likes collected by a certain account

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| Body | body | Params:<ul><li><i>TargetAccount</i>: The account name we like</li><ul> | Yes | [routes.getLikeCountParams](#routes.getLikeCountParams) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Example: <code>{'likes': likesCount}</code> | string |
| 400 | Example: <code>{'error': 'Error description'}</code> | string |

### Models


#### routes.doLikeParams

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| currentAccount | string |  | Yes |
| targetAccount | string |  | Yes |

#### routes.getLikeCountParams

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| targetAccount | string |  | Yes |

## Troubleshooting
I am still new in Go world, so any improvements or suggestions are appreciated. Just submit a new issue [here][submitIssue].

[submitIssue]: <https://github.com/MyMarvel/goLikes/issues/new>
[routes/routes.go]: <https://github.com/MyMarvel/goLikes/blob/main/routes/routes.go>
[controllers/database.go]: <https://github.com/MyMarvel/goLikes/blob/main/controllers/database.go>
[tests/account_test.go]: <https://github.com/MyMarvel/goLikes/blob/main/tests/account_test.go>
[settings.go]: <https://github.com/MyMarvel/goLikes/blob/main/settings.go>
