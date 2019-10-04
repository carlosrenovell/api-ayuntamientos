package politics

import (
	. "gitlab.com/euthinkia/api-ayuntamientos/modules/commons/entities"
)

type (
	TownHallsRepositoryInterface interface {
		FindAll() (*[]TownHall, error)
		FindById(id int) (*TownHall, error)
		Save(*TownHall) error
	}
	PoliticiansRepository interface {
		FindAll() (*[]Politician, error)
		FindById(id int) (*Politician, error)
		Save(*Politician) error
	}
	PartiesRepository interface {
		FindAll() (*[]PoliticalParty, error)
		FindById(id int) (*PoliticalParty, error)
		Save(*PoliticalParty) error
	}
	ProposalsRepository interface {
		FindAll() (*[]Proposal, error)
		FindById(id int) (*Proposal, error)
		Save(*Proposal) error
	}
)
