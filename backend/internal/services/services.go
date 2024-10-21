package services

import (
	"github.com/gin-gonic/gin"
	"github.com/subhammahanty235/portfolio-maker/internal/handlers"

	"go.mongodb.org/mongo-driver/mongo"
)

func Services(r *gin.Engine, g *handlers.PortfolioMakerApp, db *mongo.Client) {
	authGroup := r.Group("/auth")
	{
		AuthServices(authGroup, g, db)
	}
	userGroup := r.Group("/user")
	{
		UserServices(userGroup, g, db)
	}
	deployGroup := r.Group("/deploy")
	{
		WebsiteService(deployGroup, g, db)
	}
	enduserGroup := r.Group("/enduser")
	{
		EndUserService(enduserGroup, g, db)
	}
}
