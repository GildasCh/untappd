package untappd

import "net/url"

// BreweryService is a "service" which allows access to API methods involving
// breweries.
type BreweryService struct {
	client *Client
}

// Brewery represents an Untappd brewery, and contains information about a
// brewery's name, location, logo, and various other metadata.
type Brewery struct {
	ID           int
	Name         string
	Slug         string
	Logo         url.URL
	Country      string
	Active       bool
	Location     BreweryLocation
	Contact      BreweryContact
	Type         string
	TypeID       int
	Independent  bool
	InProduction int
	Rating       BreweryRating
	Description  string
	Stats        BreweryStats
}

// BreweryLocation represent's an Untappd brewery's location, and contains
// information such as the brewery's city, state, and latitude/longitude.
type BreweryLocation struct {
	Address    string  `json:"brewery_address"`
	City       string  `json:"brewery_city"`
	State      string  `json:"brewery_state"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lng"`
	BLatitude  float64 `json:"brewery_lat"`
	BLongitude float64 `json:"brewery_lng"`
}

type BreweryRating struct {
	Count int     `json:"count"`
	Score float64 `json:"rating_score"`
}

type BreweryStats struct {
	TotalCount   int     `json:"total_count"`
	UniqueCount  int     `json:"unique_count"`
	MonthlyCount int     `json:"monthly_count"`
	WeeklyCount  int     `json:"weekly_count"`
	AgeOnService float64 `json:"age_on_service"`
}

// BreweryContact represents an Untappd brewery's contact social media
// and website contact information.
type BreweryContact struct {
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	URL       string `json:"url"`
}

// rawBrewery is the raw JSON representation of an Untappd brewery.  Its data is
// unmarshaled from JSON and then exported to a Brewery struct.
type rawBrewery struct {
	ID           int             `json:"brewery_id"`
	Name         string          `json:"brewery_name"`
	Slug         string          `json:"brewery_slug"`
	Logo         responseURL     `json:"brewery_label"`
	Country      string          `json:"country_name"`
	Active       responseBool    `json:"brewery_active"`
	Location     BreweryLocation `json:"location"`
	Contact      BreweryContact  `json:"contact"`
	Type         string          `json:"brewery_type"`
	TypeID       int             `json:"brewery_type_id"`
	Independent  responseBool    `json:"is_independent"`
	InProduction int             `json:"brewery_in_production"`
	Rating       BreweryRating   `json:"rating"`
	Description  string          `json:"brewery_description"`
	Stats        BreweryStats    `json:"stats"`
}

// export creates an exported Brewery from a rawBrewery struct, allowing for
// more useful structures to be created for client consumption.
func (r *rawBrewery) export() *Brewery {
	return &Brewery{
		ID:           r.ID,
		Name:         r.Name,
		Slug:         r.Slug,
		Logo:         url.URL(r.Logo),
		Country:      r.Country,
		Active:       bool(r.Active),
		Location:     r.Location,
		Contact:      r.Contact,
		Type:         r.Type,
		TypeID:       r.TypeID,
		Independent:  bool(r.Independent),
		InProduction: r.InProduction,
		Rating:       r.Rating,
		Description:  r.Description,
		Stats:        r.Stats,
	}
}
