package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Uttkarsh-raj/steganography/ui"
)

func main() {
	// init new fyne app
	newApp := app.New()

	// Create a new window for the app
	window := newApp.NewWindow("Steg-GO")
	window.Resize(fyne.NewSize(800, 600))

	// Create the ui
	ui.Create(window)

	// Run the Project
	window.ShowAndRun()
}
