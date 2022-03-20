package main

import (
	"restapi/router"
)

func main() {
	r := router.Init()

	r.Run() // listen and serve on 0.0.0.0:8080
}
