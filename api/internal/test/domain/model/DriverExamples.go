package test_domain_model

import (
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

var drivers []domain_model.Driver

func InitDrivers() {

	if len(drivers) > 0 {
		return
	}

	drivers = []domain_model.Driver{
		{
			Number:   1,
			FullName: "Max VERSTAPPEN",
			Country:  "NED",
			Avatar:   "https://www.formula1.com/content/dam/fom-website/drivers/M/MAXVER01_Max_Verstappen/maxver01.png.transform/1col/image.png",
			TeamName: "Red Bull Racing",
		},
		{
			Number:   2,
			FullName: "Fernando ALONSO",
			Country:  "ESP",
			Avatar:   "https://www.formula1.com/content/dam/fom-website/drivers/F/FERALO01_Fernando_Alonso/feralo01.png.transform/1col/image.png",
			TeamName: "Aston Martin",
		},
		{
			Number:   3,
			FullName: "Lewis HAMILTON",
			Country:  "GBR",
			Avatar:   "https://www.formula1.com/content/dam/fom-website/drivers/L/LEWHAM01_Lewis_Hamilton/lewham01.png.transform/1col/image.png",
			TeamName: "Mercedes",
		},
		{
			Number:   2,
			FullName: "Carlos SAINZ",
			Country:  "ESP",
			Avatar:   "https://www.formula1.com/content/dam/fom-website/drivers/C/CARSAI01_Carlos_Sainz/carsai01.png.transform/1col/image.png",
			TeamName: "Ferrari",
		},
	}
}

func GetAllDriverExamples() []domain_model.Driver {
	InitDrivers()

	return drivers
}

func GetAllDriverExamplesMap() map[string]domain_model.Driver {

	var response = make(map[string]domain_model.Driver)
	drivers := GetAllDriverExamples()

	for _, val := range drivers {

		if val.FullName == "" {
			continue
		}

		response[val.FullName] = val
	}

	return response
}

func GetMaxVerstappen() domain_model.Driver {
	drivers := GetAllDriverExamples()

	return drivers[0]
}

func GetFernandoAlonso() domain_model.Driver {
	drivers := GetAllDriverExamples()

	return drivers[1]
}

func GetLewisHamilton() domain_model.Driver {
	drivers := GetAllDriverExamples()

	return drivers[2]
}
