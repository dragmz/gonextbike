package nb

import (
	"encoding/xml"
	"strings"
)

type Numbers []string

func (n *Numbers) UnmarshalXMLAttr(attr xml.Attr) error {
	*n = strings.Split(attr.Value, ",")
	return nil
}

type Place struct {
	Uid          int     `xml:"uid,attr"`
	Name         string  `xml:"name,attr"`
	Latitude     string  `xml:"lat,attr"`
	Longitude    string  `xml:"lng,attr"`
	Spot         string  `xml:"spot,attr"`
	Number       string  `xml:"number,attr"`
	Bikes        string  `xml:"bikes,attr"`
	Racks        string  `xml:"bike_racks,attr"`
	TerminalType string  `xml:"terminal_type,attr"`
	Numbers      Numbers `xml:"bike_numbers,attr"`
}

type City struct {
	Uid       string  `xml:"uid,attr"`
	Latitude  string  `xml:"lat,attr"`
	Longitude string  `xml:"lng,attr"`
	Zoom      string  `xml:"zoom,attr"`
	MapsIcon  string  `xml:"maps_icon,attr"`
	Alias     string  `xml:"alias,attr"`
	Break     string  `xml:"break,attr"`
	Name      string  `xml:"name,attr"`
	Places    []Place `xml:"place"`
}

type Country struct {
	Latitude    string `xml:"lat,attr"`
	Longitude   string `xml:"lng,attr"`
	Zoom        string `xml:"zoom,attr"`
	Name        string `xml:"name,attr"`
	Hotline     string `xml:"hotline,attr"`
	Domain      string `xml:"domain,attr"`
	Country     string `xml:"country,attr"`
	CountryName string `xml:"country_name,attr"`
	Terms       string `xml:"terms,attr"`
	Website     string `xml:"website,attr"`
	Cities      []City `xml:"city"`
}

type Markers struct {
	Countries []Country `xml:"country"`
}
