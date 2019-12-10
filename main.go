package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/superbkibbles/visa-checker/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	cc := controllers.NewCountryController(getSession())
	r.GET("/countries", cc.GetCountries)
	r.POST("/country", cc.AddCountry)
	r.DELETE("/country/:id", cc.DeleteCountry)
	r.PATCH("/country/:id/:name", cc.EditCountry)
	r.GET("/country/:id", cc.GetCountry)

	handler := setCors(r)
	http.ListenAndServe(getPort(), handler)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}

// Setting up the cors
func setCors(r *httprouter.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(r)
	return handler
}

func getPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
