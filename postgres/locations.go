package postgres

import "locationapi/structs"

type postgresRepository struct {
	Address string
}

func (p *postgresRepository) ListLocations() []structs.Location {
	res := make([]structs.Location, 0)
	return res
}

func (p *postgresRepository) StoreLocation(structs.Location) {
	return
}
