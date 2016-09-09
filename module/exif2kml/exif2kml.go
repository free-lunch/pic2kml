package exif2kml

import (
	"os"
	"log"

	"github.com/free-lunch/pic2kml/module/pic2exif"
	"github.com/twpayne/go-kml"
	"fmt"
)

type Exif2Kml struct {
	exifGroup pic2exif.ExifGroup
	root *kml.CompoundElement
	numDays   int
	startDay  string
	daySplit bool
	//styles []*kml.SharedElement
}

func NewExif2Kml() *Exif2Kml {
	return  &Exif2Kml{
		root:kml.Document(),
		daySplit:true,
	}
}

func (e *Exif2Kml) SetExifGroup(exifGroup pic2exif.ExifGroup) {
	e.exifGroup = exifGroup
	e.startDay = exifGroup[0].Date
	e.numDays = GetDiffDay(exifGroup[0].Date, exifGroup[e.exifGroup.Len()-1].Date)+1
	//e.styles = make([]*kml.SharedElement, e.numDays)
}

func MakePoints(folder *kml.CompoundElement, exifGroup pic2exif.ExifGroup) {
	for _, v := range exifGroup {
		placemark := kml.Placemark(
			kml.Point(
				kml.CoordinatesArray([]float64{v.Lon, v.Lat}),
			),
		)
		folder.Add(placemark)
	}
}

func MakeLines(folder *kml.CompoundElement, exifGroup pic2exif.ExifGroup) {
	for i := 0; i < exifGroup.Len()-1; i++ {
		v1 := exifGroup[i]
		v2 := exifGroup[i+1]
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
}

// Now, Style was ignored when loading kml in google maps
func MakeStyles(folder *kml.CompoundElement, n int) {
	//element_color := GetColorPallete(n)
	//for i := 0; i < n; i ++{
	//	style := kml.Style(
	//		kml.LineStyle(
	//			kml.Color(element_color[i]),
	//			kml.Scale(1.0),
	//		),
	//		kml.IconStyle(
	//			kml.Color(element_color[i]),
	//			kml.Scale(1.0),
	//		),
	//	)
	//	e.styles[i] = style
	//	folder.Add(style)
	//}
}

func MakeFolder(document *kml.CompoundElement, exifGroup pic2exif.ExifGroup,folder_name string) {
	folder := kml.Folder()
	folder.Add(kml.Name(folder_name))
	MakePoints(folder, exifGroup)
	MakeLines(folder, exifGroup)
	document.Add(folder)
}

func (e *Exif2Kml) exifGroupByDay_Generator() chan pic2exif.ExifGroup {
	channel := make(chan pic2exif.ExifGroup)
	idx := 0
	day := 0
	current_day := e.startDay
	start_idx := idx
	go func() {
		for {
			idx += 1
			if e.exifGroup[idx].Date != current_day {
				current_day = e.exifGroup[idx].Date
				channel <- e.exifGroup[start_idx:idx]
				start_idx = idx
				day++
			}

			if idx == e.exifGroup.Len() - 1 {
				channel <- e.exifGroup[start_idx:]
				break
			}
		}
		close(channel)
	} ()
	return channel

}


func (e *Exif2Kml) MakeKml(file_name string) error {
	if e.daySplit {
		exifGroup_gen := e.exifGroupByDay_Generator()
		day := 1
		for v := range exifGroup_gen {
			MakeFolder(e.root, v, fmt.Sprintf("%d day", day))
			day ++
		}
	} else  {

		MakeFolder(e.root, e.exifGroup, fmt.Sprintf("Travel for %d days", e.numDays))
	}


	// Make KML file
	f, err := os.Create(file_name)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if  err := e.root.WriteIndent(f, "", "  "); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}