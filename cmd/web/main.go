package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/cknolla/go-webserver/internal/config"
	"github.com/cknolla/go-webserver/internal/handlers"
	"github.com/cknolla/go-webserver/internal/helpers"
	"github.com/cknolla/go-webserver/internal/models"
	"github.com/cknolla/go-webserver/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const port = ":8899"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := run()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting webserver on port", port)

	serve := &http.Server {
		Addr:    port,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatalln(err)
}

func run() error {
	// allow storing Reservation type in the session
	gob.Register(models.Reservation{})
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.GetTemplateCache()
	if err != nil {
		log.Fatalln("Can't create template cache", err)
		return err
	}
	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	helpers.NewHelpers(&app)

	return nil
}
