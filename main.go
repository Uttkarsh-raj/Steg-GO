package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowFilePicker(window fyne.Window, selectedFile *widget.Label, image *canvas.Image) {
	dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
		file := "No files selected!!"
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		if uc == nil {
			return
		}
		file = uc.URI().Path()
		selectedFile.SetText(file)
		image.File = file
		image.Refresh()
	}, window)
}

func HideTextInImage(window fyne.Window, image *canvas.Image, text *widget.Entry) {

}

func main() {
	fmt.Println("Hello Gophers")
	newApp := app.New()
	window := newApp.NewWindow("Ste-Go")

	window.Resize(fyne.NewSize(800, 600))

	// Image View
	imageLabel := widget.NewLabel("NONE")
	image := canvas.NewImageFromFile("")
	image.FillMode = canvas.ImageFillContain

	// For Image Layout
	borderColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	borderImage := canvas.NewRectangle(borderColor)
	borderImage.Resize(fyne.NewSize(400, 400))

	imageContainer := container.NewStack(borderImage, image)

	// Text Entry widget
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	input.MultiLine = true

	// For Text Layout
	borderText := canvas.NewRectangle(borderColor)
	borderText.Resize(fyne.NewSize(400, 400))

	TextContainer := container.NewStack(borderText, input)

	// Buttons
	addImage := widget.NewButton("Add Image", func() {
		ShowFilePicker(window, imageLabel, image)
	})

	hideText := widget.NewButton("Hide Text", func() {
		fmt.Println("Hide Text !!")
	})

	extractText := widget.NewButton("Extract Text", func() {
		fmt.Println("Extract Text !!!")
	})

	downloadImage := widget.NewButton("Download Image", func() {
		fmt.Println("Download Image !!!")
	})

	// Grid Layouts
	buttons := container.NewGridWithRows(4, addImage, hideText, extractText, downloadImage)
	body := container.NewGridWithColumns(2, imageContainer, TextContainer)

	// Final Layout (UI)
	window.SetContent(
		container.NewBorder(
			widget.NewLabel("Ste-Go"),
			buttons,
			nil,
			nil,
			body,
		),
	)

	// Run the Project
	window.ShowAndRun()
}
