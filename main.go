package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/free-lunch/pic2kml/module/pic2kml"
	"strings"
)


var (
	apikey *string = flag.String("apikey", "", "Google API Key for changing from gps location to address")
	daysplit *bool = flag.Bool("daysplit", false, "Support spliting result using days")
	result_filename *string = flag.String("result_filename", "default_result.kml", "Reusult KML file name ")
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please check your command, Must need to PICTURES_FOLDER")
		return
	}

	flag.Parse()
	pictures_folder := os.Args[1]

	result_fn := *result_filename
	if !strings.Contains(result_fn, ".kml") {
		result_fn += ".kml"
	}

	if err := pic2kml.MakeKml(pictures_folder, result_fn); err != nil {
		fmt.Printf("Failed makekml, Error : %s", err)
	}
}