package helper

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
		image.Hidden = false
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
	image.File = ""
	image.Resource = nil
	image.Hidden = true
	image.Refresh()
	text.Text = ""
	text.Refresh()
	dialog.ShowInformation("Steg-GO", "Reset Successful !!", window)
}
