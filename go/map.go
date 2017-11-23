package main
import "fmt"

func main() {
    var countryCapitalMap map[string]string

    countryCapitalMap = make(map[string]string)

    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Rome"

    for country := range countryCapitalMap {
        fmt.Println("Captial of", country, "is", countryCapitalMap[country])
    }

    capital, ok := countryCapitalMap["United States"]

    if(ok) {
        fmt.Println("Captial of United States is", capital)
    }else {
        fmt.Println("Captial of United States is not a pointer")
    }
}
