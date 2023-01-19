package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/omeg21/project-repo/pkg/config"
	"github.com/omeg21/project-repo/pkg/controller"
	"github.com/omeg21/project-repo/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	// change this to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := controller.NewRepo(&app)
	controller.NewController(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", controller.Repo.Home)
	// http.HandleFunc("/about", controller.Repo.About)
	// http.HandleFunc("/new", controller.New)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
	fmt.Println("what")
}
