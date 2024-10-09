package test_domain_model

import (
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

var drivers = make(map[string]domain_model.Driver)

func InitDrivers() {

	if len(drivers) > 0 {
		return
	}

	drivers["Max VERSTAPPEN"] = domain_model.Driver{
		Number:   1,
		FullName: "Max VERSTAPPEN",
		Country:  "NED",
		Avatar:   "https://www.formula1.com/content/dam/fom-website/drivers/M/MAXVER01_Max_Verstappen/maxver01.png.transform/1col/image.png",
		TeamName: "Red Bull Racing",
	}

	drivers["Fernando ALONSO"] = domain_model.Driver{
		Number:   2,
		FullName: "Fernando ALONSO",
		Country:  "ESP",
		Avatar:   "https://www.formula1.com/content/dam/fom-website/drivers/F/FERALO01_Fernando_Alonso/feralo01.png.transform/1col/image.png",
		TeamName: "Aston Martin",
	}

	drivers["Lewis HAMILTON"] = domain_model.Driver{
		Number:   3,
		FullName: "Lewis HAMILTON",
		Country:  "GBR",
		Avatar:   "https://www.formula1.com/content/dam/fom-website/drivers/L/LEWHAM01_Lewis_Hamilton/lewham01.png.transform/1col/image.png",
		TeamName: "Mercedes",
	}

	drivers["Carlos SAINZ"] = domain_model.Driver{
		Number:   2,
		FullName: "Carlos SAINZ",
		Country:  "ESP",
		Avatar:   "https://www.formula1.com/content/dam/fom-website/drivers/C/CARSAI01_Carlos_Sainz/carsai01.png.transform/1col/image.png",
		TeamName: "Ferrari",
	}
}

func GetAllDriverExamples() map[string]domain_model.Driver {
	InitDrivers()

	return drivers
}

func GetAllDriverExamplesAsArray() []domain_model.Driver {
	drivers := GetAllDriverExamples()

	response := make([]domain_model.Driver, 0, 10)

	for _, val := range drivers {
		response = append(response, val)
	}

	return response
}
