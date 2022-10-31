package repository

import "belajar-golang-unit-test/entity"

type PersonRepository interface {
	FindByAge(age string) *entity.Person
	FindById(id string) *entity.Person
	FindByName(name string) *entity.Person
}
