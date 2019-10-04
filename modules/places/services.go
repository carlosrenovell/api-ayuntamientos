package places

import (
	. "/entities"
	// "time"
)

type (
	CountriesService struct {
		Repo CountriesRepo
	}
	CitiesService struct {
		Repo CitiesRepo
	}
	ComunitiesService struct {
		Repo ComunitiesRepo
	}
	ProvincesService struct {
		Repo ProvincesRepo
	}
)

func NewCountriesService() CountriesService {
	s := CountriesService{}
	s.Repo = NewCountriesRepo()
	return s
}
func (s *CountriesService) CreateCountry(name string, description string) Country {
	country := Country{}
	country.Name = name
	country.Description = description
	return country
}

func NewCitiesService() CitiesService {
	s := CitiesService{}
	s.Repo = NewCitiesRepo()
	return s
}
func NewComunitiesService() ComunitiesService {
	s := ComunitiesService{}
	s.Repo = NewComunitiesRepo()
	return s
}
func NewProvincesService() ProvincesService {
	s := ProvincesService{}
	s.Repo = NewProvincesRepo()
	return s
}
