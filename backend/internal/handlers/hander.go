package handlers

import (
	"github.com/subhammahanty235/portfolio-maker/internal/config"
	"github.com/subhammahanty235/portfolio-maker/internal/drivers"
	"github.com/subhammahanty235/portfolio-maker/internal/drivers/query"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioMakerApp struct {
	App *config.PortFolioAppTools
	DB  drivers.DBRepo
}

func NewPortfolioMakerApp(app *config.PortFolioAppTools, db *mongo.Client) *PortfolioMakerApp {
	return &PortfolioMakerApp{
		App: app,
		DB:  query.NewPortfolioMakerAppDB(app, db),
	}
}
