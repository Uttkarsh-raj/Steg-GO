package main

import (
	"strings"
	"testing"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/Uttkarsh-raj/steganography/helper"
)

func TestHideTextInImage(t *testing.T) {
	a := app.New()
	w := a.NewWindow("Test")

	image := canvas.NewImageFromFile("assets\\Logo.png")
	text := widget.NewMultiLineEntry()
	text.Text = "Test Data"
	hiddenText := ""

	helper.HideTextInImage(w, image, text, &hiddenText)

	contents := strings.Split(string(hiddenText), "\n##hiddenText##: ")
	expected := contents[len(contents)-1]
	if expected != "Test Data" {
		t.Errorf("expected %s, got %s", expected, hiddenText)
	}
}

func TestExtractTextFromImage(t *testing.T) {
	a := app.New()
	w := a.NewWindow("Test")

	image := canvas.NewImageFromFile("assets\\test_extract.jpg")
	text := widget.NewMultiLineEntry()

	helper.ExtractTextFromImage(w, image, text)

	expected := "Hello there this is a letter to specify that i am in need\nPlease help me....."
	if text.Text != expected {
		t.Errorf("expected %s, got %s", expected, text.Text)
	}
}

func TestResetAll(t *testing.T) {
	a := app.New()
	w := a.NewWindow("Test")

	image := canvas.NewImageFromFile("assets\\Logo.png")
	text := widget.NewMultiLineEntry()

	helper.ResetAll(w, image, text)

	if image.File != "" {
		t.Errorf("expected empty image file, got %s", image.File)
	}
	if text.Text != "" {
		t.Errorf("expected empty text, got %s", text.Text)
	}
}
