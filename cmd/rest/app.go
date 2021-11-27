package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/f-person/analytics-service/domain"
	"github.com/f-person/analytics-service/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	EventsDB domain.EventsDB
}

func NewApp() App {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	mongoURI := os.Getenv(mongoURIEnvKey)
	client, err := connectToMongoDB(mongoURI)
	if err != nil {
		app.ErrorLog.Fatal(err)
		os.Exit(1)
	}

	return App{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		EventsDB: infrastructure.NewEventsDB(client),
	}
}

func connectToMongoDB(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().
		ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, clientOptions)
}
