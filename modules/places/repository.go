package places

import (
	"gitlab.com/euthinkia/api-ayuntamientos/modules/commons"
	. "gitlab.com/euthinkia/api-ayuntamientos/modules/commons/entities"

	// "labix.org/v2/mgo"
	"log"
)

type (
	CountriesRepo struct {
		CollectionName string
	}
	ComunitiesRepo struct {
		CollectionName string
	}
	ProvincesRepo struct {
		CollectionName string
	}
	CitiesRepo struct {
		CollectionName string
	}
)

func NewCountriesRepo() CountriesRepo {
	repo := CountriesRepo{}
	repo.CollectionName = "Countries"
	return repo
}
func (r *CountriesRepo) FindAll() (*[]Country, error) {
	var results []Country
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Find(nil).All(&results); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &results, nil

}
func (r *CountriesRepo) FindById(id int) (*Country, error) {
	var result Country
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.FindId(id).One(&result); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &result, nil
}
func (r *CountriesRepo) Save(input *Country) error {
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Insert(input); err != nil {
		log.Println("Failed to write results:", err)
		return err
	}
	return nil
}

func NewComunitiesRepo() ComunitiesRepo {
	repo := ComunitiesRepo{}
	repo.CollectionName = "Countries"
	return repo
}

func (r *ComunitiesRepo) FindAll() (*[]Comunity, error) {
	var results []Comunity
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Find(nil).All(&results); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &results, nil
}
func (r *ComunitiesRepo) FindById(id int) (*Comunity, error) {
	var result Comunity
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.FindId(id).One(&result); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &result, nil
}
func (r *ComunitiesRepo) Save(input *Comunity) error {
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Insert(input); err != nil {
		log.Println("Failed to write results:", err)
		return err
	}
	return nil
}
func NewProvincesRepo() ProvincesRepo {
	repo := ProvincesRepo{}
	repo.CollectionName = "Provinces"
	return repo
}
func (r *ProvincesRepo) FindAll() (*[]Province, error) {
	var results []Province
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Find(nil).All(&results); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &results, nil
}
func (r *ProvincesRepo) FindById(id int) (*Province, error) {
	var result Province
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.FindId(id).One(&result); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &result, nil
}
func (r *ProvincesRepo) Save(input *Province) error {
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Insert(input); err != nil {
		log.Println("Failed to write results:", err)
		return err
	}
	return nil
}
func NewCitiesRepo() CitiesRepo {
	repo := CitiesRepo{}
	repo.CollectionName = "Cities"
	return repo
}
func (r *CitiesRepo) FindAll() (*[]City, error) {
	var results []City
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.Find(nil).All(&results); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &results, nil
}
func (r *CitiesRepo) FindById(id int) (*City, error) {
	var result City
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")
	if err := c.FindId(id).One(&result); err != nil {
		log.Println("Failed to write results:", err)
		return nil, err
	}

	return &result, nil
}
func (r *CitiesRepo) Save(input *City) error {
	con := common.NewConnection()
	defer con.Session.Close()
	c := con.Session.DB(common.DBNAME).C(r.CollectionName)
	log.Println("Getting records of " + r.CollectionName + "...")

	if err := c.Insert(input); err != nil {
		log.Println("Failed to write results:", err)
		return err
	}
	return nil
}
