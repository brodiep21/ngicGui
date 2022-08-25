package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	// "github.com/rocketlaunchr/google-search"
)

func main() {
	a := app.New()

	w := a.NewWindow("MDS Work")

	// .
	w.SetContent(widget.NewLabel(""))
	w.ShowAndRun()

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
