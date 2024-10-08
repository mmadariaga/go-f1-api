package test_domain_model

import (
	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

var podiums = [3]domain_model.Podium{}

func InitPodiums() {

	if podiums[0].FullName != "" {
		return
	}

	drivers := GetAllDriverExamples()

	podiums[0] = domain_model.Podium{
		Position: 1,
		Driver:   drivers["Max VERSTAPPEN"],
	}

	podiums[1] = domain_model.Podium{
		Position: 2,
		Driver:   drivers["Fernando ALONSO"],
	}

	podiums[2] = domain_model.Podium{
		Position: 3,
		Driver:   drivers["Lewis HAMILTON"],
	}
}

func GetPodiumWithDrivers(driver1, driver2, driver3 string) [3]domain_model.Podium {
	podiums := [3]domain_model.Podium{}

	podiums[0] = domain_model.Podium{
		Position: 1,
		Driver:   drivers[driver1],
	}

	podiums[1] = domain_model.Podium{
		Position: 2,
		Driver:   drivers[driver2],
	}

	podiums[2] = domain_model.Podium{
		Position: 3,
		Driver:   drivers[driver3],
	}

	return podiums
}

func GetAllPodiumExamples() [3]domain_model.Podium {
	InitPodiums()

	return podiums
}
