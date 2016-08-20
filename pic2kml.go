package pic2kml

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	kml "github.com/twpayne/go-kml"
)

type Pic2Kml struct {
	apikey string
}

type Result struct {
	Address string `json:"formatted_address"`
}

type GeoAPIResponse struct {
	Results []Result `json:"results"`
	Status  string   `json:"status"`
}

type Exif struct {
	Lat  float64
	Lon  float64
	Time string
	Date string
	Addr string
}

func (p *Pic2Kml) SetApiKey(key string) {
	p.apikey = key
}

func (p *Pic2Kml) GetExif(fn string) (*Exif, error) {
	var s = new(Exif)
	f, err := os.Open(fn)
	defer f.Close()

	if err != nil {
		return s, err
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		return s, err
	}

	lat, lon, err := x.LatLong()
	if err != nil {
		return s, err
	}

	t, _ := x.DateTime()
	timeArray := strings.Split(t.String(), " ")

	s.Lat = lat
	s.Lon = lon
	s.Date = timeArray[0]
	s.Time = timeArray[1]

	if p.apikey != "" {
		s.Addr, _ = p.GetAddress(lat, lon)
	}
	return s, nil
}

func (p *Pic2Kml) GetResults(body []byte) (*GeoAPIResponse, error) {
	var s = new(GeoAPIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (p *Pic2Kml) GetAddress(lat, lon float64) (string, error) {
	if p.apikey == "" {
		return "", errors.New("Not exist Google API key")
	}

	base_url := "https://maps.googleapis.com/maps/api/geocode/json"
	query := fmt.Sprintf("?latlng=%f,%f&key=%s", lat, lon, p.apikey)
	resp, err := http.Get(base_url + query)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s, err := p.GetResults([]byte(body))
	if err != nil {
		return "", err
	}
	return s.Results[0].Address, nil
}

func (p *Pic2Kml) MakePoint(n int, exif *Exif) (*kml.CompoundElement, error) {
	// TODO: Make markpoins for kml
	if *exif == (Exif{}) {
		return nil, errors.New("Not exist EXIF")
	}

	k := kml.Placemark(
		kml.Name("#"+string(n)), // Time
		kml.Description(exif.Date+"\n"+exif.Time+"\n"+exif.Addr),
		kml.Point(
			kml.Coordinates(kml.Coordinate{Lat: exif.Lat, Lon: exif.Lon}),
		),
	)

	return k, nil
}

func (p *Pic2Kml) MakeLine(exifs []Exif) (*kml.CompoundElement, error) {
	// TODO: Make linestrings for kml
	return nil, nil
}

func (p *Pic2Kml) MakeKml(file_name string) {

	// Now, This is test code for making kml
	// TODO : Call MakePoints and MakeLines and then make a kml file

	k := kml.KML()
	folder := kml.Folder(

		kml.Placemark(
			kml.Name("#Point1"), // Time
			kml.Description("#Desc1"),
			kml.Point(
				kml.Coordinates(kml.Coordinate{Lat: -8.73320, Lon: 115.16377}),
			),
		),
		kml.Placemark(
			kml.Name("#Points2"), // Time
			kml.Description("#Desc2"),
			kml.Point(
				kml.Coordinates(kml.Coordinate{Lat: -8.73132, Lon: 115.16414}),
			),
		),
		kml.Placemark(
			kml.Name("#Line1-2"), // Time
			kml.Description("#Desc3"),
			kml.LineString(
				kml.Coordinates(
					kml.Coordinate{Lat: -8.73320, Lon: 115.16377},
					kml.Coordinate{Lat: -8.73132, Lon: 115.16414},
				),
			),
		),
	)
	k.Add(folder)

	f, err := os.Create(file_name)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}
	if err := k.WriteIndent(f, "", "  "); err != nil {
		log.Fatal(err)
	}
}
