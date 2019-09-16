package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/location"
	"net/http"
	"time"
)

func (r *Registry) GetLocation(c *gin.Context) {
	ds := location.NewDataStore(r.db)
	l, err := ds.GetLatest()

	if err != nil {
		presenters.ViewInternalServerError(c, err)
	}

	presenters.JSON(c, http.StatusOK, l)
}

func (r *Registry) CreateLocation(c *gin.Context) {
	ds := location.NewDataStore(r.db)

	var ipt inputLocation
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(c, err)
	}

	l, err := ds.Store(&location.Location{
		ID:        models.GenerateUUID(),
		Latitude:  ipt.Latitude,
		Longitude: ipt.Longitude,
		CreatedAt: time.Now(),
	})

	if err != nil {
		presenters.ViewInternalServerError(c, err)
	}

	presenters.JSON(c, http.StatusOK, l)
}

type inputLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
