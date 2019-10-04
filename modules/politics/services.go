//Business logic about politics
package politics

import (
	. "gitlab.com/euthinkia/api-ayuntamientos/modules/commons/entities"
	"time"
)

type (
	PartiesService struct {
		repo PartiesRepo
	}
	PoliticiansService struct {
		repo PoliticiansRepo
	}
)

//Crea el objeto de partido politico para que el repositorio haga la insercion
//TODO: Generar aleatoriamente el campo Hash
func (this *PartiesService) CreateParty(name string, birthday time.Time) (*PoliticalParty, error) {
	party := PoliticalParty{}
	party.Name = name
	party.Birthday = birthday
	party.CreatedAt = time.Now()
	party.UpdatedAt = time.Now()
	return &party, nil
}

func (this *PartiesService) InsertParty(party *PoliticalParty) error {
	this.repo = PartiesRepo{}
	err := this.repo.Save(party)
	return err
}

//Crea el objeto de politico para que el repositorio haga la insercion
//TODO: Generar aleatoriamente el campo Hash
func (this *PoliticiansService) CreatePolitician(name string, birthday time.Time) (*Politician, error) {
	politician := Politician{}
	politician.Name = name
	politician.Birthday = birthday
	politician.CreatedAt = time.Now()
	politician.UpdatedAt = time.Now()
	return &politician, nil
}
func (this *PoliticiansService) InsertPolitician(politician *Politician) error {
	this.repo = PoliticiansRepo{}

	err := this.repo.Save(politician)

	return err
}

// func (this *PoliticiansService) CreateObject(title string, description string,proposing ProposingInterface, proposed) error {
// 	this.repo = PoliticiansRepo{}

// 	err := this.repo.Save(politician)

// 	return err
// }
