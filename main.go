package main

import (
	"github.com/yuuis/PersonalDataRepository/api"
	"github.com/yuuis/PersonalDataRepository/infrastructures"
	"log"
)

func main() {
	mongo, err := infrastructures.OpenMongo()

	if err != nil {
		log.Fatal(err)
	}

	s := infrastructures.NewServer()
	api.Router(s, mongo)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
