package main

import (
	"flag"
	"fmt"
	"http/geo"
	"http/weather"
)

func main() {
	city := flag.String("city", "", "Город")
	format := flag.Int("format", 1, "Тип ответа")

	flag.Parse()
	geoData, _ := geo.GetMyLocation(*city)
	fmt.Println(geoData)
	
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}