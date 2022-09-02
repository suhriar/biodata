package handler

import (
	"biodata/models"
	"biodata/seeder"
)

func GetBiodataById(id string) models.Students {
	var biodata models.Students

	biodatas := seeder.Biodata()

	for _, val := range biodatas {
		if val.Id == id {
			biodata = val
			break
		}
	}
	return biodata
}