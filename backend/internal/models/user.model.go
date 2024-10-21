package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" Usage:"required"`
	MobileNum string             `json:"mobile_number"`
	Email     string             `json:"email" Usage:"required,alphanumeric"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `json:"created_at"`
}

/*
	Profile Will include fields such as
		1. Personal Details -
		2. Education Details -
		3. Experience -
		4. What You are doing -
		5. Testimonials -
		6. Skills -
		7. Portfolio
		8. Blogs
		9. Social Media

*/

type PersonalDetails struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	// UserProfileId primitive.ObjectID `json:"userProfileId" bson:"userProfileId"`
	Name       string `json:"name"`
	ProfilePic string `json:"profilePic"`
	MobileNum  string `json:"mobile_number"`
	Email      string `json:"email" Usage:"required,alphanumeric"`
	BirthDate  string `json:"birthDate"`
	Location   string `json:"location"`
	Bio        string `json:"bio"`
}

type Education struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	InstituteName string             `json:"instituteName"`
	Course        string             `json:"course"`
	StartYear     string             `json:"startYear"`
	EndYear       string             `json:"endYear"`
	Description   string             `json:"description"`
}

type Experience struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	CompanyName string             `json:"CompanyName"`
	Designation string             `json:"designation"`
	StartYear   string             `json:"startYear"`
	EndYear     string             `json:"endYear"`
	Description string             `json:"description"`
}

type Skill struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	SkillName string             `json:"skillName"`
	Knowledge string             `json:"knowledge"`
}

type WhatYouDoing struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Header    string             `json:"header"`
	SubHeader string             `json:"subHeader"`
	Icon      string             `json:"icon"`
}

type Testimonials struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Header    string             `json:"header"`
	SubHeader string             `json:"subHeader"`
	Image     string             `json:"image"`
	Date      time.Time          `json:"date"`
}

type PortfolioProjects struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	HomePageImage string             `json:"homepageImage"`
	ProjectName   string             `json:"projectName"`
	Category      string             `json:"category"`
	LiveLink      string             `json:"liveLink"`
	GithubLink    string             `json:"githubLink"`
}

type SocialMedia struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	SocialMediaName string             `json:"socialMediaName"`
	Icon            string             `json:"icon"`
	ProfileLink     string             `json:"profileLink"`
}

type UserProfile struct {
	ID                     primitive.ObjectID   `json:"_id" bson:"_id"`
	UserId                 primitive.ObjectID   `json:"userId"`
	PersonalDetailsId      primitive.ObjectID   `json:"personalDetailsId"`
	EducationDetail        []primitive.ObjectID `json:"educationDetaiil"`
	ExperienceDetail       []primitive.ObjectID `json:"experienceDetail"`
	WhatAreYouDoingDetails []primitive.ObjectID `json:"whatAreYouDoingDetails"`
	Testimonials           []primitive.ObjectID `json:"testimonials"`
	Skills                 []primitive.ObjectID `json:"skills"`
	Portfolio              []primitive.ObjectID `json:"portfolio"`
	SocialMedia            []primitive.ObjectID `json:"socialMedia"`
}
