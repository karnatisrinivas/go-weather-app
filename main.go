package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type WeatherResponse struct {
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
}

type PageData struct {
	Weather WeatherResponse
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	location := r.FormValue("location")

	// Read the API key from the configuration file using viper
	apiKey := viper.GetString("app_id")

	// Replace the app_id in the API URL with the API key
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", location, apiKey)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var weather WeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&weather); err != nil {
		log.Fatal(err)
	}

	data := PageData{
		Weather: weather,
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initialize viper
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}

	http.HandleFunc("/weather", weatherHandler)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
