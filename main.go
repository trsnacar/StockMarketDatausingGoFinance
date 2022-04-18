package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/piquette/finance-go/quote"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		logrus.Fatalf("Bir hisse senedi sembolü girin", os.Args[0])
	}
	symbol := flag.Args()[0]
	a, _ := quote.Get(symbol)
	fmt.Printf(" --%v--\n", a.ShortName)
	fmt.Printf("Mevcut fiyat:$%v\n", a.Ask)
	fmt.Printf("52Hft Yüksek:$%v\n", a.FiftyTwoWeekHigh)
	fmt.Printf("52Hft Düşük:$%v\n", a.FiftyTwoWeekLow)

}
