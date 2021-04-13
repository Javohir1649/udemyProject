package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Javohir1649/udemyProject/internal/config"
	"github.com/Javohir1649/udemyProject/internal/driver"
	"github.com/Javohir1649/udemyProject/internal/handlers"
	"github.com/Javohir1649/udemyProject/internal/models"
	"github.com/Javohir1649/udemyProject/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	db, err := driver.ConnectSql("host=localhost port=5432 dbname=bookings user=tcs password=")
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	defer db.SQL.Close()

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)

	// http.HandleFunc("/", handlers.Repo.Homepage)
	// http.HandleFunc("/about", handlers.Repo.Aboutpage)
	fmt.Printf("starting server on port %s", portNumber)
	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
