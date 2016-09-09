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
	result_filename *string = flag.String("rf", "default_result.kml", "Reusult KML file name ")
	pictures_folder *string = flag.String("pf", "./samples", "Pictures folder name")
)

func FileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	flag.Parse()

	if !FileExist(*pictures_folder) {
		fmt.Println("Please check path of pictures : ", pictures_folder)
		return
	}

	result_fn := *result_filename
	if !strings.Contains(result_fn, ".kml") {
		result_fn += ".kml"
	}

	if err := pic2kml.MakeKml(*pictures_folder, result_fn); err != nil {
		fmt.Printf("Failed makekml, Error : %s", err)
	}

	fmt.Println("Picture Folder : ", *pictures_folder)
	fmt.Println("Result File  : ", *result_filename)
}