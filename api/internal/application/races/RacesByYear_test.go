package application_races

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
	test_domain_model "github.com/mmadariaga/go-api/internal/test/domain/model"
)

func TestGetRacesByYear(t *testing.T) {

	assert := assert.New(t)

	suzuka := test_domain_model.Suzuka()
	drivers := test_domain_model.GetAllDriverExamples()
	podiums := test_domain_model.GetPodiumWithDrivers(
		"Max VERSTAPPEN",
		"Fernando ALONSO",
		"Lewis HAMILTON",
	)

	mockInfraDependencies := new(MockInfraDependencies)
	mockInfraDependencies.On(
		"FetchRacesByYear",
		2024,
	).Return(
		suzuka,
		nil,
	)
	mockInfraDependencies.On(
		"FetchPodiumByRace",
		suzuka[0].Id,
		drivers,
	).Return(
		podiums,
		nil,
	)
	mockInfraDependencies.On(
		"FetchDriversByRace",
		suzuka[0].Id,
	).Return(
		drivers,
		nil,
	)

	RacesByYear := NewRacesByYear(mockInfraDependencies)

	result, error := RacesByYear.Get(2024)

	assert.Equal(error, nil)
	assert.Equal(result[0].Race.CircuitName, "Suzuka")
	assert.Equal(result[0].Podium[0].Position, 1)
	assert.Equal(result[0].Podium[0].FullName, "Max VERSTAPPEN")
	assert.Equal(result[0].Podium[1].FullName, "Fernando ALONSO")

	mockInfraDependencies.AssertExpectations(t)
}

func TestFetchRacesByYearErrorHandling(t *testing.T) {

	assert := assert.New(t)

	mockInfraDependencies := new(MockInfraDependencies)
	mockInfraDependencies.On(
		"FetchRacesByYear",
		2024,
	).Return(
		make([]domain_model.Race, 0),
		errors.New("Unable to fetch races"),
	)

	RacesByYear := NewRacesByYear(mockInfraDependencies)
	_, error := RacesByYear.Get(2024)

	assert.NotEqual(error, nil)
	assert.Equal(error.Error(), "Unable to fetch races")

	mockInfraDependencies.AssertExpectations(t)
}

func TestFetchDriversByRaceErrorHandling(t *testing.T) {

	assert := assert.New(t)

	suzuka := test_domain_model.Suzuka()

	mockInfraDependencies := new(MockInfraDependencies)
	mockInfraDependencies.On(
		"FetchRacesByYear",
		2024,
	).Return(
		suzuka,
		nil,
	)
	mockInfraDependencies.On(
		"FetchDriversByRace",
		suzuka[0].Id,
	).Return(
		make([]domain_model.Driver, 0),
		errors.New("Unable to fetch drivers"),
	)

	RacesByYear := NewRacesByYear(mockInfraDependencies)
	_, error := RacesByYear.Get(2024)

	assert.NotEqual(error, nil)
	assert.Equal(error.Error(), "Unable to fetch drivers")

	mockInfraDependencies.AssertExpectations(t)
}

func TestFetchPodiumByRaceErrorHandling(t *testing.T) {

	assert := assert.New(t)

	suzuka := test_domain_model.Suzuka()

	mockInfraDependencies := new(MockInfraDependencies)
	mockInfraDependencies.On(
		"FetchRacesByYear",
		2024,
	).Return(
		suzuka,
		nil,
	)
	mockInfraDependencies.On(
		"FetchDriversByRace",
		suzuka[0].Id,
	).Return(
		make([]domain_model.Driver, 0),
		nil,
	)
	mockInfraDependencies.On(
		"FetchPodiumByRace",
		suzuka[0].Id,
		[]domain_model.Driver{},
	).Return(
		[3]domain_model.Podium{},
		errors.New("Unable to fetch podiums"),
	)

	RacesByYear := NewRacesByYear(mockInfraDependencies)
	_, error := RacesByYear.Get(2024)

	assert.NotEqual(error, nil)
	assert.Equal(error.Error(), "Unable to fetch podiums")

	mockInfraDependencies.AssertExpectations(t)
}

type MockInfraDependencies struct {
	mock.Mock
}

func (m *MockInfraDependencies) FetchRacesByYear(year int) ([]domain_model.Race, error) {
	args := m.Called(year)
	return args.Get(0).([]domain_model.Race), args.Error(1)
}
func (m *MockInfraDependencies) FetchPodiumByRace(raceId int, drivers []domain_model.Driver) ([3]domain_model.Podium, error) {
	args := m.Called(raceId, drivers)
	return args.Get(0).([3]domain_model.Podium), args.Error(1)
}
func (m *MockInfraDependencies) FetchDriversByRace(raceId int) ([]domain_model.Driver, error) {
	args := m.Called(raceId)
	return args.Get(0).([]domain_model.Driver), args.Error(1)
}
