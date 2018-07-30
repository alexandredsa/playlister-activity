package routes

import (
	"os"

	"bitbucket.com/devplaylister/playlister-activity/controllers"
	"github.com/gin-gonic/gin"
)

func SetUp() {
	router := gin.Default()

	router.POST("/activities", controllers.AddActivity)
	router.GET("/activities/nearest/:lat/:lon/:distance", controllers.FindNearestActivities)
	router.Run(os.Getenv("PORT"))
}
