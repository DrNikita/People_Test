package service

import (
	apiModel "github.com/DrNikita/People/internal/api/model"
	"github.com/DrNikita/People/internal/api/service"
	"github.com/DrNikita/People/internal/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AddInfo(person *model.PersonFullInfo, tx *gorm.DB) {
	ageInfo, err := service.GetAgeInfo()
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return
	}
	log.Debug("get age info")
	genderInfo, err := service.GetGenderInfo()
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return
	}
	log.Debug("get gender info")
	countryInfo, err := service.GetCountryInfo()
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return
	}
	log.Debug("get country info")

	err = tx.Model(&model.PersonFullInfo{}).Where("id = ?", person.ID).Updates(map[string]interface{}{
		"age":        ageInfo.Age,
		"gender":     genderInfo.Gender,
		"country_id": getMostLikelyCountry(countryInfo.Country),
	}).Error
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return
	}
	log.Debug("person created; age, gender, country added")
	tx.Commit()
}

func getMostLikelyCountry(countries []apiModel.County) string {
	if len(countries) == 0 {
		return ""
	}
	var maxProbability float64
	var mostLikelyCountryIndex int
	for i, country := range countries {
		if maxProbability < country.Probability {
			maxProbability = country.Probability
			mostLikelyCountryIndex = i
		}
	}
	return countries[mostLikelyCountryIndex].CountryId
}
