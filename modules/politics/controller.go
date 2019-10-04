package politics

import (
	"encoding/json"
	"fmt"
	"gitlab.com/euthinkia/api-ayuntamientos/modules/commons"
	. "gitlab.com/euthinkia/api-ayuntamientos/modules/commons/entities"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

//Controller : Todo
type TownHallsController struct {
	Repository TownHallsRepo
}

func NewTownHallsController() TownHallsController {
	fmt.Println("Ejecutando: func NewTownHallsController() TownHallsController ")
	controller := TownHallsController{}
	controller.Repository = NewTownHallsRepo()
	return controller
}

//GetAllTownHalls : Method to return all townhalls
func (c *TownHallsController) GetAllTownHalls(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Ejecutando func (c *TownHallsController) GetAllTownHalls(w http.ResponseWriter, r *http.Request) {")
	// if townhall := c.Repository.GetAllTownhalls(); townhall != nil {
	townhall, err := c.Repository.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	commons.RespondWithJSON(w, http.StatusCreated, townhall)
}

//InsertTHalls : Methods to insert townhalls
func (c *TownHallsController) InsertTHalls(w http.ResponseWriter, r *http.Request) {
	var townhall TownHall

	defer r.Body.Close()
	townhall.ID = bson.NewObjectId()
	if err := json.NewDecoder(r.Body).Decode(&townhall); err != nil {
		commons.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	// if err := c.Repository.InsertTownHalls(townhall); err != nil {
	if err := c.Repository.Save(townhall); err != nil {
		commons.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	commons.RespondWithJSON(w, http.StatusCreated, townhall)

}
