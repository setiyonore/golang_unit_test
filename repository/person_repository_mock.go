package repository

import (
	"belajar-golang-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

type PersonRepositoryMock struct {
	Mock mock.Mock
}

func (repository *PersonRepositoryMock) FindByAge(age string) *entity.Person {
	arguments := repository.Mock.Called(age)
	if arguments.Get(0) == nil {
		return nil
	} else {
		person := arguments.Get(0).(entity.Person)
		return &person
	}
}

func (repository *PersonRepositoryMock) FindById(id string) *entity.Person {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		person := arguments.Get(0).(entity.Person)
		return &person
	}
}

func (repository *PersonRepositoryMock) FindByName(name string) *entity.Person {
	arguments := repository.Mock.Called(name)
	if arguments.Get(0) == nil {
		return nil
	} else {
		person := arguments.Get(0).(entity.Person)
		return &person
	}
}
