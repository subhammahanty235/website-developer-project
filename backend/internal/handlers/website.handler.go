package handlers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/subhammahanty235/portfolio-maker/internal/drivers/query"
	"github.com/subhammahanty235/portfolio-maker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ga *PortfolioMakerApp) DeployWebsite(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// userID, _ := ctx.Get("userID")

		var deploymentInput *models.WebsiteData
		if err := ctx.ShouldBindJSON(&deploymentInput); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}

		dbctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		apiUrl := os.Getenv("API_URL")
		apiToken := os.Getenv("DEPLOYMENT_API_TOKEN")
		repoUrl := os.Getenv("REPO_URL") //currently using only one repo or website template
		repoId := os.Getenv("REPO_ID")

		print(apiUrl)
		print(apiToken)

		profileId := deploymentInput.ProfileId
		websiteName := deploymentInput.Name

		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", "Bearer "+apiToken).
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]interface{}{
				"name": websiteName,
				"gitSource": map[string]interface{}{
					"type":   "github",
					"repoId": repoId,
					"repo":   repoUrl,
					"ref":    "master",
				},
				"builds": []map[string]interface{}{
					{
						"src": "package.json",
						"use": "@vercel/node",
					},
				},
				"env": map[string]string{
					"NODE_ENV":   "production",
					"PROFILE_ID": profileId.Hex(),
				},
			}).
			Post(apiUrl)
		print("deployment code")
		print(resp.RawResponse)

		if err != nil {
			id := primitive.NewObjectID()
			deploymentInput.ID = id
			deploymentInput.DeploymentStatus = false

			webDeployment := query.WebsiteDeployment(*db)

			_, insertErr := webDeployment.InsertOne(dbctx, deploymentInput)

			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message":     "Error while deplying the website",
				"error":       err.Error(),
				"insertError": insertErr.Error(),
				"success":     false,
			})
			return
		}

		id := primitive.NewObjectID()

		deploymentInput.ID = id
		deploymentInput.DeploymentStatus = true

		webDeployment := query.WebsiteDeployment(*db)

		_, insertErr := webDeployment.InsertOne(dbctx, webDeployment)

		if insertErr != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Error while deplying the website",
				"error":   insertErr.Error(),
				"success": false,
			})
			return
		}
		print("Deployment completed")
		ctx.JSON(http.StatusOK, gin.H{
			"websiteDetails": deploymentInput,
			"message":        "added Successfully",
			"success":        true,
			"response":       resp.Body(),
		})

	}
}

func (ga *PortfolioMakerApp) CheckEligiblityForDeployment(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type ProfileIdInput struct {
			ProfileId primitive.ObjectID `json:"profileId"`
		}

		var profileIdInput *ProfileIdInput
		if err := ctx.ShouldBindJSON(&profileIdInput); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}

		dbctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.M{"_id": profileIdInput.ProfileId}

		var responseData *models.UserProfile

		userCollection := query.UserProfile(*db)
		print(profileIdInput.ProfileId.Hex())

		findErr := userCollection.FindOne(dbctx, filter).Decode(&responseData)
		if findErr != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Error while creating the token xx",
				"error":   findErr.Error(),
			})
			return
		}

		// Check if eligible
		var eligible bool
		if responseData.PersonalDetailsId == primitive.NewObjectID() && len(responseData.EducationDetail) > 0 && len(responseData.ExperienceDetail) > 0 &&
			len(responseData.Skills) > 0 && len(responseData.Testimonials) > 0 && len(responseData.SocialMedia) > 0 {
			eligible = true
		} else {
			eligible = false
		}

		//check if the website is already deployed
		webDeployment := query.WebsiteDeployment(*db)
		var websiteData *models.WebsiteData

		websiteFindErr := webDeployment.FindOne(dbctx, bson.M{"profileId": profileIdInput.ProfileId}).Decode(&websiteData)
		if websiteFindErr != nil {
			print("Deployment not found, okay!")
		} else {
			if websiteData.DeploymentStatus {
				ctx.JSON(http.StatusOK, gin.H{
					"message":         "Fetched Successfully",
					"eligible":        false,
					"success":         true,
					"alreadyDeployed": true,
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{

			"message":  "Fetched Successfully",
			"eligible": eligible,
			"success":  true,
		})

	}

}
