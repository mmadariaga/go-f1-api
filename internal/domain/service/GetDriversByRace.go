package domain_service

import (
	"sync"

	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

type GetDriversByRaceDependencies interface {
	FetchDriversByRace(raceId int) ([]domain_model.Driver, error)
}

type fetchDriversResponse struct {
	raceId   int
	response []domain_model.Driver
	error    error
}

func GetDriversByRace(races []domain_model.Race, dependencies GetDriversByRaceDependencies) (map[int][]domain_model.Driver, error) {

	driversByRace := make(map[int][]domain_model.Driver)

	var waitGroup sync.WaitGroup
	respChannel := make(chan fetchDriversResponse, 8)

	for _, race := range races {
		waitGroup.Add(1)
		go fetchDrivers(
			race.Id,
			dependencies,
			respChannel,
			&waitGroup,
		)
	}

	go func() {
		waitGroup.Wait()
		close(respChannel)
	}()

	for resp := range respChannel {

		if resp.error != nil {
			return nil, resp.error
		}

		driversByRace[resp.raceId] = resp.response
	}

	return driversByRace, nil
}

func fetchDrivers(
	raceId int,
	dependencies GetDriversByRaceDependencies,
	respChan chan<- fetchDriversResponse,
	waitGroup *sync.WaitGroup,
) {
	defer waitGroup.Done()

	result, error := dependencies.FetchDriversByRace(raceId)

	if error != nil {
		respChan <- fetchDriversResponse{
			raceId:   raceId,
			response: nil,
			error:    error,
		}
		return
	}

	respChan <- fetchDriversResponse{
		raceId:   raceId,
		response: result,
		error:    nil,
	}
}
