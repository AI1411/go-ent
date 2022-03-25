package routes

import (
	"github.com/AI1411/go-ecom/database"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/", database.CreateUser)
}
