package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"github.com/brodiep21/ngicGui/search"
	// "fyne.io/fyne/v2/theme"
	// "bytes"
	// "fmt"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "image/png"
	// "io"
	// "os"
)

func main() {
	a := app.New()
	win := a.NewWindow("NGIC")

	zip := widget.NewEntry()
	zip.SetPlaceHolder("Enter Zip")

	//button fills zip placeholder
	button := widget.NewButton("Get Zip", func() {
		tax, err := search.SearchForTaxRate(zip.Text)
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

	//create a canvas object from a file image
	img := canvas.NewImageFromFile("positionstatement.PNG")

	//main menu within same page
	menu := container.NewGridWithRows(1,
		container.NewGridWithColumns(2,
			container.NewVBox(
				widget.NewButton("Scan Position Statement", func() {
					img.FillMode = canvas.ImageFillContain
					win2 := a.NewWindow("Scans")
					win2.SetContent(img)
					win2.Resize(fyne.NewSize(550, 975))
					win2.CenterOnScreen()
					win2.Show()
				}),
				widget.NewButton("2", func() {}),
				widget.NewButton("3", func() {}),
				widget.NewButton("4", func() {}),
				widget.NewButton("5", func() {}),
				widget.NewButton("6", func() {}),
				widget.NewButton("7", func() {}),
				widget.NewButton("8", func() {}),
				widget.NewButton("9", func() {}),
			),
			// container.NewHBox(
			// 	widget.NewButton("Hello", func() {}),
			// 	widget.NewButton("Hello", func() {}),
			// 	widget.NewButton("Hello", func() {}),
			// ),
			container.NewGridWithColumns(3,
				layout.NewSpacer(),
				container.NewVBox(
					zip,
					button,
				),
			),
		),
	)

	// win.SetContent(hBox)
	win.SetContent(menu)

	//when using the X close button, the app is only closed and minimuzed to the tray
	win.SetCloseIntercept(func() {
		win.Hide()
	})
	win.SetMaster()
	win.Show()
	a.Run()
}
