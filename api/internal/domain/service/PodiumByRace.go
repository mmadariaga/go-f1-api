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

type PodiumByRace struct {
	Dependencies GetPodiumByRaceDependencies
}

func (p *PodiumByRace) Get(
	races []domain_model.Race,
) (map[int][3]domain_model.Podium, error) {

	podiumsByRace := make(map[int][3]domain_model.Podium)

	driversByRace := DriversByRace{
		Dependencies: p.Dependencies,
	}

	drivers, error := driversByRace.Get(races)
	if error != nil {
		return nil, error
	}

	var waitGroup sync.WaitGroup
	respChannel := make(chan fetchPodiumsResponse, 8)
	for _, race := range races {
		waitGroup.Add(1)
		go p.fetchPodiums(
			race.Id,
			drivers[race.Id],
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

func (p *PodiumByRace) fetchPodiums(
	raceId int,
	drivers []domain_model.Driver,
	respChan chan<- fetchPodiumsResponse,
	waitGroup *sync.WaitGroup,
) {
	defer waitGroup.Done()

	result, error := p.Dependencies.FetchPodiumByRace(raceId, drivers)

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
