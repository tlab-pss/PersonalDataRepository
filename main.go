package main

import (
	"github.com/yuuis/PersonalDataRepository/api"
	"github.com/yuuis/PersonalDataRepository/infrastructures"
	"github.com/yuuis/PersonalDataRepository/models"
	"log"
)

func main() {
	mysql, err := infrastructures.OpenMysql()

	if err != nil {
		log.Fatal(err)
	}

	mysql.AutoMigrate(&models.Basic{}, &models.Location{})

	s := infrastructures.NewServer()
	api.Router(s)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
