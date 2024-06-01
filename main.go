package main

import (
	"bytes"
	"fmt"
	"image/color"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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

func HideTextInImage(window fyne.Window, image *canvas.Image, text *widget.Entry, hidenText *string) {
	content, err := os.ReadFile(image.File)
	if err != nil {
		dialog.ShowError(err, window)
		return
	}
	if text.Text == "" || text == nil {
		dialog.ShowError(fmt.Errorf("error: Please enter text to be hidden"), window)
		return
	}
	*hidenText = string(content)
	*hidenText += "\n##hiddenText##: " + text.Text
	dialog.ShowInformation("Steg-GO", "Operation completed successfully! Your image is ready for download in JPG or PNG format.", window)
}

func DownloadImage(window fyne.Window, image *canvas.Image, hidenText string) {

	if image == nil || image.File == "" {
		dialog.ShowError(fmt.Errorf("error: No image selected"), window)
		return
	}

	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if writer == nil {
			return
		}
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		defer writer.Close()
		buffer := bytes.NewBufferString(hidenText)
		_, err = writer.Write(buffer.Bytes())
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		dialog.ShowInformation("Steg-GO", "Download completed successfully! Please check the designated location.", window)
	}, window)
}

func ExtractTextFromImage(window fyne.Window, image *canvas.Image, text *widget.Entry) {
	if image.File == "" || image == nil {
		dialog.ShowError(fmt.Errorf("error: No image selected"), window)
		return
	}
	content, err := os.ReadFile(image.File)
	if err != nil {
		dialog.ShowError(err, window)
		return
	}
	info := strings.Split(string(content), "##hiddenText##: ")
	text.Text = info[len(info)-1]
	text.Refresh()
	dialog.ShowInformation("Steg-GO", "Text extracted Successfully !!", window)
}

func ResetAll(window fyne.Window, image *canvas.Image, text *widget.Entry) {
	*image = canvas.Image{}
	text.Text = ""
	image.Refresh()
	text.Refresh()
	dialog.ShowInformation("Steg-GO", "Reset Successful !!", window)
}

func main() {
	newApp := app.New()
	window := newApp.NewWindow("Steg-GO")
	hiddenText := ""

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
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Wrapping = fyne.TextWrapWord

	// For Text Layout
	borderText := canvas.NewRectangle(borderColor)
	borderText.Resize(fyne.NewSize(400, 400))

	TextContainer := container.NewStack(borderText, input)

	// Buttons
	addImage := widget.NewButtonWithIcon("Add Image", theme.FileImageIcon(), func() {
		ShowFilePicker(window, imageLabel, image)
	})

	hideText := widget.NewButtonWithIcon("Hide Text", theme.VisibilityOffIcon(), func() {
		HideTextInImage(window, image, input, &hiddenText)
	})

	extractText := widget.NewButtonWithIcon("Extract Text", theme.VisibilityIcon(), func() {
		ExtractTextFromImage(window, image, input)
	})

	downloadImage := widget.NewButtonWithIcon("Download Image", theme.DownloadIcon(), func() {
		DownloadImage(window, image, hiddenText)
	})

	resetAll := widget.NewButtonWithIcon("Reset", theme.ViewRefreshIcon(), func() {
		ResetAll(window, image, input)
	})

	// Grid Layouts
	buttons := container.NewGridWithRows(4, addImage, hideText, extractText, downloadImage)
	body := container.NewGridWithColumns(2, imageContainer, TextContainer)
	header := container.NewHBox(
		container.New(layout.NewGridWrapLayout(fyne.NewSize(0, 0)), widget.NewLabel("Steg-GO")),
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

	// Run the Project
	window.ShowAndRun()
}
