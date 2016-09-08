package pic2exif

type Exif struct {
	Number int
	Lat    float64
	Lon    float64
	Time   string
	Date   string
	Addr   string
}

type ExifGroup []*Exif

func (e ExifGroup) Len() int {
	return len(e)
}

func (e ExifGroup) Less(i, j int) bool {
	if e[i].Date == e[j].Date {
		return e[i].Time < e[j].Time
	} else {
		return e[i].Date < e[j].Date
	}
}

func (e ExifGroup) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}