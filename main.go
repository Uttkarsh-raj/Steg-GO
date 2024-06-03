package main

import (
	"embed"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Uttkarsh-raj/Steg-GO/helper"
	"github.com/Uttkarsh-raj/Steg-GO/ui"
)

// Embed assets to the project (allows the application to run anywhere and does not depend on the resources)

//go:embed assets/*
var assets embed.FS

func main() {

	// init new fyne app
	newApp := app.New()

	// Initialize resources
	logoImage, err := helper.LoadResourceFromEmbed("assets/Logo.png", &assets)
	if err != nil {
		fmt.Println(err)
	}
	titleImage, err := helper.LoadResourceFromEmbed("assets/title.png", &assets)
	if err != nil {
		fmt.Println(err)
	}

	// Create a new window for the app
	window := newApp.NewWindow("Steg-GO")
	window.SetIcon(logoImage)
	window.Resize(fyne.NewSize(800, 600))

	// Create the ui
	ui.Create(window, titleImage)

	// Run the Project
	window.ShowAndRun()
}
