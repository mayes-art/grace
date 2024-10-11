package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JunController struct{}

func shareComm(n int) []int {
	var ord []int
	ch := make(chan int)

	for i := 0; i < n; i++ {
		go func(ch chan<- int, order int) {
			ch <- order
		}(ch, i)
	}

	for i := range ch {
		ord = append(ord, i)

		if len(ord) == n {
			break
		}
	}

	close(ch)

	return ord
}

func (j *JunController) GetReqHeader(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "OK", "data": shareComm(5)})
}
