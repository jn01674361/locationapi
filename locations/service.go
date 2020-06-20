package locations

import (
	"locationapi/structs"
)

type locationRepository interface {
	ListLocations() []structs.Location
	StoreLocation(structs.Location)
}

type locationService struct {
	repository locationRepository
}

func (ls *locationService) GetLocations() []structs.Location {
	return ls.repository.ListLocations()
}

func (ls *locationService) PostLocation(location structs.Location) {
	ls.repository.StoreLocation(location)
}
