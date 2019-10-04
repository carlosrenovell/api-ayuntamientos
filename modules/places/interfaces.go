package places

import (
	. "gitlab.com/euthinkia/api-ayuntamientos/modules/commons/entities"
)

//Repositories
type (
	CountriesRepositoryInterface interface {
		FindAll() ([]*Country, error)
		FindById(id int) (*Country, error)
		Save(*Country) error
	}
	ComunitiesRepositoryInterface interface {
		FindAll() (*[]Comunity, error)
		FindById(id int) (*Comunity, error)
		Save(*Comunity) error
	}
	ProvincesRepositoryInterface interface {
		FindAll() (*[]Province, error)
		FindById(id int) (*Province, error)
		Save(*Province) error
	}
	CitiesRepositoryInterface interface {
		FindAll() (*[]City, error)
		FindById(id int) (*City, error)
		Save(*City) error
	}
)
