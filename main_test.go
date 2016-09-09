package main

import (
	"testing"
	"os"
)


func TestDefault(t *testing.T){
	os.Args = []string{"__", "-pf=./samples"}
	main()
	if !FileExist("./default_result.kml") {
		t.Errorf("Not exist Default result file ")
	}
}

func TestResultFileName(t *testing.T){
	os.Args = []string{"__", "-pf=./samples", "-rf=./TestResult.kml"}
	main()
	FileExist("./TestResult.kml")
	if !FileExist("./TestResult.kml") {
		t.Errorf("Not exist TestResult.kml")
	}
}