package router

import (
	"github.com/AlexandreLima658/gopportunities/handler/command/create"
	deleteOpening "github.com/AlexandreLima658/gopportunities/handler/command/delete"
	"github.com/AlexandreLima658/gopportunities/handler/command/update"
	"github.com/AlexandreLima658/gopportunities/handler/query/filter"
	"github.com/AlexandreLima658/gopportunities/handler/query/id"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", id.ShowOpeningHandler)
		v1.POST("/opening", create.CreateOpeningHandler)
		v1.DELETE("/opening", deleteOpening.DeleteOpeningHandler)
		v1.PUT("/opening", update.UpdateOpeningHandler)
		v1.GET("/openings", filter.ListOpeningHandler)

	}
}
