package main

import (
	"github.com/yuuis/PersonalDataRepository/api"
	"github.com/yuuis/PersonalDataRepository/infrastructures"
	"github.com/yuuis/PersonalDataRepository/models/basic"
	"github.com/yuuis/PersonalDataRepository/models/location"
	"log"
)

func main() {
	mysql, err := infrastructures.OpenMysql()

	if err != nil {
		log.Fatal(err)
	}

	mysql.AutoMigrate(&basic.Basic{}, &location.Location{})

	s := infrastructures.NewServer()
	api.Router(s, mysql)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
