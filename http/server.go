package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// LocationService is an interface for CRUD operations on locations.
type LocationService interface {
	GetLocations()
}

// Server serves
type Server struct {
	locationService LocationService
}

// Open opens
func (s *Server) Open() {
	muxRouter := mux.NewRouter()

	muxRouter.Handle("/", NotImplemented).Methods(http.MethodGet)
	muxRouter.Handle("/locations", GetLocationsFunc).Methods(http.MethodGet)

	http.ListenAndServe(":8080", muxRouter)
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
func NewServer() Server { return Server{} }
