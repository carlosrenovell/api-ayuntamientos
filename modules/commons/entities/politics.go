package entities

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	/*Politics*/
	PoliticalParty struct {
		gorm.Model
		ID        bson.ObjectId `bson:"_id" json:"id"`
		Hash      string        `bson:"hash" json:"hash"`
		Name      string        `bson:"name" json:"name"`
		Birthday  time.Time     `bson:"birthday" json:"birthday"`
		CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
		UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
		DeletedAt time.Time     `bson:"deletedAt" json:"deletedAt"`
	}
	//Politician  representa políticos
	Politician struct {
		gorm.Model
		ID               bson.ObjectId    `bson:"_id" json:"id"`
		Hash             string           `bson:"hash" json:"hash"`
		Name             string           `bson:"name" json:"name"`
		Lastnames        string           `bson:"lastnames" json:"lastnames"`
		Birthday         time.Time        `bson:"birthday" json:"birthday"`
		Photo            string           `bson:"photo" json:"photo"`
		PoliticalParties []PoliticalParty `bson:"politicalParties" json:"politicalParties"`
		CreatedAt        time.Time        `bson:"createdAt" json:"createdAt"`
		UpdatedAt        time.Time        `bson:"updatedAt" json:"updatedAt"`
		DeletedAt        time.Time        `bson:"deletedAt" json:"deletedAt"`
	}
	//TownHalls object with all key:value
	//TODO quizas habria que cambiar TownHalls por Governments para una mejor estructura
	TownHall struct {
		ID        bson.ObjectId `bson:"_id" json:"id"`
		Hash      string        `bson:"hash" json:"hash"`
		Cities    []int         `bson:"cities" json:"cities"`
		Comunity  int           `bson:"comunity" json:"comunity"`
		Country   int           `bson:"country" json:"country"`
		CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
		UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
		DeletedAt time.Time     `bson:"deletedAt" json:"deletedAt"`
	}
	Proposal struct {
		ID   bson.ObjectId `bson:"_id" json:"id"`
		Hash string        `bson:"hash" json:"hash"`
		//Quienes realizan propuesta: Partido político o un político hace esa propuesta
		PoliticalPartyProposing int `bson:"politicalPartyProposing" json:"politicalPartyProposing"`
		PoliticianProposing     int `bson:"politicianProposing" json:"politicianProposing"`
		//ElectionProposed es para qué elecciones está propuesto
		ElectionProposed int    `bson:"election" json:"election"`
		Title            string `bson:"title" json:"title"`
		Description      string `bson:"description" json:"description"`

		CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
		UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
		DeletedAt time.Time `bson:"deletedAt" json:"deletedAt"`
	}
)

func (PoliticalParty) TableName() string {
	return "PoliticalParties"
}
func (Politician) TableName() string {
	return "Politicians"
}
