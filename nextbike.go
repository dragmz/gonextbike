package nb

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

type Numbers []string

func (n *Numbers) UnmarshalXMLAttr(attr xml.Attr) error {
	*n = strings.Split(attr.Value, ",")
	return nil
}

type Bike struct {
	Number        int    `xml:"number,attr"`
	Type          int    `xml:"type,attr"`
	Active        int    `xml:"active,attr"`
	State         string `xml:"state,attr"`
	ElectricLock  bool   `xml:"electric_lock,attr"`
	BoardComputer string `xml:"board_computer,attr"`
}

type Types map[string]int

func (t *Types) UnmarshalXMLAttr(attr xml.Attr) error {
	return json.Unmarshal([]byte(attr.Value), t)
}

type Place struct {
	UID          int     `xml:"uid,attr"`
	Name         string  `xml:"name,attr"`
	Latitude     float32 `xml:"lat,attr"`
	Longitude    float32 `xml:"lng,attr"`
	Spot         int     `xml:"spot,attr"`
	Number       int     `xml:"number,attr"`
	Bikes        int     `xml:"bikes,attr"`
	Racks        int     `xml:"bike_racks,attr"`
	FreeRacks    int     `xml:"free_racks,attr"`
	TerminalType string  `xml:"terminal_type,attr"`
	Numbers      Numbers `xml:"bike_numbers,attr"`
	Types        Types   `xml:"bike_types,attr"`
	PlaceType    string  `xml:"place_type,attr"`
	RackLocks    string  `xml:"rack_locks,attr"`
	BikeArray    []Bike  `xml:"bike"`
}

type Coord struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

type Bounds struct {
	SouthWest Coord `json:"south_west"`
	NorthEast Coord `json:"north_east"`
}

func (b *Bounds) UnmarshalXMLAttr(attr xml.Attr) error {
	return json.Unmarshal([]byte(attr.Value), b)
}

type City struct {
	UID            int     `xml:"uid,attr"`
	Latitude       float32 `xml:"lat,attr"`
	Longitude      float32 `xml:"lng,attr"`
	Zoom           int     `xml:"zoom,attr"`
	MapsIcon       string  `xml:"maps_icon,attr"`
	Alias          string  `xml:"alias,attr"`
	Break          string  `xml:"break,attr"`
	Name           string  `xml:"name,attr"`
	NumPlaces      int     `xml:"num_places,attr"`
	RefreshRate    int     `xml:"refresh_rate,attr"`
	Bounds         Bounds  `xml:"bounds,attr"`
	BookedBikes    int     `xml:"booked_bikes,attr"`
	SetPointBikes  int     `xml:"set_point_bikes,attr"`
	AvailableBikes int     `xml:"available_bikes,attr"`

	Places []Place `xml:"place"`
}

type Country struct {
	Latitude             float32 `xml:"lat,attr"`
	Longitude            float32 `xml:"lng,attr"`
	Zoom                 int     `xml:"zoom,attr"`
	Name                 string  `xml:"name,attr"`
	Hotline              string  `xml:"hotline,attr"`
	Domain               string  `xml:"domain,attr"`
	Country              string  `xml:"country,attr"`
	CountryName          string  `xml:"country_name,attr"`
	Terms                string  `xml:"terms,attr"`
	Policy               string  `xml:"policy,attr"`
	Website              string  `xml:"website,attr"`
	BookedBikes          int     `xml:"booked_bikes,attr"`
	SetPointBikes        int     `xml:"set_point_bikes,attr"`
	AvailableBikes       int     `xml:"available_bikes,attr"`
	CappedAvailableBikes int     `xml:"capped_available_bikes,attr"`
	Cities               []City  `xml:"city"`
}

type Markers struct {
	Countries []Country `xml:"country"`
}
