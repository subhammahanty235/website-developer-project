package query

import (
	"github.com/subhammahanty235/portfolio-maker/internal/config"
	"github.com/subhammahanty235/portfolio-maker/internal/drivers"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioMakerAppDB struct {
	App *config.PortFolioAppTools
	DB  *mongo.Client
}

func NewPortfolioMakerAppDB(app *config.PortFolioAppTools, db *mongo.Client) drivers.DBRepo {
	return &PortfolioMakerAppDB{
		App: app,
		DB:  db,
	}
}
