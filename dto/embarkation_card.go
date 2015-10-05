package dto

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type EmbarkationCardRequest struct {
	Fullname            string `json:"full_name"`
	Gender              string `json:"gender"`
	PassportNumber      string `json:"passport_number"`
	PlaceOfIssue        string `json:"place_of_issue"`
	ExpiryDate          string `json:"expiry_date"`
	CountryOfBirth      string `json:"country_of_birth"`
	DateOfBirth         string `json:"date_of_birth"`
	Nationality         string `json:"nationality"`
	FlightCode          string `json:"flight_code"`
	EmbarkationLocation string `json:"embarkation_location"`
	MalaysianIC         string `json:"malaysian_ic"`
	Format              string
}

func (r EmbarkationCardRequest) CheckFormat() error {
	errorStrings := make([]string, 0, 10)
	if r.Fullname == "" {
		errorStrings = append(errorStrings, "full_name field missing")
	}
	genderRegex, _ := regexp.Compile("^[mMfF]$")
	if !genderRegex.MatchString(r.Gender) {
		errorStrings = append(errorStrings, "gender invalid")
	}
	if r.PassportNumber == "" {
		errorStrings = append(errorStrings, "passport_number field missing")
	}
	if r.PlaceOfIssue == "" {
		errorStrings = append(errorStrings, "place_of_issue field missing")
	}
	if _, err := time.Parse("02-01-2006", r.ExpiryDate); err != nil {
		errorStrings = append(errorStrings, "expiry_date field is invalid")
	}
	if r.CountryOfBirth == "" {
		errorStrings = append(errorStrings, "country_of_birth field missing")
	}
	if _, err := time.Parse("02-01-2006", r.DateOfBirth); err != nil {
		errorStrings = append(errorStrings, "date_of_birth field is invalid")
	}
	if r.Nationality == "" {
		errorStrings = append(errorStrings, "nationality field is missing")
	}
	regex, _ := regexp.Compile("^[A-Z]+[0-9]+$")
	if !regex.MatchString(r.FlightCode) {
		errorStrings = append(errorStrings, "flight_code field is invalid")
	}
	if r.EmbarkationLocation == "" {
		errorStrings = append(errorStrings, "embarkation_location field is missing")
	}

	if len(errorStrings) > 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

func (r *EmbarkationCardRequest) Parse() {
	r.Fullname = strings.ToUpper(r.Fullname)
	r.Gender = strings.ToUpper(r.Gender)
	r.PassportNumber = strings.ToUpper(r.PassportNumber)
	r.PlaceOfIssue = strings.ToUpper(r.PlaceOfIssue)
	r.CountryOfBirth = strings.ToUpper(r.CountryOfBirth)
	r.Nationality = strings.ToUpper(r.Nationality)
	r.FlightCode = strings.ToUpper(r.FlightCode)
	r.EmbarkationLocation = strings.ToUpper(r.EmbarkationLocation)
	r.MalaysianIC = strings.ToUpper(r.MalaysianIC)
}
