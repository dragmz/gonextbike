package nb

import (
	"encoding/xml"
	"io"
	"net/http"
)

func parseMarkers(r io.Reader) (Markers, error) {
	state := Markers{}

	d := xml.NewDecoder(r)
	if err := d.Decode(&state); err != nil {
		return state, err
	}

	return state, nil
}

func GetMarkers() (Markers, error) {
	c := http.Client{}
	url := "http://nextbike.net/maps/nextbike-official.xml"

	resp, err := c.Get(url)
	if err != nil {
		return Markers{}, err
	}

	defer resp.Body.Close()

	return parseMarkers(resp.Body)
}
