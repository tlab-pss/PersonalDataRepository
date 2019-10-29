package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/basic_location"
	"github.com/yuuis/PersonalDataRepository/models/location"
	"time"
)

func (r *Registry) GetBasicLocation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := basic_location.NewDataStore(r.db)

	b, err := ds.Get()

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.BasicLocationView(ctx, *b)
	}
}

func (r *Registry) CreateBasicLocation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := basic_location.NewDataStore(r.db)

	var ipt inputBasicLocation
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
		return
	}

	b, err := ds.Store(&basic_location.BasicLocation{
		ID: models.GenerateUUID(),
		House: location.Location{
			ID:        models.GenerateUUID(),
			Latitude:  ipt.House.Lat,
			Longitude: ipt.House.Lng,
			CreatedAt: time.Now(),
		},
		Office: location.Location{
			ID:        models.GenerateUUID(),
			Latitude:  ipt.Office.Lat,
			Longitude: ipt.Office.Lng,
			CreatedAt: time.Now(),
		},
		Route:     nil,
		CreatedAt: time.Time{},
	})

	if err != nil {
		presenters.ViewInternalServerError(ctx, err)
		return
	}

	presenters.BasicLocationView(ctx, *b)
}

type inputBasicLocation struct {
	House struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"house"`
	Office struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"office"`
	// Route todo: いつか
}
