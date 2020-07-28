package main

import (
	"github.com/gin-gonic/gin"
)

func initRoutes() {
	router.GET("/", showIndex)
}

func showIndex(c *gin.Context) {
	c.JSON(200, jobList)
}
