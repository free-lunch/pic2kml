package pic2kml

import (
	"fmt"
	"testing"
)

const key string = "AIzaSyCGrUW1AhRTaYKW4x9sD0AnQg3nQzRYGQQ"

func TestSetApiKey(t *testing.T) {
	p := new(Pic2Kml)
	p.SetApiKey(key)
}

func TestGetAddress(t *testing.T) {
	result := "Jl. Samudra, Kuta, Kabupaten Badung, Bali 80361, Indonesia"
	lat, lon := -8.733203944444444, 115.16377158333334

	p := new(Pic2Kml)
	p.SetApiKey(key)
	addr, err := p.GetAddress(lat, lon)
	if err != nil {
		t.Errorf("GetAddress() == %#v, want nil", err)
		return
	}

	if addr != result {
		t.Errorf("Address is wrong %#v,\n Right Address : %#v", err, result)
		return
	}

	t.Logf("GetAddress() is Success, Address : %#v", addr)
}

func TestGetExif(t *testing.T) {
	p := new(Pic2Kml)
	exif, err := p.GetExif("./samples/sample.jpg")
	if err != nil {
		t.Errorf("GetExif() == %#v, want nil", err)
		return
	}

	t.Log(fmt.Sprintf("GetExif() is Success, exif is %+v", exif))
}

func TestGetExif_WithKey(t *testing.T) {
	p := new(Pic2Kml)
	p.SetApiKey(key)
	exif, err := p.GetExif("./samples/sample.jpg")
	if err != nil {
		t.Errorf("GetExif() == %#v, want nil", err)
		return
	}

	t.Log(fmt.Sprintf("GetExif() is Success, exif is %+v", exif))
}

func TestMakeKml(t *testing.T) {
	p := new(Pic2Kml)
	p.MakeKml("result.kml")
}
