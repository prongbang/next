package connect

import (
	"gopkg.in/mgo.v2"
	"os"
	"fmt"
)

type MgoConfig struct {
	DATABASE string
	USERNAME string
	PASSWORD string
	LOCAL_OR_HOST string
	MONGODB_DB_HOST string
	MONGODB_DB_PORT string
}

func (mg *MgoConfig) NewSession() *mgo.Session {

	var session *mgo.Session
	var err interface{}

	if mg.LOCAL_OR_HOST == 1 {
		Host := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_MONGODB_DB_HOST"), os.Getenv("OPENSHIFT_MONGODB_DB_PORT"))
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{Host},
			Username: mg.USERNAME,
			Password: mg.PASSWORD,
			Database: mg.DATABASE,
		})
	} else {
		session, err = mgo.Dial(fmt.Sprintf("%s:%s", mg.MONGODB_DB_HOST, mg.MONGODB_DB_PORT))
	}
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

func (mg *MgoConfig) NewCollection(session *mgo.Session, collection string) *mgo.Collection {
	return session.DB(mg.DATABASE).C(collection)
}

func (mg *MgoConfig) Close(session *mgo.Session) {
	defer session.Close()
}

func (mg *MgoConfig) Index(column string) mgo.Index {
	return mgo.Index{
		Key: []string{"$2dsphere:" + column },
		Bits: 26,
	}
}
