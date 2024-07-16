package simqueryapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getSimData(context *gin.Context) {

	msisdn, _ := strconv.ParseInt(context.Param("msisdn"), 10, 64)

	context.JSON(http.StatusOK, gin.H{"message": "test was successful", "msisdn": msisdn})
}
