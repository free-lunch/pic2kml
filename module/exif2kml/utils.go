package exif2kml

import (
	"time"
	"math/rand"
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
	"math"
)

func GetDiffDay(d1 string, d2 string) int {
	timeFormat := "2006-01-02"
	start, _ := time.Parse(timeFormat, d1)
	now, _ := time.Parse(timeFormat, d2)
	ret := int(now.Sub(start).Hours()/24)*1.0
	if ret < 0 {
		ret = -ret
	}
	return ret
}

func GetColorPallete(n int)[]color.Color{
	golden_ratio_conjugate := 0.618033988749895*360
	h := rand.Float64()*360
	ret := make([]color.Color, n)
	for i := 0; i < n; i++ {
		ret[i] = colorful.Hsv(h,0.8,0.8)
		h += golden_ratio_conjugate
		h = math.Mod(h, 360.0)
	}
	return ret
}