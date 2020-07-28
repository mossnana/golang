package main

import "github.com/gin-gonic/gin"

/* func main() {
	// New Type
	type job struct {
		Title   string `bson:"title"`
		Company string `bson:"company"`
	}

	type test struct {
		Name   string `bson:"name"`
		Number int    `bson:"number"`
	}

	newJob := job{Title: "Web Engineer", Company: "MO IO"}
	// t := test{Name: "test", Number: 200}

	session, error := mgo.Dial("localhost:27017")
	if error != nil {
		log.Fatal(error)
	}

	mongoConnection := session.DB("local").C("golangTest")
	error = mongoConnection.Insert(newJob)
	if error != nil {
		log.Fatal(error)
	}
	session.Close()
} */

var router *gin.Engine

func main() {
	router = gin.Default()
	initRoutes()
	router.Run()
}
