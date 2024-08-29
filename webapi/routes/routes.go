package main

import (
	"net/http"
	"webapi/trail"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/trails/:id", func(c *gin.Context) {
		id := c.Param("id")
		trail := trail.GetTrailByID(id)
		if trail == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Trail not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, trail)
	})

	router.GET("/freebiketrails", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, trail.GetFreeBikingTrails())
	})

	router.GET("/difficultbiketrails", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, trail.GetDifficultBikeTrails())
	})

	router.GET("/freefishingtrails", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, trail.GetFreeFishingTrails())
	})

	router.GET("/picnicwithrestroom", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, trail.GetPicnicWithRestrooms())
	})

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Welcome to Boulder Trails..")
	})

	return router
}
