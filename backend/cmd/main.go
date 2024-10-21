package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/subhammahanty235/portfolio-maker/internal/config"
	"github.com/subhammahanty235/portfolio-maker/internal/drivers"
	"github.com/subhammahanty235/portfolio-maker/internal/handlers"
	"github.com/subhammahanty235/portfolio-maker/internal/services"
)

var app config.PortFolioAppTools

func main() {
	InfoLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger

	err := godotenv.Load()
	if err != nil {
		app.ErrorLogger.Fatal("No .env file available " + err.Error())
	}

	uri := os.Getenv("MONGODB_URI")
	println(uri)

	if uri == "" {
		app.ErrorLogger.Fatal("MongoDb uri string is not available")
	}

	client := drivers.Connection(uri)
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			app.ErrorLogger.Fatal(err)
			return
		}
	}()

	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	portfolioMaker := handlers.NewPortfolioMakerApp(&app, &client)
	services.Services(router, portfolioMaker, &client)

	router.Run(":8080")
}
