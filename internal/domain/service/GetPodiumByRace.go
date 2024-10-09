package domain_service

import (
	"sync"

	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
)

type GetPodiumByRaceDependencies interface {
	GetDriversByRaceDependencies
	FetchPodiumByRace(raceId int, drivers []domain_model.Driver) ([3]domain_model.Podium, error)
}

type fetchPodiumsResponse struct {
	raceId   int
	response [3]domain_model.Podium
	error    error
}

func GetPodiumByRace(
	races []domain_model.Race,
	dependencies GetPodiumByRaceDependencies,
) (map[int][3]domain_model.Podium, error) {

	podiumsByRace := make(map[int][3]domain_model.Podium)

	drivers, error := GetDriversByRace(races, dependencies)
	if error != nil {
		return nil, error
	}

	var waitGroup sync.WaitGroup
	respChannel := make(chan fetchPodiumsResponse, 8)
	for _, race := range races {
		waitGroup.Add(1)
		go fetchPodiums(
			race.Id,
			drivers[race.Id],
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

		podiumsByRace[resp.raceId] = resp.response
	}

	return podiumsByRace, nil
}

func fetchPodiums(
	raceId int,
	drivers []domain_model.Driver,
	dependencies GetPodiumByRaceDependencies,
	respChan chan<- fetchPodiumsResponse,
	waitGroup *sync.WaitGroup,
) {
	defer waitGroup.Done()

	result, error := dependencies.FetchPodiumByRace(raceId, drivers)

	if error != nil {
		respChan <- fetchPodiumsResponse{
			raceId:   raceId,
			response: [3]domain_model.Podium{},
			error:    error,
		}
		return
	}

	respChan <- fetchPodiumsResponse{
		raceId:   raceId,
		response: result,
		error:    nil,
	}
}
