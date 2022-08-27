package search

import (
	"context"
	"errors"
	"github.com/rocketlaunchr/google-search"
	"strings"
)

//simple google search for tax rate from zipcode - replaces an API until a free API is available or found
func SearchForTaxRate(zipCode string) (string, error) {

	search, err := googlesearch.Search(context.TODO(), "tax rate of "+zipCode)
	if err != nil {
		return "Search error", err
	}

	for _, v := range search {
		if strings.Contains(v.Description, "%") {
			pos := strings.Index(v.Description, "%")
			if pos == -1 {
				break
			}
			tax := v.Description[(pos - 5) : pos+1]
			tax = strings.TrimSpace(tax)
			tax = strings.TrimPrefix(tax, " is ")
			tax = strings.TrimPrefix(tax, "is")
			tax = strings.TrimPrefix(tax, " ")
			tax = strings.TrimPrefix(tax, "s")
			tax = strings.TrimPrefix(tax, "a")
			return tax, nil
		}
	}

	return "Zip Not found", errors.New("could not locate tax")
}
