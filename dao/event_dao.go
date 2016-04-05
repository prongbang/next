package dao

import (
	"github.com/prongbang/next/structs"
	"github.com/prongbang/next/connect"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
)

const (
	INDEX = "position"
)

type DaoConfig struct {
	m *connect.MgoConfig
	COLLECTION_EVENT string
}

func (dao *DaoConfig) FindAll() []structs.Event {

	session := dao.m.NewSession()
	c := dao.m.NewCollection(session, dao.COLLECTION_EVENT)
	var result []structs.Event
	err := c.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}
	dao.m.Close(session)

	return result
}

func (dao *DaoConfig) FindByLatLng(locat structs.Request) []structs.Event {

	fmt.Println("Location:", locat)

	// TODO Find By Lat & Lng

	session := dao.m.NewSession()
	c := dao.m.NewCollection(session, dao.COLLECTION_EVENT)
	var result []structs.Event

	// ref : https://www.safaribooksonline.com/blog/2013/07/24/analyze-data-with-mongo-and-go/
	// ref : https://stackoverflow.com/questions/26710271/how-can-i-find-nearby-place-with-latitude-and-longitude-in-mongodb
	// ref : https://dba.stackexchange.com/questions/22813/mongodb-returning-the-points-nearest-to-a-location-with-distance
	// ref : https://blog.codecentric.de/en/2012/02/spring-data-mongodb-geospatial-queries/
	// ref : https://docs.mongodb.org/manual/reference/operator/query/near/
	// ref : https://docs.mongodb.org/manual/reference/operator/query/near/#examples

	// :: Find event by between latitude and longitude :: //
	// TODO openshift error => can't find any special indices: 2d (needs index), 2dsphere (needs index),
	// http://grokbase.com/t/gg/mgo-users/13ahpaqctt/doing-ensure-index-for-2dsphere

	// TODO BUG

	// Creating the indexes
	//$2dsphere:position
	//$2d:position
	index := dao.m.Index(INDEX)
	//c.EnsureIndexKey("position")
	c.EnsureIndex(index)
	//$minDistance: <distance in meters>
	//{"position": {"$near" : {"$geometry" : { "type" : "Point", "coordinates": [-84.096466, 9.934351] } } } }
	err := c.Find(bson.M{
		"position": bson.M{
			"$near": bson.M{
				"$geometry" : bson.M{
					"type" : "Point",
					"coordinates":[]float64{ locat.Lat, locat.Lng },
				},
			},
			"$maxDistance": locat.Dist,
		},
	}).All(&result)

	if err != nil {
		panic(err)
	} else {
		//c.DropIndex(INDEX)
	}

	dao.m.Close(session)

	return result
}

func (dao *DaoConfig) FindById(id uint32) structs.Event {

	session := dao.m.NewSession()
	c := dao.m.NewCollection(session, dao.COLLECTION_EVENT)
	var result structs.Event
	err := c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		panic(err)
	}
	dao.m.Close(session)

	return result
}

func (dao *DaoConfig) Save(inf interface{}, collection string) int {
	status := 1
	session := dao.m.NewSession()
	c := dao.m.NewCollection(session, collection)
	err := c.Insert(inf)
	if err != nil {
		status = 0
		log.Fatal(err)
	} else {
		c.EnsureIndexKey(INDEX)
	}
	return status
}

func (dao *DaoConfig) clearEvent() int {
	status := 1
	session := dao.m.NewSession()
	c := dao.m.NewCollection(session, dao.COLLECTION_EVENT)
	_, err := c.RemoveAll(nil)
	if err != nil {
		status = 0
	}
	return status
}