package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/subhammahanty235/portfolio-maker/internal/config"
	"github.com/subhammahanty235/portfolio-maker/internal/drivers"
	"github.com/subhammahanty235/portfolio-maker/internal/drivers/query"
	"github.com/subhammahanty235/portfolio-maker/internal/models"
	"github.com/subhammahanty235/portfolio-maker/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioMakerAppDB struct {
	App *config.PortFolioAppTools
	DB  *drivers.DBRepo
}

func NewPortfolioMakerAppDB(app *config.PortFolioAppTools, db *mongo.Client) *PortfolioMakerApp {
	return &PortfolioMakerApp{
		App: app,
		DB:  query.NewPortfolioMakerAppDB(app, db),
	}
}

func (ga *PortfolioMakerApp) SignUpWithPassword(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *models.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Password, _ = utils.Encrypt(user.Password)
		// ok, status, err := ga.DB.AddUser(user)

		dbCtx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var g *query.PortfolioMakerAppDB

		collection := query.User(*db)
		filter := bson.D{{Key: "email", Value: user.Email}}
		var res bson.M

		findErr := collection.FindOne(dbCtx, filter).Decode(&res)
		if findErr != nil {
			if findErr == mongo.ErrNoDocuments {
				user.ID = primitive.NewObjectID()
				_, insertErr := collection.InsertOne(dbCtx, user)

				if insertErr != nil {
					g.App.ErrorLogger.Fatalf("Cannot add user to database %v", insertErr)

				}
				token, err := utils.Generate(user.Email, user.ID)
				if err != nil {
					ctx.JSON(http.StatusUnprocessableEntity, gin.H{
						"message": "Error while creating the token",
						"error":   err,
					})
					return
				}

				ctx.JSON(http.StatusOK, gin.H{
					"token":   token,
					"message": "Registered Successfully",
					"email":   user.Email,
					"success": true,
				})
				return
			}
			g.App.ErrorLogger.Fatal(findErr)
		}

		ctx.JSON(http.StatusFound, gin.H{
			"message": "Exisiting account, go to the login page",
			"success": false,
		})

	}
}

func (ga *PortfolioMakerApp) LoginWithPassword(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type CustomStruct struct {
			EmailId  string `json:"emailId"`
			Password string `json:"password"`
		}

		var reqData *CustomStruct

		if err := ctx.ShouldBindJSON(&reqData); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}

		dbctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		collection := query.User(*db)
		filter := bson.D{{Key: "email", Value: reqData.EmailId}}

		var response bson.M
		findErr := collection.FindOne(dbctx, filter).Decode(&response)
		if findErr != nil {
			if findErr == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "User with this EmailId doesn't exists",
					"error":   findErr,
					"success": false,
				})
				return
			}
		}

		result, _ := utils.VerifyHash(reqData.Password, response["password"].(string))
		if result {

			idObject := response["_id"].(primitive.ObjectID)

			token, err := utils.Generate(response["email"].(string), idObject)

			if err != nil {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Error while creating the token",
					"error":   err,
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"token":   token,
				"message": "Logged in Successfully",
				"email":   response["email"],
				"success": true,
			})
			return
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Wrong Password/Credentials",
				"error":   "Wrong Password",
				"success": false,
			})
			return
		}

	}
}
