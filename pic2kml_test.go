package pic2kml

import (
	"fmt"
	"testing"

	kml "github.com/twpayne/go-kml"
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
	exif, err := p.GetExif("./samples/sample1.jpg")
	if err != nil {
		t.Errorf("GetExif() == %#v, want nil", err)
		return
	}

	t.Log(fmt.Sprintf("GetExif() is Success, exif is %+v", exif))
}

func TestGetExif_WithKey(t *testing.T) {
	p := new(Pic2Kml)
	p.SetApiKey(key)
	exif, err := p.GetExif("./samples/sample1.jpg")
	if err != nil {
		t.Errorf("GetExif() == %#v, want nil", err)
		return
	}

	t.Log(fmt.Sprintf("GetExif() is Success, exif is %+v", exif))
}

func TestMakePoint(t *testing.T) {
	p := new(Pic2Kml)
	exif, err := p.GetExif("./samples/sample1.jpg")
	if err != nil {
		t.Errorf("GetExif() == %#v, want nil", err)
		return
	}
	exif.Number = 1

	point, err := p.MakePoint(exif)
	if err != nil {
		t.Errorf("MakePoint() == %#v,Error is %#v", point, err)
	}
}

func TestMakePoint_Withkey(t *testing.T) {
	p := new(Pic2Kml)
	p.SetApiKey(key)

	exif, err := p.GetExif("./samples/sample1.jpg")
	if err != nil {
		t.Errorf("GetExif() == %#v, want nil", err)
		return
	}
	exif.Number = 1

	point, err := p.MakePoint(exif)
	if err != nil {
		t.Errorf("MakePoint() == %#v,Error is %#v", point, err)
	}
}

func TestMakeLine(t *testing.T) {
	p := new(Pic2Kml)
	start, err := p.GetExif("./samples/sample1.jpg")
	if err != nil {
		t.Errorf("GetExif() == %#v, want nil", err)
		return
	}
	end, err := p.GetExif("./samples/sample2.jpg")
	if err != nil {
		t.Errorf("GetExif() == %#v, want nil", err)
		return
	}

	start.Number = 1
	end.Number = 2

	line, err := p.MakeLine(start, end)
	if err != nil {
		t.Errorf("MakePoint() == %#v,Error is %#v", line, err)
	}
}

func TestMakeExifs(t *testing.T) {
	p := new(Pic2Kml)
	err := p.MakeExifs("./samples")
	if err != nil {
		t.Errorf("MakeExifs(), Error is %#v", err)
	}
}

func TestMakePoints(t *testing.T) {
	p := new(Pic2Kml)
	p.MakeExifs("./samples")
	folder := kml.Folder()
	err := p.MakePoints(folder)
	if err != nil {
		t.Errorf("MakePoints(),Error is %#v", err)
	}
	t.Log(folder)
}

func TestMakeLines(t *testing.T) {
	p := new(Pic2Kml)
	p.MakeExifs("./samples")
	folder := kml.Folder()
	err := p.MakeLines(folder)
	if err != nil {
		t.Errorf("MakeLines(),Error is %#v", err)
	}
	t.Log(folder)
}

func TestMakeKml(t *testing.T) {
	p := new(Pic2Kml)
	err := p.MakeExifs("./samples")
	if err != nil {
		t.Errorf("MakeKml(),Error is %#v", err)
	}

	folder := kml.Folder()
	p.MakePoints(folder)
	p.MakeLines(folder)
	p.MakeKml(folder, "result.kml")

}
