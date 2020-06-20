package postgres

import (
	"database/sql"
	"fmt"
	"locationapi/structs"

	"github.com/google/uuid"
)

// Repository p
type Repository struct {
	db *sql.DB
}

// ListLocations l
func (p Repository) ListLocations() []structs.Location {
	res := make([]structs.Location, 0)
	sqlStatement := "SELECT * FROM locations;"

	var id uuid.UUID
	var lat, lon float32
	var name, locationType string

	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&id, &lat, &lon, &name, &locationType); err {
		case sql.ErrNoRows:
			panic("No rows were returned!")
		case nil:
			loc := structs.Location{
				ID: id,
				Coordinates: structs.Coordinates{
					Lat: lat,
					Lon: lon,
				},
				Metadata: structs.Metadata{
					// Type: locationType,
					// Name: name,
				},
			}
			res = append(res, loc)
		default:
			panic(err)

		}
	}

	return res
}

// StoreLocation s
func (p Repository) StoreLocation(structs.Location) {

	sqlStatement := `
		INSERT INTO locations (id, latitude, longitude, location_name, location_type)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := p.db.Exec(sqlStatement, uuid.New(), 10.2, 1.1, "King's Cross Station", "public")
	if err != nil {
		panic(err)
	}

	return
}

func buildSQL(ID uuid.UUID) string {
	return fmt.Sprintf("SELECT * FROM locations WHERE id='%s';", ID)
}

// NewRepository n
func NewRepository(database *sql.DB) Repository {
	return Repository{
		db: database,
	}
}
