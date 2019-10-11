package main

import (
	"github.com/yuuis/PersonalDataRepository/api"
	"github.com/yuuis/PersonalDataRepository/db/mongo"
	"github.com/yuuis/PersonalDataRepository/infrastructures"
	"log"
)

func main() {
	m, err := infrastructures.OpenMongo()

	if err != nil {
		log.Fatal(err)
	}

	// migrate
	if err := mongo.Seed(m); err != nil {
		log.Fatal(err)
	}

	s := infrastructures.NewServer()
	api.Router(s, m)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
