package commons

import (
	"fmt"
	"gitlab.com/euthinkia/api-ayuntamientos/modules/commons/configuration"
	"labix.org/v2/mgo"
	"log"
)

const (
	// DBNAME the name of the DB instance
	DBNAME = "Euthinkia"
)

//Aqui deberiamos realizar las conexiones que los repos deben usar
type Connection struct {
	DbName  string
	Session *mgo.Session
	err     error
}

func NewConnection() Connection {
	con := Connection{}
	return con
}
func (con *Connection) Connect() (*mgo.Session, error) {
	fmt.Println("Ejecutando Connect()")
	con.Session, con.err = noSslConnect()
	return con.Session, con.err
}

//noSslConnect: TODO The idea of this is to migrate to an ssl connection to mongoDB
func noSslConnect() (*mgo.Session, error) {
	fmt.Println("Ejecutando noSslConnect()")
	// session, err := mgo.Dial(os.Getenv("MONGO_URI"))
	conf := configuration.NewConfigDriver()
	url := conf.Url
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal("Error en la conexion:", err)
		panic(err)
	}

	return session, err
}
