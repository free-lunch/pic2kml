package pic2kml

import (
	"os"
	"github.com/free-lunch/pic2kml/module/pic2exif"
	"fmt"
	"github.com/free-lunch/pic2kml/module/exif2kml"
)

func Folder_Check(path string ) (bool, error){
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err

}

func MakeKml(pictures_folder string, result_name string, options ...interface{} ) error {
	p2e := pic2exif.NewPic2Exif()
	exifGroup, err := p2e.MakeExifGroupFromFolder(pictures_folder)
	if err != nil {
		fmt.Printf("[Error] MakeExif :%s ", err)
		return err
	}

	e2k := exif2kml.NewExif2Kml()
	e2k.SetExifGroup(*exifGroup)
	if err := e2k.MakeKml(result_name); err != nil {
		fmt.Println("[Error] Pic2Kml : MakeKml : ", err)
		return err
	}

	return nil

}

