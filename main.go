package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	// "fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "image/color"
	"context"
	"errors"
	"github.com/rocketlaunchr/google-search"
	"strings"
)

func main() {
	a := app.New()
	win := a.NewWindow("NGIC")

	zip := widget.NewEntry()
	zip.SetPlaceHolder("Enter Zip")

	button := widget.NewButton("Get Zip", func() {
		tax, err := searchForTaxRate(zip.Text)
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			zip.SetText(tax)
		}
	})

	//sets a system tray that will only close from the tray/toolbar if using a desktop
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("NGICapp",
			fyne.NewMenuItem("Show", func() {
				win.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	menu := container.NewGridWithRows(3,
		container.NewGridWithColumns(3,
			// layout.NewSpacer(),
			container.NewVBox(
				zip,
				button,
				layout.NewSpacer(),
			),
		),
	)

	win.SetContent(widget.NewLabel("Fyne System Tray"))

	// win.SetContent(hBox)
	win.SetContent(menu)

	//when using the X close button, the app is only closed and minimuzed to the tray
	win.SetCloseIntercept(func() {
		win.Hide()
	})
	win.ShowAndRun()
}

func searchForTaxRate(zipCode string) (string, error) {

	search, err := googlesearch.Search(context.TODO(), "tax rate of "+zipCode)
	if err != nil {
		return "", err
	}

	for _, v := range search {
		if strings.Contains(v.Description, "%") {
			pos := strings.Index(v.Description, "%")
			if pos == -1 {
				break
			}
			tax := v.Description[(pos - 6) : pos+1]
			tax = strings.TrimPrefix(tax, "is ")
			tax = strings.TrimPrefix(tax, " is ")
			tax = strings.TrimPrefix(tax, "is")
			tax = strings.TrimPrefix(tax, " ")
			tax = strings.TrimPrefix(tax, "s")
			tax = strings.TrimPrefix(tax, "a")
			return tax, nil
		}
	}

	return "", errors.New("could not locate tax")
}
