package services

import (
	"github.com/gin-gonic/gin"
	"github.com/subhammahanty235/portfolio-maker/internal/handlers"
	"github.com/subhammahanty235/portfolio-maker/internal/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserServices(r *gin.RouterGroup, g *handlers.PortfolioMakerApp, db *mongo.Client) {
	router := r.Use(gin.Logger(), gin.Recovery(), middlewares.JWTMiddleware())

	router.POST("/create", g.CreateInitialProfile(db))
	router.POST("/adddata/:seg", g.AddDataToProfile(db))
	router.GET("/getuser", g.GetUserDetails(db))
	router.GET("/getprofiles", g.FetchProfileOfUsers(db))

}
