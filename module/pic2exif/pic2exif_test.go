package pic2exif

import (
	"testing"
	"fmt"
)

// TODO : Modify to properly test cases
//func TestGetAddress(t *testing.T) {
//	result := "Jl. Samudra, Kuta, Kabupaten Badung, Bali 80361, Indonesia"
//	lat, lon := -8.733203944444444, 115.16377158333334
//
//	p := new(Pic2Kml)
//	p.SetApiKey(key)
//	addr, err := p.GetAddress(lat, lon)
//	if err != nil {
//		t.Errorf("GetAddress() == %#v, want nil", err)
//		return
//	}
//
//	if addr != result {
//		t.Errorf("Address is wrong %#v,\n Right Address : %#v", err, result)
//		return
//	}
//
//	t.Logf("GetAddress() is Success, Address : %#v", addr)
//}
//
//func TestGetExif(t *testing.T) {
//	p := new(Pic2Kml)
//	exif, err := p.GetExif(sampleFolder + "sample1.jpg")
//	if err != nil {
//		t.Errorf("GetExif() == %#v, want nil", err)
//		return
//	}
//
//	t.Log(fmt.Sprintf("GetExif() is Success, exif is %+v", exif))
//}
//
//func TestGetExif_WithKey(t *testing.T) {
//	p := new(Pic2Kml)
//	p.SetApiKey(key)
//	exif, err := p.GetExif("./Samples/sample1.jpg")
//	if err != nil {
//		t.Errorf("GetExif() == %#v, want nil", err)
//		return
//	}
//
//	t.Log(fmt.Sprintf("GetExif() is Success, exif is %+v", exif))
//}
//func TestMakeExifs(t *testing.T) {
//	p := new(Pic2Kml)
//	err := p.MakeExifs(sampleFolder)
//	if err != nil {
//		t.Errorf("MakeExifs(), Error is %#v", err)
//	}
//}