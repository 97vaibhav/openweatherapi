package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//helps us getting open weather api key

type apiConfigData struct {
	ApiKey string `json:"ApiKey`
}

type weatherData struct {
	Name string `json:"name`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApi(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	// unmarshel () =json(byte) data----->struct
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to weather api this is the starting point. To get information about weather hit end point weather/{ city name} !\n"))
}

func Getweatherdata(city string) (weatherData, error) {
	apiConfig, err := loadApi(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}
	//api.openweathermap.org/data/2.5/find?q=London&units=metric
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.ApiKey + "&q=" + city + "&units=metric")
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}

func main() {
	http.HandleFunc("/welcome", welcome)

	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := Getweatherdata(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
		})

	http.ListenAndServe(":8080", nil)
}
