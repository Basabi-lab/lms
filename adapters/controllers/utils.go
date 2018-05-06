package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(http.StatusInternalServerError, &gin.H{"Error": fmt.Sprintf("%v", err)})
}
