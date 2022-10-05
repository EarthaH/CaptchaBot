package reader

import (
	"testing"
)

var filename = "sample0.png"
var filename2 = "sample1.jpg"

func TestReadHelloWorldImage(t *testing.T) {
	expectedText := "\"Hello World\""

	result, err := ReadImage(filename)

	if err != nil || result != expectedText {
		t.Failed()
	}
}

func TestReadURLImage(t *testing.T) {
	expectedText := "https://discord.gg/makizushi"

	result, err := ReadImage(filename2)

	if err != nil || result != expectedText {
		t.Failed()
	}
}

func TestErrorFilename(t *testing.T) {
	expectedText := ""

	result, err := ReadImage("")

	if err == nil || result != expectedText {
		t.Failed()
	}
}
