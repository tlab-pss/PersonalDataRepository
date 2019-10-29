package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/health"
	"time"
)

func (r *Registry) GetHealth(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := health.NewDataStore(r.db)
	h, err := ds.GetLatest()

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.HealthView(ctx, *h)
	}
}

func (r *Registry) CreateHealth(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := health.NewDataStore(r.db)

	var ipt inputHealth
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
		return
	}

	h, err := ds.Store(&health.Health{
		ID:        models.GenerateUUID(),
		Weight:    ipt.Weight,
		Height:    ipt.Height,
		HeartRate: ipt.HeartRate,
		CreatedAt: time.Time{},
	})

	if err != nil {
		presenters.ViewInternalServerError(ctx, err)
		return
	}

	presenters.HealthView(ctx, *h)
}

type inputHealth struct {
	Weight    float64 `json:"weight"`
	Height    float64 `json:"height"`
	HeartRate int     `json:"heartRate"`
}
