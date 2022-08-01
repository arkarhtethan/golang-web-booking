package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/arkarhtethan/golang-web-booking/internal/config"
	"github.com/arkarhtethan/golang-web-booking/internal/handlers"
	"github.com/arkarhtethan/golang-web-booking/internal/models"
	"github.com/arkarhtethan/golang-web-booking/internal/render"
)

const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	fmt.Printf("Server is running at %s\n", PORT)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	gob.Register(models.Reservation{})

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache.")
		return err
	}
	app.TempateCache = tc

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	return nil
}
