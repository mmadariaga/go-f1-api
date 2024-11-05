package application_drivers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	domain_model "github.com/mmadariaga/go-api/internal/domain/model"
	test_domain_model "github.com/mmadariaga/go-api/internal/test/domain/model"
)

func TestGetLastRaceDrivers(t *testing.T) {

	assert := assert.New(t)

	races := test_domain_model.Suzuka()
	drivers := test_domain_model.GetAllDriverExamples()

	mockInfraDependencies := new(MockInfraDependencies)
	mockInfraDependencies.On(
		"FetchLastRace",
	).Return(
		races[0],
		nil,
	)
	mockInfraDependencies.On(
		"FetchDriversByRace",
		races[0].Id,
	).Return(
		drivers,
		nil,
	)

	LastRaceDrivers := NewLastRaceDrivers(mockInfraDependencies)

	results, error := LastRaceDrivers.Get()

	assert.Greater(len(results), 1)

	assert.Equal(error, nil)
	assert.Equal(results[0].Number, 1)
	assert.Equal(results[0].FullName, "Max VERSTAPPEN")

	mockInfraDependencies.AssertExpectations(t)
}

func TestGetLastRaceDriversErrorHandling(t *testing.T) {

	assert := assert.New(t)

	mockInfraDependencies := new(MockInfraDependencies)
	mockInfraDependencies.On(
		"FetchLastRace",
	).Return(
		domain_model.Race{},
		errors.New("Unable to fetch races"),
	)

	LastRaceDrivers := NewLastRaceDrivers(mockInfraDependencies)
	_, error := LastRaceDrivers.Get()

	assert.NotEqual(error, nil)
	assert.Equal(error.Error(), "Unable to fetch races")

	mockInfraDependencies.AssertExpectations(t)
}

type MockInfraDependencies struct {
	mock.Mock
}

func (m *MockInfraDependencies) FetchLastRace() (domain_model.Race, error) {
	args := m.Called()
	return args.Get(0).(domain_model.Race), args.Error(1)
}
func (m *MockInfraDependencies) FetchDriversByRace(raceId int) ([]domain_model.Driver, error) {
	args := m.Called(raceId)
	return args.Get(0).([]domain_model.Driver), args.Error(1)
}
