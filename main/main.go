package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Handle(http.MethodGet, "/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "555")
	})
	r.Run("localhost:3333")
}
