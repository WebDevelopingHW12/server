package service

import (
	"fmt"
	"os"

	"github.com/WebDevelopingHW12/server/models"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"

	"github.com/WebDevelopingHW12/server/database/database"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			fmt.Println(root)
		}
	}
	database.Start("database/database/test.db")

	//mx.HandleFunc("/api/", apiHandler(formatter)).Methods("GET")
	mx.Handle("/api/", negroni.New(
		negroni.HandlerFunc(models.ValidateMid),
		negroni.HandlerFunc(apiHandler(formatter)),
	))

	mx.HandleFunc("/login", loginHandler).Methods("POST")
	mx.HandleFunc("/register", registerHandler).Methods("POST")

	mx.HandleFunc("/api/films/", filmsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/films/pages", filmsPagesHandler).Methods("GET")
	filmsSubRouter := mx.PathPrefix("/api/films").Subrouter()
	filmsSubRouter.HandleFunc("/{id:[0-9]+}", getFilmsById).Methods("GET")

	mx.HandleFunc("/api/people/", peopleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/people/pages", peoplePagesHandler).Methods("GET")
	peopleSubRouter := mx.PathPrefix("/api/people").Subrouter()
	peopleSubRouter.HandleFunc("/{id:[0-9]+}", getPeopleById).Methods("GET")

	mx.HandleFunc("/api/planets/", planetsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/planets/pages", planetsPagesHandler).Methods("GET")
	planetsSubRouter := mx.PathPrefix("/api/planets").Subrouter()
	planetsSubRouter.HandleFunc("/{id:[0-9]+}", getPlanetsById).Methods("GET")

	mx.HandleFunc("/api/species/", speciesHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/species/pages", speciesPagesHandler).Methods("GET")
	speciesSubRouter := mx.PathPrefix("/api/species").Subrouter()
	speciesSubRouter.HandleFunc("/{id:[0-9]+}", getSpeciesById).Methods("GET")

	mx.HandleFunc("/api/starships/", starshipsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/starships/pages", starshipsPagesHandler).Methods("GET")
	starshipsSubRouter := mx.PathPrefix("/api/starships").Subrouter()
	starshipsSubRouter.HandleFunc("/{id:[0-9]+}", getStarshipsById).Methods("GET")

	mx.HandleFunc("/api/vehicles/", vehiclesHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/vehicles/pages", vehiclesPagesHandler).Methods("GET")
	vehiclesSubRouter := mx.PathPrefix("/api/vehicles").Subrouter()
	vehiclesSubRouter.HandleFunc("/{id:[0-9]+}", getVehiclesById).Methods("GET")
}
