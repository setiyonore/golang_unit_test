package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var personRepository = &repository.PersonRepositoryMock{Mock: mock.Mock{}}
var personService = PersonService{Repository: personRepository}

func TestPersonService_GetNotFound(t *testing.T) {
	//program mock
	personRepository.Mock.On("FindByAge", "1").Return(nil)
	person, err := personService.Get("1")
	assert.Nil(t, person)
	assert.NotNil(t, err)
}

func TestPersonService_GetSuccess(t *testing.T) {
	person := entity.Person{
		Id:      "2",
		Name:    "Eko",
		Age:     "30",
		Married: true,
	}
	personRepository.Mock.On("FindByAge", "2").Return(person)
	result, err := personService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, person.Id, result.Id)
	assert.Equal(t, person.Name, result.Name)
	assert.Equal(t, person.Age, result.Age)
	assert.Equal(t, person.Married, result.Married)
}

func TestPersonService_GetByIdNotFound(t *testing.T) {
	//program mock
	personRepository.Mock.On("FindById", "1").Return(nil)
	result, err := personService.GetId("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestPersonService_GetByIdSuccess(t *testing.T) {
	//program mock
	person := entity.Person{
		Id:      "2",
		Name:    "Eko",
		Age:     "30",
		Married: true,
	}
	personRepository.Mock.On("FindById", "2").Return(person)
	result, err := personService.GetId("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, person.Id, result.Id)
	assert.Equal(t, person.Name, result.Name)
	assert.Equal(t, person.Age, result.Age)
	assert.Equal(t, person.Married, result.Married)
}

func TestPersonService_GetByNameNotFound(t *testing.T) {
	//program mock
	personRepository.Mock.On("FindByName", "Eko").Return(nil)
	result, err := personService.GetName("Eko")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestPersonService_GetNameSuccess(t *testing.T) {
	person := entity.Person{
		Id:      "3",
		Name:    "Budi",
		Age:     "29",
		Married: false,
	}
	personRepository.Mock.On("FindByName", "Budi").Return(person)
	result, err := personService.GetName("Budi")
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, person.Id, result.Id)
	assert.Equal(t, person.Name, result.Name)
	assert.Equal(t, person.Age, result.Age)
	assert.Equal(t, person.Married, result.Married)
}
