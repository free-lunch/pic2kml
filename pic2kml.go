package pic2kml

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
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
	Number int
	Lat    float64
	Lon    float64
	Time   string
	Date   string
	Addr   string
}

type Exifs []*Exif

func (exifs Exifs) Len() int {
	return len(exifs)
}

func (exifs Exifs) Less(i, j int) bool {
	if exifs[i].Date == exifs[j].Date {
		return exifs[i].Time < exifs[j].Time
	} else {
		return exifs[i].Date < exifs[j].Date
	}
}

func (exifs Exifs) Swap(i, j int) {
	exifs[i], exifs[j] = exifs[j], exifs[i]
}

type Pic2Kml struct {
	apikey string
	exifs  Exifs
	root   *kml.CompoundElement
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
		return nil, err
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

func (p *Pic2Kml) MakePoint(exif *Exif) (*kml.CompoundElement, error) {
	if *exif == (Exif{}) {
		return nil, errors.New("Not exist EXIF")
	}

	k := kml.Placemark(
		kml.Name(fmt.Sprintf("#%d", exif.Number)),
		kml.Description(exif.Date+"\n"+exif.Time+"\n"+exif.Addr),
		kml.Point(
			kml.Coordinates(kml.Coordinate{Lat: exif.Lat, Lon: exif.Lon}),
		),
	)

	return k, nil
}

func (p *Pic2Kml) MakeLine(start *Exif, end *Exif) (*kml.CompoundElement, error) {
	if *start == (Exif{}) || *end == (Exif{}) {
		return nil, errors.New("Not exist EXIF")
	}

	k := kml.Placemark(
		kml.Name(fmt.Sprintf("#%d --> #%d", start.Number, end.Number)),
		kml.LineString(
			kml.Coordinates(
				kml.Coordinate{Lat: start.Lat, Lon: start.Lon},
				kml.Coordinate{Lat: end.Lat, Lon: end.Lon},
			),
		),
	)

	return k, nil
}

func (p *Pic2Kml) SortExifs() {
	sort.Sort(p.exifs)
}

func (p *Pic2Kml) MakeExifs(folder string) error {
	f := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		// Check a  type of picture
		ext := strings.ToLower(filepath.Ext(path))
		switch ext {
		case ".jpg":
		case ".bmp":
		case ".gif":
		case ".png":
		default:
			return nil
		}

		exif, err := p.GetExif(path)

		//Case : Not exist GPS data
		if exif == nil {
			return nil
		}
		if err != nil {
			return err
		}

		p.exifs = append(p.exifs, exif)
		return nil
	}

	err := filepath.Walk(folder, f)
	if err != nil {
		return err
	}
	p.SortExifs()

	for i, exif := range p.exifs {
		exif.Number = i + 1
	}
	return nil
}

func (p *Pic2Kml) MakePoints(k *kml.CompoundElement) error {
	if k == nil {
		return errors.New("Invaild kml.CompoundElement")
	}

	for i := 0; i < p.exifs.Len(); i++ {
		point, _ := p.MakePoint(p.exifs[i])
		k.Add(point)
	}

	return nil
}

func (p *Pic2Kml) MakeLines(k *kml.CompoundElement) error {
	if k == nil {
		return errors.New("Invaild kml.CompoundElement")
	}

	for i := 0; i < p.exifs.Len()-1; i++ {
		line, _ := p.MakeLine(p.exifs[i], p.exifs[i+1])
		k.Add(line)
	}
	return nil
}

func (p *Pic2Kml) MakeKml(k *kml.CompoundElement, file_name string) error {
	if p.root == nil {
		p.root = kml.KML()
	}

	if k == nil {
		return errors.New("Invaild kml.CompoundElement")
	}

	p.root.Add(k)

	f, err := os.Create(file_name)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := p.root.WriteIndent(f, "", "  "); err != nil {
		log.Fatal(err)
	}
	return nil
}
