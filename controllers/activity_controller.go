package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.com/devplaylister/playlister-activity/dto"
	"bitbucket.com/devplaylister/playlister-activity/services"
	"github.com/gin-gonic/gin"
)

type ActivityController struct{}

func AddActivity(c *gin.Context) {
	var userSongActivityData dto.UserSongActivityData
	if err := c.ShouldBindJSON(&userSongActivityData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.IndexActivity(userSongActivityData)

	if err != nil {
		fmt.Errorf(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func FindNearestActivities(c *gin.Context) {
	lat, _ := strconv.ParseFloat(c.Param("lat"), 64)
	lon, _ := strconv.ParseFloat(c.Param("lon"), 64)
	distance, _ := strconv.Atoi(c.Param("distance"))
	activities, err := services.FindNearest(lat, lon, distance)

	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, activities)
}
