package query

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// User returns the 'user' collection from the 'portfolioMakerDev' database
func User(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("user")
}

// PersonalDetails returns the 'personalDetails' collection from the 'portfolioMakerDev' database
func PersonalDetails(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("personalDetails")
}

// Education returns the 'education' collection from the 'portfolioMakerDev' database
func Education(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("education")
}

// Experience returns the 'experience' collection from the 'portfolioMakerDev' database
func Experience(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("experience")
}

// Skill returns the 'skills' collection from the 'portfolioMakerDev' database
func Skill(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("skills")
}

// WhatYouDoing returns the 'whatYouDoing' collection from the 'portfolioMakerDev' database
func WhatYouDoing(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("whatYouDoing")
}

// Testimonials returns the 'testimonials' collection from the 'portfolioMakerDev' database
func Testimonials(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("testimonials")
}

// PortfolioProjects returns the 'portfolioProjects' collection from the 'portfolioMakerDev' database
func PortfolioProjects(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("portfolioProjects")
}

// SocialMedia returns the 'socialMedia' collection from the 'portfolioMakerDev' database
func SocialMedia(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("socialMedia")
}

// UserProfile returns the 'userProfile' collection from the 'portfolioMakerDev' database
func UserProfile(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("userProfile")
}

func WebsiteDeployment(db mongo.Client) *mongo.Collection {
	return db.Database("portfolioMakerDev").Collection("websiteDeployment")
}
