package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Vehicle struct {
	Name     string
	Price    int
	Currency string
	Year     string
	Details  string
	Link     string
	// Image    string
}

func main() {
	c := colly.NewCollector()

	var Vehicles []Vehicle
	// All div with class car
	c.OnHTML("div.car", func(e *colly.HTMLElement) {
		price, currency := priceConvert(e.ChildText(".price"))
		vehicle := Vehicle{
			Name:     e.ChildText(".title"),
			Price:    price,
			Currency: currency,
			Year:     e.ChildText(".uk-text-primary"),
			Details:  e.ChildText(".dtil"),
			Link:     e.ChildAttr("a", "href"),
			// Image:    e.ChildAttr("img", "src"),
		}

		Vehicles = append(Vehicles, vehicle)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:",
			r.Request.URL,
			"failed with response:",
			r, "\nError:", err)
	})

	c.Visit("https://carrosrd.com/supercarros")
	for _, car := range Vehicles {
		text := `
Vehiculo: %s
Price: %d
Currency: %s
Year: %s
Details: %s
Link: %s
`
		fmt.Printf(text, car.Name,
			car.Price,
			car.Currency,
			car.Year,
			car.Details, car.Link)

		printSeparator(45, "*")
	}
	fmt.Println(len(Vehicles))
}

func priceConvert(price string) (int, string) {
	var currency string

	// US$42,800

	if strings.Contains(price, "US$") {
		currency = "USD"
		price = strings.ReplaceAll(price, "US$", "")

	} else {
		currency = "DOP"
		price = strings.ReplaceAll(price, "RD$", "")

	}

	price = strings.ReplaceAll(price, ",", "")
	price = strings.TrimSpace(price)

	value, _ := strconv.Atoi(price)

	return value, currency

}

func printSeparator(times int, sep string) {
	for i := 0; i <= times; i++ {
		fmt.Print(sep)
	}
	fmt.Println()
}

// <div class="car ">
// <a title="Mercedes-Benz Clase GLE" href="https://carrosrd.com/carros/Mercedes-Benz/Clase GLE/350 4Matic/l-79130">
// <div class="overlay-car">
// <i class="uk-icon-sign-out"></i>
// <span>Ver detalles</span>
// </div>
// <div class="price uk-position-bottom-left">US$43,900</div>
// <img width="227" height="171"
// src="https://carrosrd-media.s3.amazonaws.com/listings/79130/m_17745516180564201.jpg?1"
//  alt="Mercedes-Benz Clase GLE 350 4Matic en venta">
// </a><div class="details-car">
// <div class="uk-flex uk-flex-space-between uk-text-bold uk-margin-small-bottom">
// span class="title">Mercedes-Benz Clase GLE 350 4Matic</span>
// <span class="uk-text-primary">2020</span>
// </div>
// <div class="dtil uk-text-muted"
// style="font-size:12px">
// Gasolina | 2000cc | AWD | 5</div>
// </div>
// </div>
