package store

//TODO: Cambiar a GIN
import (
	"github.com/gorilla/mux"
	"gitlab.com/euthinkia/api-ayuntamientos/modules/politics"
	// "gitlab.com/euthinkia/api-ayuntamientos/places"
	"log"
	"net/http"
)

var townHallsController = politics.NewTownHallsController()

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Method to get all townhalls",
		"GET",
		"/api/v1/townhalls/all",
		townHallsController.GetAllTownHalls,
	}, {
		"Method to insert townhalls",
		"POST",
		"/api/v1/townhalls/insert",
		townHallsController.InsertTHalls,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
