package pic2kml

import (
	"encoding/json"
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

var apikey string

func SetApiKey(key string) {
	apikey = key
}

func GetExif(fn string, useAddr bool) (*Exif, error) {
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

	if useAddr && apikey != "" {
		s.Addr, _ = GetAddress(lat, lon, apikey)
	}
	return s, nil
}

func GetResults(body []byte) (*GeoAPIResponse, error) {
	var s = new(GeoAPIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func GetAddress(lat, lon float64, key string) (string, error) {
	base_url := "https://maps.googleapis.com/maps/api/geocode/json"
	query := fmt.Sprintf("?latlng=%f,%f&key=%s", lat, lon, key)
	resp, err := http.Get(base_url + query)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s, err := GetResults([]byte(body))
	if err != nil {
		return "", err
	}
	return s.Results[0].Address, nil
}

func MakePoints(k *kml.CompoundElement, exifs []Exif) error {
	// TODO: Make markpoins for kml
	return nil
}

func MakeLines(k *kml.CompoundElement, exifs []Exif) error {
	// TODO: Make linestrings for kml
	return nil
}

func MakeKml(file_name string) {

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
