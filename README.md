# rest-api
Template REST api

## To run

```bash
$ make local
```

## Frame get health

## Collect metrics

## Add opentracing

## Add docs api by swagger

go get swagger: `go get -u github.com/swaggo/swag/cmd/swag`

generate doc: `swag init -g cmd/server/main.go`

Add some description:

```go
// @title Go Learn REST API
// @version 1.0
// @description Golang REST API
// @contact.name Duc PX
// @contact.url https://github.com/ducpx
// @contact.email ducpx96@gmail.com
// @BasePath /api/v1
```

In handler, register docs api to server

```go
import (
    _ "github.com/ducpx/rest-api/docs"
    echoSwagger "github.com/swaggo/echo-swagger"
)

...

    e.GET("/swagger/*", echoSwagger.WrapHandler)
```