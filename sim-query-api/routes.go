package simqueryapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "home page"})
	})
	server.GET("/ev101/api/sims/:msisdn", getSimData)
}
