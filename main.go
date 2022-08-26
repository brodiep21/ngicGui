package main

import (
	"fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "image/color"
	// "github.com/rocketlaunchr/google-search"
)

func main() {
	myApp := app.New()
	win := myApp.NewWindow("NGIC")
	// g := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	username := &widget.Entry{PlaceHolder: "Username"}
	// text1 := canvas.NewText("Hello", g)
	// text2 := canvas.NewText("There", g)
	// text2.Move(fyne.NewPos(20, 20))
	// content := container.NewWithoutLayout(text1, text2)
	menu := container.NewGridWithRows(3,
		container.NewGridWithColumns(3,
			layout.NewSpacer(),
			container.NewVBox(
				username,
				&widget.Button{Text: "Login", Icon: theme.ConfirmIcon()},
			),
		),
	)
	// content := container.New(layout.NewGridLayout(2), text1, text2)

	// win.SetContent(content)
	win.SetContent(menu)
	win.ShowAndRun()
}

// func searchForTaxRate(zipCode string) (string, error) {

// 	search, err := googlesearch.Search(context.TODO(), "tax rate of "+zipCode)
// 	if err != nil {
// 		return "", err
// 	}

// 	for _, v := range search {
// 		if strings.Contains(v.Description, "%") {
// 			pos := strings.Index(v.Description, "%")
// 			if pos == -1 {
// 				break
// 			}
// 			tax := v.Description[(pos - 6) : pos+1]
// 			tax = strings.TrimPrefix(tax, "is ")
// 			tax = strings.TrimPrefix(tax, " is ")
// 			tax = strings.TrimPrefix(tax, "is")
// 			tax = strings.TrimPrefix(tax, " ")
// 			tax = strings.TrimPrefix(tax, "s")
// 			tax = strings.TrimPrefix(tax, "a")
// 			return tax, nil
// 		}
// 	}

// 	return "", errors.New("could not locate tax")
// }
