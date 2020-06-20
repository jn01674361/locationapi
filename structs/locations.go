package structs

// Location represents a client's location
type Location struct {
	Coordinates Coordinates `json:"coordinates"`
	Metadata    Metadata    `json:"metadata"`
}

// Coordinates holds the latitude and longitude of a location
type Coordinates struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

// Metadata holds additional information about a Location.
type Metadata struct {
	Name *string      `json:"name"`
	Type LocationType `json:"location_type"`
}

// LocationType indicates what category a location belongs to
type LocationType string

const (
	// LocationTypeLeisure is a location associated with leisure activities.
	// For example, an amusement park.
	LocationTypeLeisure LocationType = "leisure"

	// LocationTypeResidence is a location where people live.
	// For example, an apartment building.
	LocationTypeResidence LocationType = "residence"

	// LocationTypePublic is a location where some public authority provides a service.
	// For example, a train station.
	LocationTypePublic LocationType = "public"

	// LocationTypeOther is a location that does not fit into any of the other categories.
	// For example, an unofficial graffiti wall under a bridge.
	LocationTypeOther LocationType = "other"
)

// ValidLocationTypes decides which LocationTypes are valid.
var ValidLocationTypes = map[LocationType]bool{
	LocationTypeLeisure:   true,
	LocationTypePublic:    true,
	LocationTypeResidence: true,
	LocationTypeOther:     true,
}
