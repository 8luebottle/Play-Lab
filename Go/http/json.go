package main

import (
	"encoding/json"
	"net/http"
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"
)

const WeatherURL = "https://www.7timer.info/bin/astro.php?lon=113.2&lat=23.1&ac=0&unit=metric&output=json&tzshift=0"

type Period string

const (
	Day   Period = "1"
	Week  Period = "7"
	Month Period = "30"
)

// Parsing JSON examples
func main() {
	dayResp := GetWeatherResp(Day)
	ExampleMarshalMap(dayResp) // Parse JSON with Map

	weekResp := GetWeatherResp(Week)
	ExampleMarshalStruct(weekResp) // Parse JSON with Struct
}

func GetWeatherResp(period Period) *http.Response {
	client := &http.Client{
		Timeout: 5 * time.Second, // Since default value is 0, it's important to set deadline.
		// This prevents server crashing (avoids resource limits).
	}
	resp, err := client.Get(WeatherURL + string(period))
	if err != nil {
		logrus.Error(err)
	}

	return resp
}

// ExampleMarshalMap binds body to `map[string]interface{}`
// This suits arbitrary JSON data of unknown schema.
func ExampleMarshalMap(resp *http.Response) {
	body := resp.Body

	var data map[string]interface{} // Value : interface{} can hold any types of value.
	// interface{} data will allocate heap memories.

	logrus.Printf("Size of Empty map : %d", unsafe.Sizeof(data)) // 8
	// The size of Empty map(8) should smaller than the size of Empty struct(56).

	if err := json.NewDecoder(body).Decode(&data); err != nil {
		logrus.Error(err)
	}
	d, err := json.MarshalIndent(data, "", "\t") // Pretty print JSON data.
	if err != nil {
		logrus.Error(err)
	}

	logrus.Printf("Size of exampleMap : %T %d", d, unsafe.Sizeof(d)) // []uint8 24
	// Ultimately, it's the same as size of struct.

	logrus.Printf("exampleMap : %+v\n", string(d)) // The result.
}

type Timer struct {
	Product    string `json:"product"`
	Init       string `json:"init"`
	DataSeries []Data `json:"dataseries"`
}

type Data struct {
	TimePoint    int          `json:"timepoint"`
	CloudCover   int          `json:"cloudcover"`
	Seeing       int          `json:"seeing"`
	Transparency int          `json:"transparency"`
	LiftedIndex  int          `json:"lifted_index"`
	RH2Minute    int          `json:"rh2m"`
	Wind10Minute Wind10Minute `json:"wind10m"`
	Temp2Minute  int          `json:"temp2m"`
	PrecType     string       `json:"prec_type"`
}

type Wind10Minute struct {
	Direction string `json:"direction"`
	Speed     int    `json:"speed"`
}

// ExampleMarshalStruct binds body to a custom nested struct, Timer.
// Use it if you know JSON fields.
func ExampleMarshalStruct(resp *http.Response) {
	body := resp.Body

	var timer Timer

	logrus.Printf("Size of Empty struct : %d", unsafe.Sizeof(timer)) // 56
	// The size of Empty struct(56) is bigger than the size of empty map(8).

	if err := json.NewDecoder(body).Decode(&timer); err != nil {
		logrus.Error(err)
	}
	d, err := json.MarshalIndent(timer, "", "\t") // Pretty print JSON data.
	if err != nil {
		logrus.Error(err)
	}
	logrus.Printf("Size of struct : %T %d", d, unsafe.Sizeof(d)) // []uint8 24
	// Ultimately, it's the same as size of map.

	logrus.Printf("\n exampleStruct : %+v", string(d)) // The result.
}
