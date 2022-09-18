package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		logrus.Errorf("%#+v", err)
		c.JSON(http.StatusOK, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
