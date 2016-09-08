package exif2kml

import (
	"testing"
	"github.com/twpayne/go-kml"
)

// TODO : Modify test cases properly
//
//func TestMakePoint(t *testing.T) {
//	p := new(Pic2Kml)
//	exif, err := p.GetExif(sampleFolder + "sample1.jpg")
//	if err != nil {
//		t.Errorf("GetExif() == %#v, want nil", err)
//		return
//	}
//	exif.Number = 1
//
//	point, err := p.MakePoint(exif)
//	if err != nil {
//		t.Errorf("MakePoint() == %#v,Error is %#v", point, err)
//	}
//}
//
//func TestMakePoint_Withkey(t *testing.T) {
//	p := new(Pic2Kml)
//	p.SetApiKey(key)
//
//	exif, err := p.GetExif(sampleFolder + "sample1.jpg")
//	if err != nil {
//		t.Errorf("GetExif() == %#v, want nil", err)
//		return
//	}
//	exif.Number = 1
//
//	point, err := p.MakePoint(exif)
//	if err != nil {
//		t.Errorf("MakePoint() == %#v,Error is %#v", point, err)
//	}
//}
//
//func TestMakeLine(t *testing.T) {
//	p := new(Pic2Kml)
//	start, err := p.GetExif(sampleFolder + "sample1.jpg")
//	if err != nil {
//		t.Errorf("GetExif() == %#v, want nil", err)
//		return
//	}
//	end, err := p.GetExif(sampleFolder + "sample2.jpg")
//	if err != nil {
//		t.Errorf("GetExif() == %#v, want nil", err)
//		return
//	}
//
//	start.Number = 1
//	end.Number = 2
//
//	line, err := p.MakeLine(start, end)
//	if err != nil {
//		t.Errorf("MakePoint() == %#v,Error is %#v", line, err)
//	}
//}
//func TestMakePoints(t *testing.T) {
//	p := new(Pic2Kml)
//	p.MakeExifs(sampleFolder)
//	folder := kml.Folder()
//	err := p.MakePoints(folder)
//	if err != nil {
//		t.Errorf("MakePoints(),Error is %#v", err)
//	}
//	t.Log(folder)
//}
//
//func TestMakeLines(t *testing.T) {
//	p := new(Pic2Kml)
//	p.MakeExifs(sampleFolder)
//	folder := kml.Folder()
//	err := p.MakeLines(folder)
//	if err != nil {
//		t.Errorf("MakeLines(),Error is %#v", err)
//	}
//	t.Log(folder)
//}
//
//func TestMakeKml(t *testing.T) {
//	p := new(Pic2Kml)
//	err := p.MakeExifs(sampleFolder)
//	if err != nil {
//		t.Errorf("MakeKml(),Error is %#v", err)
//	}
//
//	folder := kml.Folder()
//	p.MakePoints(folder)
//	p.MakeLines(folder)
//	p.MakeKml(folder, )
//
//}