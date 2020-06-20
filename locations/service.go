package locations

import (
	"locationapi/structs"
)

type locationRepository interface {
	ListLocations() []structs.Location
	StoreLocation(structs.Location)
}

// LocationService location service
type LocationService struct {
	repository locationRepository
}

// GetLocations gets locations
func (ls LocationService) GetLocations() []structs.Location {
	return ls.repository.ListLocations()
}

// PostLocation posts locations
func (ls LocationService) PostLocation(location structs.Location) {
	ls.repository.StoreLocation(location)
}

// NewLocationService returns a new location
func NewLocationService(repo locationRepository) LocationService {
	return LocationService{
		repository: repo,
	}
}
