package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/superbkibbles/visa-checker/models"
)

type CountryController struct {
	session *mgo.Session
}

// session *mgo.session

func NewCountryController(s *mgo.Session) *CountryController {
	return &CountryController{s}
}

// get all countries
func (cc CountryController) GetCountries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cs := []models.Country{}
	err := cc.session.DB("heroku_69bmrctm").C("countries").Find(nil).All(&cs)

	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(cs)
}

// Get a single country
func (cc CountryController) GetCountry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	// search in the query
	country := models.Country{}
	c := cc.session.DB("heroku_69bmrctm").C("countries")
	c.FindId(bson.ObjectIdHex(id)).One(&country)

	json.NewEncoder(w).Encode(country)
}

// add new country
func (cc CountryController) AddCountry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// w.Header().Set("content-type", "application/json")
	c := models.Country{}

	json.NewDecoder(r.Body).Decode(&c)
	c.ID = bson.NewObjectId()

	cc.session.DB("heroku_69bmrctm").C("countries").Insert(c)

	err := json.NewEncoder(w).Encode(c)
	if err != nil {
		log.Fatalln(err)
	}
}

// Delete a single country by passing id
func (cc CountryController) DeleteCountry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	c := cc.session.DB("heroku_69bmrctm").C("countries")
	// delete that county
	err := c.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		panic(err)
	}
}

// edit a country
func (cc CountryController) EditCountry(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
