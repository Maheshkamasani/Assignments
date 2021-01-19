package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name        string `json:"name"`
	Time        string `json:"time"`
	Information Info   `json:"info"`
}

type Info struct {
	Description string `json:"des"`
}

func main() {
	fmt.Println("Hello world")
	jsonString := `{"name": "Maheswar Reddy kamasani", "time": "2020-11-13", "info": {
		"des": "A Kloudone Trainee"
	}}`
	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}