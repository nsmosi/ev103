package simqueryapi

import (
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {

	server.GET("/ev101/api/sims/:msisdn", getSimCard)
	server.DELETE("/ev101/api/sims/:msisdn", deleteSimCard)

}
