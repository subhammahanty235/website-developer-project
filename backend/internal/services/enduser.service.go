package services

import (
	"github.com/gin-gonic/gin"
	"github.com/subhammahanty235/portfolio-maker/internal/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

func EndUserService(r *gin.RouterGroup, g *handlers.PortfolioMakerApp, db *mongo.Client) {
	router := r.Use(gin.Logger(), gin.Recovery())

	router.POST("/getdetails/profile/:profileId", g.FetchProfileDetails(db))
}
