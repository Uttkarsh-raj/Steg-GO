package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Uttkarsh-raj/steganography/helper"
)

func Create(window fyne.Window) {

	hiddenText := ""

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
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Wrapping = fyne.TextWrapWord

	// For Text Layout
	borderText := canvas.NewRectangle(borderColor)
	borderText.Resize(fyne.NewSize(400, 400))

	TextContainer := container.NewStack(borderText, input)

	// Buttons
	addImage := widget.NewButtonWithIcon("Add Image", theme.FileImageIcon(), func() {
		helper.ShowFilePicker(window, imageLabel, image)
	})

	hideText := widget.NewButtonWithIcon("Hide Text", theme.VisibilityOffIcon(), func() {
		helper.HideTextInImage(window, image, input, &hiddenText)
	})

	extractText := widget.NewButtonWithIcon("Extract Text", theme.VisibilityIcon(), func() {
		helper.ExtractTextFromImage(window, image, input)
	})

	downloadImage := widget.NewButtonWithIcon("Download Image", theme.DownloadIcon(), func() {
		helper.DownloadImage(window, image, hiddenText)
	})

	resetAll := widget.NewButtonWithIcon("Reset", theme.ViewRefreshIcon(), func() {
		helper.ResetAll(window, image, input)
	})

	// Grid Layouts
	buttons := container.NewGridWithRows(4, addImage, hideText, extractText, downloadImage)
	body := container.NewGridWithColumns(2, imageContainer, TextContainer)
	header := container.NewHBox(
		container.New(layout.NewGridWrapLayout(fyne.NewSize(130, 40)), canvas.NewImageFromFile("assets\\title.png")),
		layout.NewSpacer(),
		resetAll,
	)

	// Final Layout (UI)
	window.SetContent(
		container.NewBorder(
			header,
			buttons,
			nil,
			nil,
			body,
		),
	)
}
