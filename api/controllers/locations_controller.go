package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/location"
	"time"
)

func (r *Registry) GetLocation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := location.NewDataStore(r.db)
	l, err := ds.GetLatest()

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.LocationView(ctx, *l)
	}
}

func (r *Registry) CreateLocation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := location.NewDataStore(r.db)

	var ipt inputLocation
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	l, err := ds.Store(&location.Location{
		ID:        models.GenerateUUID(),
		Latitude:  ipt.Latitude,
		Longitude: ipt.Longitude,
		CreatedAt: time.Now(),
	})

	if err != nil {
		presenters.ViewInternalServerError(ctx, err)
	}

	presenters.LocationView(ctx, *l)
}

type inputLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
