package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/subhammahanty235/portfolio-maker/internal/drivers/query"
	"github.com/subhammahanty235/portfolio-maker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	We are using one single API to store all types of Profile related data
*/

// API to create a new profile based on the userId

// func (ga *PortfolioMakerApp) ProtectedRouteHandler() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		userEmail, _ := ctx.Get("userEmail")
// 		userID, _ := ctx.Get("userID")
// 		ctx.JSON(http.StatusOK, gin.H{"userEmail": userEmail, "userID": userID})
// 	}
// }

func (ga *PortfolioMakerApp) CreateInitialProfile(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, _ := ctx.Get("userID")

		dbCtx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var g *query.PortfolioMakerAppDB

		collection := query.UserProfile(*db)

		id := primitive.NewObjectID()

		_, insertErr := collection.InsertOne(dbCtx, gin.H{
			"_id":    id,
			"userId": userID,
		})

		if insertErr != nil {
			g.App.ErrorLogger.Fatalf("Cannot add user to database %v", insertErr)

		}

		var res models.UserProfile
		findErr := collection.FindOne(dbCtx, bson.D{{Key: "_id", Value: id}}).Decode(&res)
		if findErr != nil {
			if findErr == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Error while finding the profile",
					"error":   findErr.Error(),
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"profile": res,
			"message": "Created Successfully",
			// "email":   user.Email,
			"success": true,
		})

	}
}

func (ga *PortfolioMakerApp) AddDataToProfile(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, _ := ctx.Get("userID")

		type SegmentType struct {
			Seg string `uri:"seg" binding:"required"`
		}
		var segData *SegmentType
		if err := ctx.ShouldBindUri(&segData); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}

		dbctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.M{"userId": userID}
		switch segData.Seg {
		case "1": //personal details add
			var personalDetailsInput *models.PersonalDetails
			if err := ctx.ShouldBindJSON(&personalDetailsInput); err != nil {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			}

			pdCollection := query.PersonalDetails(*db)
			id := primitive.NewObjectID()

			personalDetailsInput.ID = id
			_, insertErr := pdCollection.InsertOne(dbctx, personalDetailsInput)
			if insertErr != nil {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Error while inserting the profile",
					"error":   insertErr.Error(),
					"success": false,
				})
				return
			}

			profileCollection := query.UserProfile(*db)
			update := bson.M{
				"$set": bson.M{
					"personalDetailsId": id,
				},
			}

			_, updateErr := profileCollection.UpdateOne(dbctx, filter, update)
			if updateErr != nil {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Error while updating the profile",
					"error":   insertErr.Error(),
					"success": false,
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"personalDetails": personalDetailsInput,
				"message":         "added Successfully",
				"success":         true,
			})
			return
		case "2": //Education details add
			var educationDetailsInput []*models.Education
			if err := ctx.ShouldBindJSON(&educationDetailsInput); err != nil {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			}

			var success bool
			var message string
			var error string

			educationCollection := query.Education(*db)

			if len(educationDetailsInput) == 0 {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Please provide all the required details",
					"error":   "Empty data provided",
					"success": false,
				})
				return
			}
			for _, education := range educationDetailsInput {
				id := primitive.NewObjectID()
				education.ID = id
				_, insertErr := educationCollection.InsertOne(dbctx, education)
				if insertErr != nil {
					success = false
					message = "Error while inserting the data"
					error = insertErr.Error()
					break
				}

				profileCollection := query.UserProfile(*db)
				update := bson.M{
					"$push": bson.M{
						"educationDetaiil": id,
					},
				}
				_, updateErr := profileCollection.UpdateOne(dbctx, filter, update)
				if updateErr != nil {
					success = false
					message = "Error while updating the data"
					error = updateErr.Error()
					break
				}

				success = true
				message = "Data added successfully"
				error = "nil"

			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": message,
				"success": success,
				"error":   error,
			})
			return

		case "3": //Experience details add
			var experienceDetailsInput []*models.Experience
			if err := ctx.ShouldBindJSON(&experienceDetailsInput); err != nil {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			}

			var success bool
			var message string
			var error string

			experienceCollection := query.Experience(*db)

			if len(experienceDetailsInput) == 0 {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Please provide all the required details",
					"error":   "Empty data provided",
					"success": false,
				})
				return
			}
			for _, experience := range experienceDetailsInput {
				id := primitive.NewObjectID()
				experience.ID = id
				_, insertErr := experienceCollection.InsertOne(dbctx, experience)
				if insertErr != nil {
					success = false
					message = "Error while inserting the data"
					error = insertErr.Error()
					break
				}

				profileCollection := query.UserProfile(*db)
				update := bson.M{
					"$push": bson.M{
						"experienceDetail": id,
					},
				}
				_, updateErr := profileCollection.UpdateOne(dbctx, filter, update)
				if updateErr != nil {
					success = false
					message = "Error while updating the data"
					error = updateErr.Error()
					break
				}

				success = true
				message = "Data added successfully"
				error = "nil"

			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": message,
				"success": success,
				"error":   error,
			})
			return

		case "4": //What you doing details add
			var whatYouDoingInput []*models.WhatYouDoing
			if err := ctx.ShouldBindJSON(&whatYouDoingInput); err != nil {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			}

			var success bool
			var message string
			var error string

			wydCollection := query.WhatYouDoing(*db)

			if len(whatYouDoingInput) == 0 {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Please provide all the required details",
					"error":   "Empty data provided",
					"success": false,
				})
				return
			}
			for _, wyd := range whatYouDoingInput {
				id := primitive.NewObjectID()
				wyd.ID = id
				_, insertErr := wydCollection.InsertOne(dbctx, wyd)
				if insertErr != nil {
					success = false
					message = "Error while inserting the data"
					error = insertErr.Error()
					break
				}

				profileCollection := query.UserProfile(*db)
				update := bson.M{
					"$push": bson.M{
						"whatAreYouDoingDetails": id,
					},
				}
				_, updateErr := profileCollection.UpdateOne(dbctx, filter, update)
				if updateErr != nil {
					success = false
					message = "Error while updating the data"
					error = updateErr.Error()
					break
				}

				success = true
				message = "Data added successfully"
				error = "nil"

			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": message,
				"success": success,
				"error":   error,
			})
			return

		case "5": //What you doing details add
			var testimonialsInput []*models.Testimonials
			if err := ctx.ShouldBindJSON(&testimonialsInput); err != nil {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			}

			var success bool
			var message string
			var error string

			tmlCollection := query.Testimonials(*db)

			if len(testimonialsInput) == 0 {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Please provide all the required details",
					"error":   "Empty data provided",
					"success": false,
				})
				return
			}
			for _, wyd := range testimonialsInput {
				id := primitive.NewObjectID()
				wyd.ID = id
				_, insertErr := tmlCollection.InsertOne(dbctx, wyd)
				if insertErr != nil {
					success = false
					message = "Error while inserting the data"
					error = insertErr.Error()
					break
				}

				profileCollection := query.UserProfile(*db)
				update := bson.M{
					"$push": bson.M{
						"testimonials": id,
					},
				}
				_, updateErr := profileCollection.UpdateOne(dbctx, filter, update)
				if updateErr != nil {
					success = false
					message = "Error while updating the data"
					error = updateErr.Error()
					break
				}

				success = true
				message = "Data added successfully"
				error = "nil"

			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": message,
				"success": success,
				"error":   error,
			})
			return
		case "6": //What you doing details add
			var skillsInput []*models.Skill
			if err := ctx.ShouldBindJSON(&skillsInput); err != nil {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			}

			var success bool
			var message string
			var error string

			wydCollection := query.Skill(*db)

			if len(skillsInput) == 0 {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Please provide all the required details",
					"error":   "Empty data provided",
					"success": false,
				})
				return
			}
			for _, wyd := range skillsInput {
				id := primitive.NewObjectID()
				wyd.ID = id
				_, insertErr := wydCollection.InsertOne(dbctx, wyd)
				if insertErr != nil {
					success = false
					message = "Error while inserting the data"
					error = insertErr.Error()
					break
				}

				profileCollection := query.UserProfile(*db)
				update := bson.M{
					"$push": bson.M{
						"skills": id,
					},
				}
				_, updateErr := profileCollection.UpdateOne(dbctx, filter, update)
				if updateErr != nil {
					success = false
					message = "Error while updating the data"
					error = updateErr.Error()
					break
				}

				success = true
				message = "Data added successfully"
				error = "nil"

			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": message,
				"success": success,
				"error":   error,
			})
			return
		case "7": //What you doing details add
			var portfolioProjectInput []*models.PortfolioProjects
			if err := ctx.ShouldBindJSON(&portfolioProjectInput); err != nil {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			}

			var success bool
			var message string
			var error string

			wydCollection := query.PortfolioProjects(*db)

			if len(portfolioProjectInput) == 0 {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Please provide all the required details",
					"error":   "Empty data provided",
					"success": false,
				})
				return
			}
			for _, wyd := range portfolioProjectInput {
				id := primitive.NewObjectID()
				wyd.ID = id
				_, insertErr := wydCollection.InsertOne(dbctx, wyd)
				if insertErr != nil {
					success = false
					message = "Error while inserting the data"
					error = insertErr.Error()
					break
				}

				profileCollection := query.UserProfile(*db)
				update := bson.M{
					"$push": bson.M{
						"portfolio": id,
					},
				}
				_, updateErr := profileCollection.UpdateOne(dbctx, filter, update)
				if updateErr != nil {
					success = false
					message = "Error while updating the data"
					error = updateErr.Error()
					break
				}

				success = true
				message = "Data added successfully"
				error = "nil"

			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": message,
				"success": success,
				"error":   error,
			})
			return
		case "8": //What you doing details add
			var smInput []*models.SocialMedia
			if err := ctx.ShouldBindJSON(&smInput); err != nil {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
			}

			var success bool
			var message string
			var error string

			wydCollection := query.SocialMedia(*db)

			if len(smInput) == 0 {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Please provide all the required details",
					"error":   "Empty data provided",
					"success": false,
				})
				return
			}
			for _, wyd := range smInput {
				id := primitive.NewObjectID()
				wyd.ID = id
				_, insertErr := wydCollection.InsertOne(dbctx, wyd)
				if insertErr != nil {
					success = false
					message = "Error while inserting the data"
					error = insertErr.Error()
					break
				}

				profileCollection := query.UserProfile(*db)
				update := bson.M{
					"$push": bson.M{
						"socialMedia": id,
					},
				}
				_, updateErr := profileCollection.UpdateOne(dbctx, filter, update)
				if updateErr != nil {
					success = false
					message = "Error while updating the data"
					error = updateErr.Error()
					break
				}

				success = true
				message = "Data added successfully"
				error = "nil"

			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": message,
				"success": success,
				"error":   error,
			})
			return
		}
	}
}

func (ga *PortfolioMakerApp) FetchProfileOfUsers(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDStr, _ := ctx.Get("userID")
		println(userIDStr.(string))

		// Convert the userID string to primitive.ObjectID
		userID, _ := primitive.ObjectIDFromHex(userIDStr.(string))

		println(userID.String())

		dbCtx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// var g *query.PortfolioMakerAppDB

		collection := query.UserProfile(*db)

		pipeline := mongo.Pipeline{
			// Match the specific user profile
			{{Key: "$match", Value: bson.D{{Key: "userId", Value: userID.Hex()}}}},

			// Lookup PersonalDetails
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "personalDetails"},
				{Key: "localField", Value: "personalDetailsId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "personalDetails"},
			}}},

			// Lookup EducationDetail array elements
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "education"},
				{Key: "localField", Value: "educationDetaiil"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "educationDetail"},
			}}},

			// Lookup Experience
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "experience"},
				{Key: "localField", Value: "experienceDetail"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "experienceDetails"},
			}}},
			// Lookup WhatYouDoing
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "whatYouDoing"},
				{Key: "localField", Value: "whatAreYouDoingDetails"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "whatAreYouDoingDetails"},
			}}},
			// Lookup Testimonials
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "testimonials"},
				{Key: "localField", Value: "testimonials"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "testimonialsDetails"},
			}}},
			// Lookup Skills
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "skills"},
				{Key: "localField", Value: "skills"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "skillsDetails"},
			}}},
			// Lookup Portfolio
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "portfolioProjects"},
				{Key: "localField", Value: "portfolio"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "portfolioDetails"},
			}}},
			// Lookup SocialMedia
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "socialMedia"},
				{Key: "localField", Value: "socialMedia"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "socialMediaDetails"},
			}}},
		}

		cursor, err := collection.Aggregate(dbCtx, pipeline)
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "Error while finding the profile",
				"error":   err.Error(),
			})
			return
		}

		var details []bson.M
		defer cursor.Close(dbCtx)

		for cursor.Next(dbCtx) {
			var city bson.M

			if err := cursor.Decode(&city); err != nil {
				ctx.JSON(500, gin.H{"error": "Failed to decode city"})
				return
			}
			details = append(details, city)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"profile": details,
			"message": "Created Successfully",
			"success": true,
		})

	}
}

// end  user's website
func (ga *PortfolioMakerApp) FetchProfileDetails(db *mongo.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		type ProfileIDInput struct {
			ProfileId string `uri:"profileId" binding:"required"`
		}

		var profileId *ProfileIDInput
		if err := ctx.ShouldBindUri(&profileId); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}

		profileIdObj, _ := primitive.ObjectIDFromHex(profileId.ProfileId)
		dbCtx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		collection := query.UserProfile(*db)
		pipeline := mongo.Pipeline{
			// Match the specific user profile
			{{Key: "$match", Value: bson.D{{Key: "_id", Value: profileIdObj}}}},

			// Lookup PersonalDetails
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "personalDetails"},
				{Key: "localField", Value: "personalDetailsId"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "personalDetails"},
			}}},

			// Lookup EducationDetail array elements
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "education"},
				{Key: "localField", Value: "educationDetaiil"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "educationDetail"},
			}}},

			// Lookup Experience
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "experience"},
				{Key: "localField", Value: "experienceDetail"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "experienceDetails"},
			}}},

			// Lookup WhatYouDoing
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "whatYouDoing"},
				{Key: "localField", Value: "whatAreYouDoingDetails"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "whatAreYouDoingDetails"},
			}}},

			// Lookup Testimonials
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "testimonials"},
				{Key: "localField", Value: "testimonials"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "testimonialsDetails"},
			}}},

			// Lookup Skills
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "skills"},
				{Key: "localField", Value: "skills"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "skillsDetails"},
			}}},

			// Lookup Portfolio
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "portfolioProjects"},
				{Key: "localField", Value: "portfolio"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "portfolioDetails"},
			}}},

			// Lookup SocialMedia
			{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "socialMedia"},
				{Key: "localField", Value: "socialMedia"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "socialMediaDetails"},
			}}},

			// Limit to one document
			{{Key: "$limit", Value: 1}},
		}

		cursor, err := collection.Aggregate(dbCtx, pipeline)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "Error while finding the profile",
					"error":   err.Error(),
				})
				return
			}
			ctx.JSON(500, gin.H{"error": "Failed to aggregate profile"})
			return
		}

		var details bson.M
		defer cursor.Close(dbCtx)

		if cursor.Next(dbCtx) {
			if err := cursor.Decode(&details); err != nil {
				ctx.JSON(500, gin.H{"error": "Failed to decode profile"})
				return
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Profile not found"})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"profile": details,
			"message": "Created Successfully",
			"success": true,
		})

	}
}
