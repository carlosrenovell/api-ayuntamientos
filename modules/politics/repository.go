package politics

import (
	"fmt"
	"gitlab.com/euthinkia/api-ayuntamientos/modules/commons"
	. "gitlab.com/euthinkia/api-ayuntamientos/modules/commons/entities"

	"log"
)

type (

	//implementaciones
	PartiesRepo struct {
		CollectionName string
	}
	PoliticiansRepo struct {
		CollectionName string
	}
	TownHallsRepo struct {
		CollectionName string
	}
	ProposalsRepo struct {
		CollectionName string
	}
)

func NewPartiesRepo() PartiesRepo {
	repo := PartiesRepo{}
	repo.CollectionName = "PoliticalParties"
	return repo
}
func (r *PartiesRepo) FindAll() (*[]PoliticalParty, error) {
	var results []PoliticalParty
	con := commons.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Find(nil).All(&results); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &results, nil
}
func (r *PartiesRepo) FindById(id int) (*PoliticalParty, error) {
	var result PoliticalParty
	con := commons.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.FindId(id).One(&result); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &result, nil
}

func (r *PartiesRepo) Save(input *PoliticalParty) error {
	con := commons.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Insert(input); err != nil {
		log.Println("Failed to write results:", err)
		return err
	}
	return nil
}

func NewPoliticiansRepo() PoliticiansRepo {
	repo := PoliticiansRepo{}
	repo.CollectionName = "Politicians"
	return repo
}

func (r *PoliticiansRepo) FindAll() (*[]Politician, error) {
	var results []Politician
	con := commons.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Find(nil).All(&results); err != nil {
		log.Println("Failed to write results:", err)
	}

	return &results, nil

}
func (r *PoliticiansRepo) FindById(id int) (*Politician, error) {
	var result Politician
	con := commons.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.FindId(id).One(&result); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &result, nil
}

func (r *PoliticiansRepo) Save(input *Politician) error {
	con := commons.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Insert(input); err != nil {
		log.Println("Failed to write results:", err)
		return err
	}
	return nil
}

func NewTownHallsRepo() TownHallsRepo {
	repo := TownHallsRepo{}
	repo.CollectionName = "TownHalls"
	return repo
}

// GetAllTownhalls returns the list of Products
func (r *TownHallsRepo) FindAll() (*[]TownHall, error) {
	fmt.Println("Ejecutando func (r *TownHallsRepo) FindAll() (*[]TownHall, error) ")
	var results []TownHall
	con := commons.NewConnection()
	defer con.Session.Close()

	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Find(nil).All(&results); err != nil {
		log.Println("Failed to write results:", err)
	}

	return &results, nil
}

/*InsertTHalls insert the apiKey TODO::Check if the appName exist
 * dont save the apiKey
 */
func (r *TownHallsRepo) Save(input TownHall) error {
	con := commons.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Insert(input); err != nil {
		log.Println("Failed to write results:", err)
		return err
	}
	return nil
}

func NewProposalsRepo() ProposalsRepo {
	repo := ProposalsRepo{}
	repo.CollectionName = "Proposals"
	return repo
}

func (r *ProposalsRepo) FindAll() (*[]TownHall, error) {
	var results []TownHall
	con := commons.NewConnection()
	defer con.Session.Close()

	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Find(nil).All(&results); err != nil {
		log.Println("Failed to write results:", err)
	}

	return &results, nil
}

func (r *ProposalsRepo) Save(input *Proposal) error {
	con := commons.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(commons.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Insert(input); err != nil {
		log.Println("Failed to write results:", err)
		return err
	}
	return nil
}
