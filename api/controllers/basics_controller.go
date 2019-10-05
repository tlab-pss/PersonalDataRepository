package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/basic"
	"time"
)

func (r *Registry) GetBasic(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := basic.NewDataStore(r.db)

	b, err := ds.Get()

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.BasicView(ctx, *b)
	}
}

func (r *Registry) CreateBasic(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := basic.NewDataStore(r.db)

	var ipt inputBasic
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	// TODO: validationどっかにまとめたい
	if err := basic.ValidateName(ipt.Name); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	if err := basic.ValidateGender(ipt.Gender); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	tb, err := time.Parse("2006-01-02", ipt.Birthday)

	if err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	b, err := ds.Store(&basic.Basic{
		ID:        models.GenerateUUID(),
		Name:      ipt.Name,
		Birthday:  tb,
		Gender:    ipt.Gender,
		CreatedAt: time.Now(),
	})

	if err != nil {
		presenters.ViewInternalServerError(ctx, err)
	}

	presenters.BasicView(ctx, *b)
}

type inputBasic struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
}
