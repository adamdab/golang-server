package application

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dabkoa/golang-server/application/handlers"
	"github.com/dabkoa/golang-server/utils"
)

type App struct {
	Configuration utils.Config
	Service       *http.Server
}

func CreateApp(configPath string) App {
	var config utils.Config
	err := config.Load("resources/config.json")
	if err != nil {
		panic(err)
	}

	log.Println("Current configuration:")
	log.Printf("%+v\n", config)

	svc := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Port),
		Handler:        handlers.CreateHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return App{
		Service:       svc,
		Configuration: config,
	}
}

func (app *App) Start() {
	log.Printf("Application: \"%s\" started!\n", app.Configuration.ApplicationName)
	log.Fatal(app.Service.ListenAndServe())
}
