package pic2exif

import (
	"os"
	"strings"
	"path/filepath"
	"sort"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

type Pic2Exif  struct {
	//exifGroup *ExifGroup
	apiKey string
}

func NewPic2Exif() *Pic2Exif{
	e := &Pic2Exif{
		apiKey: "",
	}
	return e
}

func (p *Pic2Exif) SetApiKey(key string) {
	p.apiKey = key
}


func (p *Pic2Exif) SortExifGroup(e *ExifGroup) {
	sort.Sort(e)
}

func GetExif(fileName string, apiKey string) (*Exif, error) {
	var s = new(Exif)
	f, err := os.Open(fileName)
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

	if apiKey != "" {
		s.Addr, _ = GetAddress(lat, lon, apiKey)
	}
	return s, nil
}

func (p *Pic2Exif) MakeExifGroupFromFolder(Folder string) (*ExifGroup, error) {
	exifGroup := ExifGroup{}
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
		exif, err := GetExif(path, p.apiKey)
		//Case : Not exist GPS data
		if exif == nil {
			return nil
		}
		if err != nil {
			return err
		}
		exifGroup = append(exifGroup, exif)
		return nil
	}

	err := filepath.Walk(Folder, f)
	if err != nil {
		return nil, err
	}
	sort.Sort(exifGroup)


	for i, exif := range exifGroup {
		exif.Number = i + 1
	}
	return &exifGroup, nil
}