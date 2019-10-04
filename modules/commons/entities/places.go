package entities

import (
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	//Coordenate struct for townhall object
	Coordenate struct {
		ID        bson.ObjectId `bson:"_id" json:"id"`
		Hash      string        `bson:"hash" json:"hash"`
		Latitude  string        `bson:"latitude" json:"latitude"`
		Longitude string        `bson:"longitude" json:"longitude"`
		CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
		UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
		DeletedAt time.Time     `bson:"deletedAt" json:"deletedAt"`
	}

	//Country struct
	Country struct {
		ID          bson.ObjectId `bson:"_id" json:"id"`
		Hash        string        `bson:"hash" json:"hash"`
		Name        string        `bson:"name" json:"name"`
		Description string        `bson:"description" json:"description"`
		// Coordenates    []Coordenate   `bson:"coordenates" json:"coordenates"` //TODO: Tengo mis dudas de si las coordenadas debería ser un slide
		// Maior          Politician   int:"maior" json:"maior"`
		Coordenates    []int          `bson:"int" json:"coordenates"` //TODO: Tengo mis dudas de si las coordenadas debería ser un slide
		Maior          int            `bson:"maior" json:"maior"`
		PoliticalParty PoliticalParty `bson:"politicalparty" json:"politicalparty"`
		POPULATION     int            `bson:"population,omitempty" json:"population"`
		WEBSITE        string         `bson:"website" json:"website"`
		CreatedAt      time.Time      `bson:"createdAt" json:"createdAt"`
		UpdatedAt      time.Time      `bson:"updatedAt" json:"updatedAt"`
		DeletedAt      time.Time      `bson:"deletedAt" json:"deletedAt"`
	}
	// Comunity struct
	Comunity struct {
		ID             bson.ObjectId  `bson:"_id" json:"id"`
		Hash           string         `bson:"hash" json:"hash"`
		Name           string         `bson:"name" json:"name"`
		Description    string         `bson:"description" json:"description"`
		Coordenates    []Coordenate   `bson:"coordenates" json:"coordenates"` //TODO: Tengo mis dudas de si las coordenadas debería ser un slide
		President      []Politician   `bson:"maior" json:"maior"`
		PoliticalParty PoliticalParty `bson:"politicalparty" json:"politicalparty"`
		POPULATION     int            `bson:"population" json:"population"`
		WEBSITE        string         `bson:"website" json:"website"`
		Country        Country        `bson:"country" json:"country"`
		CreatedAt      time.Time      `bson:"createdAt" json:"createdAt"`
		UpdatedAt      time.Time      `bson:"updatedAt" json:"updatedAt"`
		DeletedAt      time.Time      `bson:"deletedAt" json:"deletedAt"`
	}
	// Province struct
	Province struct {
		ID             bson.ObjectId `bson:"_id" json:"id"`
		Hash           string        `bson:"hash" json:"hash"`
		Name           string        `bson:"name" json:"name"`
		Description    string        `bson:"description" json:"description"`
		Coordenates    []int         `bson:"coordenates" json:"coordenates"` //TODO: Tengo mis dudas de si las coordenadas debería ser un slide
		President      int           `bson:"president" json:"president"`
		PoliticalParty int           `bson:"politicalparty" json:"politicalparty"`
		POPULATION     int           `bson:"population" json:"population"`
		WEBSITE        string        `bson:"website" json:"website"`
		Country        int           `bson:"country" json:"country"`
		CreatedAt      time.Time     `bson:"createdAt" json:"createdAt"`
		UpdatedAt      time.Time     `bson:"updatedAt" json:"updatedAt"`
		DeletedAt      time.Time     `bson:"deletedAt" json:"deletedAt"`
	}
	// City struct
	City struct {
		ID          bson.ObjectId `bson:"_id" json:"id"`
		Hash        string        `bson:"hash" json:"hash"`
		Name        string        `bson:"name" json:"name"`
		Description string        `bson:"description" json:"description"`
		// Coordenates    []Coordenate   `bson:"coordenates" json:"coordenates"` //TODO: Tengo mis dudas de si las coordenadas debería ser un slide
		// President      []Politician   `bson:"president" json:"president"`
		// PoliticalParty PoliticalParty `bson:"politicalparty" json:"politicalparty"`

		Coordenates    []int `bson:"coordenates" json:"coordenates"` //TODO: Tengo mis dudas de si las coordenadas debería ser un slide
		President      int   `bson:"president" json:"president"`
		PoliticalParty int   `bson:"politicalparty" json:"politicalparty"`

		POPULATION int       `bson:"population" json:"population"`
		WEBSITE    string    `bson:"website" json:"website"`
		CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
		UpdatedAt  time.Time `bson:"updatedAt" json:"updatedAt"`
		DeletedAt  time.Time `bson:"deletedAt" json:"deletedAt"`
	}
)

//TownHall Obj to slice
// type TownHalls []TownHall

//TableName TODO: Fill this
func (Coordenate) TableName() string {
	return "Coordenates"
}

//EntityName TODO: Fill
func (Coordenate) EntityName() string {
	return "Coordenate"
}

//TableName TODO: Fill this
func (Province) TableName() string {
	return "Provinces"
}

//EntityName TODO: Fill this
func (Province) EntityName() string {
	return "Provinces"
}

//TableName TODO: Fill this
func (City) TableName() string {
	return "Cities"
}

//EntityName TODO: Fill this
func (City) EntityName() string {
	return "City"
}

//TableName TODO: Fill this
func (TownHall) TableName() string {
	return "TownHalls"
}

//EntityName TODO: Fill this
func (TownHall) EntityName() string {
	return "TownHall"
}
