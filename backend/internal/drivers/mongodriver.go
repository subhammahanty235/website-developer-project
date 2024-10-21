package drivers

import (
	"context"
	"time"

	"github.com/subhammahanty235/portfolio-maker/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var app config.PortFolioAppTools

func Connection(URI string) mongo.Client {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10000*time.Microsecond)
	defer cancelCtx()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		println(err)
	}

	return *client
}
