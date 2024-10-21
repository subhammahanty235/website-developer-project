package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/subhammahanty235/portfolio-maker/internal/drivers/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ga *PortfolioMakerApp) ProtectedRouteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userEmail, _ := ctx.Get("userEmail")
		userID, _ := ctx.Get("userID")
		ctx.JSON(http.StatusOK, gin.H{"userEmail": userEmail, "userID": userID})
	}
}

func (ga *PortfolioMakerApp) GetUserDetails(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, _ := ctx.Get("userID")
		userEmail, _ := ctx.Get("userEmail")
		println(userEmail)
		println(userID)
		type UserDetails struct {
			ID        primitive.ObjectID `json:"_id" bson:"_id"`
			Name      string             `json:"name" Usage:"required"`
			MobileNum string             `json:"mobile_number"`
			Email     string             `json:"email" Usage:"required,alphanumeric"`
		}

		dbCtx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		collection := query.User(*db)
		var userDetails *UserDetails
		filter := bson.M{"email": userEmail}
		findErr := collection.FindOne(dbCtx, filter).Decode(&userDetails)

		if findErr != nil {
			if findErr == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "User not found",
					"success": false,
					"id":      userID,
					"error":   findErr.Error(),
				})
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Some Error Occured, Please try again later",
					"success": false,
					"error":   "Server error occured",
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user":    userDetails,
			"message": "User fetched Successfully",
			"success": true,
			"error":   nil,
		})

	}
}
