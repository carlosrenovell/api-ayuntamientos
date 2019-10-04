package entities

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//Election representa cada eleccion
	Election struct {
		ID        bson.ObjectId `bson:"_id" json:"id"`
		Hash      string        `bson:"hash" json:"hash"`
		Date      time.Time
		CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
		UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
		DeletedAt time.Time `bson:"deletedAt" json:"deletedAt"`
	}
	//Representa unas elecciones con cada partido en concreto
	ResultElection struct {
		ID   bson.ObjectId `bson:"_id" json:"id"`
		Hash string        `bson:"hash" json:"hash"`
		// Election       Election       `bson:"election" json:"election"`
		// PoliticalParty PoliticalParty `bson:"politicalParty" json:"politicalParty"`
		Election       int `bson:"election" json:"election"`
		PoliticalParty int `bson:"politicalParty" json:"politicalParty"`
		Date           time.Time
		CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
		UpdatedAt      time.Time `bson:"updatedAt" json:"updatedAt"`
		DeletedAt      time.Time `bson:"deletedAt" json:"deletedAt"`
	}
)

// TableName table name
func (Election) TableName() string {
	return "Elections"
}

// TableName table name
func (ResultElection) TableName() string {
	return "ResultsElections"
}
