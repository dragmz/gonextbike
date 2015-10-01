package nb

import (
	"os"
	"path"
	"testing"
)

func TestGetMarkers(t *testing.T) {
	f, err := os.Open(path.Join("testdata", "nb1.xml"))
	if err != nil {
		t.Fatal(err)
	}

	m, err := parseMarkers(f)
	if err != nil {
		t.Error(err)
	}

	if cc := len(m.Countries); cc != 47 {
		t.Error("Expected 47 countries but got:", cc)
	}

	c := m.Countries[6]

	if c.Latitude != "52.2413" {
		t.Error("Unexpected lat:", c.Latitude)
	}

	if c.Longitude != "18.479" {
		t.Error("Unexpected lng:", c.Longitude)
	}

	if c.Zoom != "6" {
		t.Error("Unexpected zoom:", c.Zoom)
	}

	if c.Name != "WRM nextbike Poland" {
		t.Error("Unexpected name:", c.Name)
	}

	if c.Hotline != "+48717381111" {
		t.Error("Unexpected hotline:", c.Hotline)
	}

	if c.Domain != "pl" {
		t.Error("Unexpected domain:", c.Domain)
	}

	if c.Country != "PL" {
		t.Error("Unexpected country:", c.Country)
	}

	if c.CountryName != "Poland" {
		t.Error("Unexpected country name:", c.CountryName)
	}

	if c.Terms != "http://wroclawskirower.pl/regulamin/" {
		t.Error("Unexpected terms:", c.Terms)
	}

	if c.Website != "http://www.wroclawskirower.pl/" {
		t.Error("Unexpected website:", c.Website)
	}

	if len(c.Cities) != 1 {
		t.Error("Unexpected cities:", len(c.Cities))
	}

	ct := c.Cities[0]
	if ct.Uid != "148" {
		t.Error("Unexpected uid:", ct.Uid)
	}

	if ct.Longitude != "17.0485" {
		t.Error("Unexpected lng:", ct.Longitude)
	}

	if ct.Latitude != "51.1097" {
		t.Error("Unexpected lat:", ct.Latitude)
	}

	if ct.Zoom != "13" {
		t.Error("Unexpected zoom:", ct.Zoom)
	}

	if ct.MapsIcon != "" {
		t.Error("Unexpected maps_icon:", ct.MapsIcon)
	}

	if ct.Alias != "" {
		t.Error("Unexpected alias:", ct.Alias)
	}

	if ct.Break != "0" {
		t.Error("Unexpected break:", ct.Break)
	}

	if ct.Name != "Wrocław" {
		t.Error("Unexpected name:", ct.Name)
	}

	if len(ct.Places) != 73 {
		t.Error("Unexpected places:", len(ct.Places))
	}

	p := ct.Places[0]
	if p.Uid != 349478 {
		t.Error("Unexpected uid:", p.Uid)
	}

	if p.Latitude != "51.13207714749649" {
		t.Error("Unexpected lat:", p.Latitude)
	}

	if p.Longitude != "17.06550121307373" {
		t.Error("Unexpected lng:", p.Longitude)
	}

	if p.Name != "Pl. Kromera" {
		t.Error("Unexpected name:", p.Name)
	}

	if p.Spot != "1" {
		t.Error("Unexpected spot:", p.Spot)
	}

	if p.Number != "5901" {
		t.Error("Unexpected number:", p.Number)
	}

	if p.Bikes != "5+" {
		t.Error("Unexpected bikes:", p.Bikes)
	}

	if p.Racks != "16" {
		t.Error("Unexpected racks:", p.Racks)
	}

	if p.TerminalType != "unknown" {
		t.Error("Unexpected terminal_type:", p.TerminalType)
	}

	if len(p.Numbers) != 5 {
		t.Error("Unexpected bike numbers:", len(p.Numbers))
	}

	expected := []string{"57710", "57266", "57082", "57655", "57697"}
	for i, num := range p.Numbers {
		if num != expected[i] {
			t.Error("Unexpected number:", num)
		}
	}
}
