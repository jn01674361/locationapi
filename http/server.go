package http

import (
	"encoding/json"
	"locationapi/structs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// LocationService is an interface for CRUD operations on locations.
type LocationService interface {
	GetLocations() []structs.Location
	PostLocation(structs.Location)
}

// Server serves
type Server struct {
	locationService LocationService
}

// Open opens
func (s *Server) Open() {
	muxRouter := mux.NewRouter()

	muxRouter.Handle("/", NotImplemented).Methods(http.MethodGet)
	muxRouter.Handle("/locations", s.getLocations()).Methods(http.MethodGet)
	muxRouter.Handle("/locations", GetLocationsFunc).Methods(http.MethodPost)

	http.ListenAndServe(":8080", muxRouter)
}

func (s Server) getLocations() http.HandlerFunc {
	log.Println("about get locations")
	return func(w http.ResponseWriter, r *http.Request) {
		res := s.locationService.GetLocations()
		j, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		w.Write([]byte(j))
	}

}

func (s Server) postLocations() http.HandlerFunc {
	log.Println("about to post locations")
	return func(w http.ResponseWriter, r *http.Request) {
		s.locationService.PostLocation(structs.Location{})
		w.Write([]byte("byte"))
	}

}

// GetLocationsFunc gets locations.
var GetLocationsFunc = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

// NotImplemented is not implemented yet
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

// NewServer returns a new router.
func NewServer(location LocationService) Server {
	return Server{locationService: location}
}
