package api

import (
	"net/http"

	"text-captcha/bot/pkg/logger"
	"text-captcha/bot/pkg/reader"

	"github.com/gin-gonic/gin"
)

var homepath = "./resources/images"

func HandleRequest() {
	r := gin.Default()
	r.GET("/health/", checkHealth)
	r.GET("/image/:filename", readImage)

	r.Run("localhost:8080")
}

func checkHealth(c *gin.Context) {
	logger.Info("Endpoint Hit: checkHealth")

	c.JSON(http.StatusOK, "")
}

func readImage(c *gin.Context) {
	logger.Info("Endpoint Hit: readImage")

	filename := c.Param("filename")
	text, err := reader.ReadImage(filename)

	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	logger.Info(text)
	c.JSON(http.StatusOK, text)
}
