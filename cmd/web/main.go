package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/cknolla/go-webserver/pkg/config"
	"github.com/cknolla/go-webserver/pkg/handlers"
	"github.com/cknolla/go-webserver/pkg/render"
	"log"
	"net/http"
	"time"
)

const port = ":8899"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.GetTemplateCache()
	if err != nil {
		log.Fatalln("Can't create template cache", err)
	}
	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Println("Starting webserver on port", port)

	serve := &http.Server {
		Addr: port,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatalln(err)
}
