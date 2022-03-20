# Diamanti Rest API Assignment

This is an assignment repo for Diamanti's Software Engineer Internship interview. It shows an example of REST APIs using Golang.

To run this app, you should have latest version of golang. Please run the below commands from the root of this directory in the terminal.

```
go get . // This gets all the dependencies
```

```
go run . // This runs the app
```

This example uses [Golang](https://go.dev/) as the primary language, [Gin](https://github.com/gin-gonic/gin) as a web framework and `JWT` tokens as a way of authenticating users.

The APIs are quite simple. It has an `authentication` endpoint and CRUD endpoints for music `albums`. Below is a list of all the endpoints.
`POST /auth`
`GET /albums`
`GET /album/:id`
`POST /album`
`PUT /album/:id`
`DELETE /album/:id`

The `/auth` endpoint is used to authenticate the user and provide a jwt token in return. Rest of the endpoints requires user to be authenticated by passing the token in query params as `token`.

The authentication for '/ablum' enpoints is done through a middleware which checks if the user is authenticated otherwise sends back an unauthorized response.

The codebase is divided into several layers to make it more maintainable and readable.

`controllers` - All endpoint handlers are defined in this layer.
`middlerware` - All middlerwares are defined in this layer.
`router` - All the routers are defined in this layer.
`service` - All the business logic for different endpoints are defined here.
`util` - Any utility like helper functions or constants are defined here.

I have just started to learn Golang and this is one of the very first assignments I did using golang ecosystem.

### TODOs

- Some part of the code is repeated, take the repeated code out into a common function or constant.
- Make separate model layer to keep the fetch/store data.

### References

go-gin-example - https://github.com/eddycjy/go-gin-example
Developing a RESTful API with Go and Gin - https://go.dev/doc/tutorial/web-service-gin
