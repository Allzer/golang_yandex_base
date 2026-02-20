package weather

import (
	"fmt"
	"http/geo"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil{
		fmt.Println(err.Error())
		return ""
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))

	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())

	body, _ := io.ReadAll(resp.Body)

	return string(body)

}