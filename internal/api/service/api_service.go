package service

import (
	"encoding/json"
	"github.com/DrNikita/People/internal/api/model"
	config "github.com/DrNikita/People/internal/db"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func GetAgeInfo(name string) (ageInfo model.AgeInfo, err error) {
	configs := config.GetConfigurationInstance()
	client := http.Client{}
	req, err := http.NewRequest("GET", configs.AgeUrl+name, nil)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &ageInfo); err != nil {
		log.Error(err)
	}
	return
}

func GetGenderInfo(name string) (genderInfo model.GenderInfo, err error) {
	configs := config.GetConfigurationInstance()
	client := http.Client{}
	req, err := http.NewRequest("GET", configs.GenderUrl+name, nil)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &genderInfo); err != nil {
		log.Error(err)
	}
	return
}

func GetCountryInfo(name string) (countryInfo model.CountryInfo, err error) {
	configs := config.GetConfigurationInstance()
	client := http.Client{}
	req, err := http.NewRequest("GET", configs.CountryUrl+name, nil)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &countryInfo); err != nil {
		log.Error(err)
	}
	return
}
