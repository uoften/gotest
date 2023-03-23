package main
import (
	"fmt"
	g "github.com/serpapi/google-search-results-golang"
)

func main(){
	parameter := map[string]string{
		"engine": "google",
		"q": "uoften2adsa",
		"api_key": "168b73db82cbe6ec04252989605977715636d5bf9b612015a07cfc30eab756f0",
	}

	search := g.NewGoogleSearch(parameter, "168b73db82cbe6ec04252989605977715636d5bf9b612015a07cfc30eab756f0")
	results, err := search.GetJSON()
	if err != nil {
		fmt.Println(err)
	}
	organic_results := results["organic_results"].([]interface{})

	fmt.Println(organic_results)
}