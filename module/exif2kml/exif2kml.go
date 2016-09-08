package exif2kml

import (
	"github.com/twpayne/go-kml"
	"fmt"
	"github.com/free-lunch/pic2kml/module/pic2exif"
	"errors"

	"os"
	"log"
)

type Exif2Kml struct {
	exifGroup *pic2exif.ExifGroup
	days   []*kml.CompoundElement
}

type Option struct {
	name string
	value string
}

func NewExif2Kml() *Exif2Kml {
	e := &Exif2Kml{
	}
	return e
}

func (e *Exif2Kml) SetExifGroup(exifGroup *pic2exif.ExifGroup) {
	e.exifGroup = exifGroup
}

func (e *Exif2Kml) MakePoints(folder *kml.CompoundElement) error {
	if folder == nil {
		return errors.New("Foloder is empty")
	}

	for _, v := range *e.exifGroup {
		placemark := kml.Placemark(
			kml.Point(
				kml.CoordinatesArray([]float64{v.Lon, v.Lat}),
			),
		)
		folder.Add(placemark)
	}

	return nil
}

func (e *Exif2Kml) MakeLines(folder *kml.CompoundElement) error {
	if folder == nil {
		return errors.New("Foloder is empty")
	}

	for i := 0; i < e.exifGroup.Len()-1; i++ {
		v1 := (*e.exifGroup)[i]
		v2 := (*e.exifGroup)[i+1]
		placemark := kml.Placemark(
			kml.LineString(
				kml.Coordinates(
					kml.Coordinate{Lat: v1.Lat, Lon: v1.Lon},
					kml.Coordinate{Lat: v2.Lat, Lon: v2.Lon},
				),

			),
		)
		folder.Add(placemark)
	}

	return nil
}

func (e *Exif2Kml) MakeKml(file_name string, options ...interface{}) error {
	if len(options) > 0 {
		fmt.Println(options)
		//if daysplit {
		//	// TODO : Split exifGroup and then handle its by each day of pictures
		//	timeFormat := "2006-01-02"
		//	start, _ := time.Parse(timeFormat, (*e.exifGroup)[0].Date)
		//	for i := 1; i < len(*e.exifGroup); i++ {
		//		now, _ := time.Parse(timeFormat, (*e.exifGroup)[i].Date)
		//		n := int(now.Sub(start).Hours() / 24)
		//		fmt.Println(n)
		//
		//	}
		//}
	}

	folder := kml.Folder()
	e.MakePoints(folder)
	e.MakeLines(folder)

	d:= kml.Document(kml.Name("Total Points and Paths"),folder)
	k := kml.KML(d)


	// Make KML file
	f, err := os.Create(file_name)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if  err := k.WriteIndent(f, "", "  "); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}