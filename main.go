package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/brodiep21/ngicGui/search"
)

func main() {
	a := app.New()

	win := a.NewWindow("NGIC")

	//sets a system tray that will only close from the tray/toolbar if using a desktop
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("NGICapp",
			fyne.NewMenuItem("Show", func() {
				win.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	zip := widget.NewEntry()
	zip.SetPlaceHolder("Enter Zip")

	//button fills zip placeholder
	button := widget.NewButton("Get Zip", func() {
		log.Println("Keyboard enter")

		tax, err := search.SearchForTaxRate(zip.Text)
		if err != nil {
			zip.SetText("Tax not found")
		} else {
			zip.SetText(tax)
		}
	})
	//top toolbar on the main page
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {

		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.FileIcon(), func() {}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {

			log.Println("Display help")
		}),
	)

	//create a canvas object from a file image
	img := canvas.NewImageFromFile("positionstatement.PNG")

	//main menu on window open
	menu := container.NewGridWithRows(2,
		// container.NewBorder(
		// 	top : layout.NewVBoxLayout(),
		// 	layout.NewSpacer(),
		// 	layout.NewSpacer(),
		// 	layout.NewSpacer(),

		// ),
		container.NewGridWithColumns(2,
			container.NewVBox(
				toolbar,
				widget.NewButton("Scan Position Statement", func() {
					img.FillMode = canvas.ImageFillOriginal
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

	//when using the X close button, the app is only closed and minimized to the tray
	win.SetCloseIntercept(func() {
		win.Hide()
	})
	win.SetMaster()
	win.Show()
	a.Run()
}
