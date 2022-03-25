package main

import (
	"github.com/AI1411/go-ecom/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.Router(r)

	r.Run()
}
