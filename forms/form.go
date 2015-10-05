/*
Package forms contains all the form data and metadata
*/
package forms

import (
	"encoding/json"
	"io/ioutil"
)

var Config *config

type config struct {
	FullName       [][]int `json:"full_name"`
	FullName2      []int   `json:"full_name_2"`
	Gender         [][]int `json:"gender"`
	PassportNumber [][]int `json:"passport_number"`
	FlightNumber   [][]int `json:"flight_no"`
	CountryOfBirth [][]int `json:"country_of_birth"`
	DateOfBirth    [][]int `json:"date_of_birth"`
	Nationality    [][]int `json:"nationality"`
	Nationality2   []int   `json:"nationality_2"`
	LastCity       [][]int `json:"last_city"`
	MalaysianIC    [][]int `json:"identity_card_no_malaysian"`
	MalaysianIC2   [][]int `json:"identity_card_no_malaysian_2"`
}

func Init() {
	configFile, err := ioutil.ReadFile("forms/form.json")
	if err != nil {
		panic("Form config file open error: " + err.Error())
	}

	Config = &config{}
	err = json.Unmarshal(configFile, Config)
	if err != nil {
		panic("Failed to parse form config json")
	}
}
