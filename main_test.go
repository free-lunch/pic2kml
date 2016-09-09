package main

import (
	"testing"
	"os"
)

func FileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func TestDefault(t *testing.T){
	os.Args[1] = "./samples"
	main()
	if !FileExist("./default_result.kml") {
		t.Errorf("Not exist Default result file ")
	}
}

func TestResultFileName(t *testing.T){
	os.Args[1] = "./samples"
	*result_filename = "TestResult.kml"
	main()
	FileExist("./TestResult.kml")
	if !FileExist("./TestResult.kml") {
		t.Errorf("Not exist TestResult.kml")
	}
}