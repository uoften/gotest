package main

import (
	"fmt"
	"github.com/negah/alexa"
)

func main(){
	url := "uoften.com"

	globalRank, err := alexa.GlobalRank(url)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s rank in alexa is %s\n", url, globalRank)
	}

	countryRank, countryName, _, err := alexa.CountryRank(url)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s has rank %s in %s", url, countryRank, countryName )
	}
}