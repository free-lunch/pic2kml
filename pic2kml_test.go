package pic2kml

import (
	"fmt"
	"testing"
)

func TestSetApiKey(t *testing.T) {
	key := "AIzaSyCGrUW1AhRTaYKW4x9sD0AnQg3nQzRYGQQ"
	SetApiKey(key)
}

func TestGetAddress(t *testing.T) {
	// Need to check API

	result := "Jl. Samudra, Kuta, Kabupaten Badung, Bali 80361, Indonesia"
	lat, lon := -8.733203944444444, 115.16377158333334
	key := "AIzaSyCGrUW1AhRTaYKW4x9sD0AnQg3nQzRYGQQ"

	addr, err := GetAddress(lat, lon, key)
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
	exif, err := GetExif("sample.jpg", true)
	if err != nil {
		t.Errorf("GetExif() == %#v, want nil", err)
		return
	}

	t.Log(fmt.Sprintf("GetExif() is Success, exif is %+v", exif))
}

func TestMakeKml(t *testing.T) {
	MakeKml("result.kml")
}
