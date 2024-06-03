package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Uttkarsh-raj/Steg-GO/ui"
)

func main() {
	// init new fyne app
	newApp := app.New()

	// Create a new window for the app
	window := newApp.NewWindow("Steg-GO")
	icon, err := fyne.LoadResourceFromPath("assets\\Logo.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	window.SetIcon(icon)
	window.Resize(fyne.NewSize(800, 600))

	// Create the ui
	ui.Create(window)

	// Run the Project
	window.ShowAndRun()
}
