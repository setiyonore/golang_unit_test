package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type PersonService struct {
	Repository repository.PersonRepository
}

func (service PersonService) Get(age string) (*entity.Person, error) {
	person := service.Repository.FindByAge(age)
	if person == nil {
		return person, errors.New("Person not found")
	} else {
		return person, nil
	}
}

func (service PersonService) GetId(id string) (*entity.Person, error) {
	person := service.Repository.FindById(id)
	if person == nil {
		return person, errors.New("Person not found")
	} else {
		return person, nil
	}
}

func (service PersonService) GetName(name string) (*entity.Person, error) {
	person := service.Repository.FindByName(name)
	if person == nil {
		return person, errors.New("Person not found")
	} else {
		return person, nil
	}
}
