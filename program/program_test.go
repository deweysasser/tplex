package program

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestOptions_Run(t *testing.T) {
	options := Options{
		Templates: []string{"testdata/templates/test.txt"},
		Data:      []string{"testdata/test0.yaml"},
		OutputDir: "testoutput/test0",
	}
	expected := []byte("[one two three]")

	options.Run()

	b, err := ioutil.ReadFile("testoutput/test0/test.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result := bytes.Compare(expected, b); result != 0 {
		t.Errorf("output incorrect.\nExpected: \"%s\"\nResult: \"%s\"\n", expected, b)
	}
}
