package services

import (
	"github.com/gin-gonic/gin"
	"github.com/subhammahanty235/portfolio-maker/internal/handlers"
	"github.com/subhammahanty235/portfolio-maker/internal/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
)

func WebsiteService(r *gin.RouterGroup, g *handlers.PortfolioMakerApp, db *mongo.Client) {
	router := r.Use(gin.Logger(), gin.Recovery(), middlewares.JWTMiddleware())

	router.POST("/deploy", g.DeployWebsite(db))
	router.POST("/checkeligiblity", g.CheckEligiblityForDeployment(db))
}
