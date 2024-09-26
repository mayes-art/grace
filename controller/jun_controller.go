package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JunController struct{}

func (j *JunController) GetReqHeader(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
