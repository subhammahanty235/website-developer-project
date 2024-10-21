package services

import (
	"github.com/gin-gonic/gin"
	"github.com/subhammahanty235/portfolio-maker/internal/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthServices(r *gin.RouterGroup, g *handlers.PortfolioMakerApp, db *mongo.Client) {
	router := r.Use(gin.Logger(), gin.Recovery())

	router.POST("/signup", g.SignUpWithPassword(db))
	router.POST("/loginpw", g.LoginWithPassword(db))
}
