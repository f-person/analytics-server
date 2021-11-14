package main

import (
	"log"
	"os"
)

type App struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewApp() App {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	return App{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}
