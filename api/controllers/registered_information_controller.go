package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/registered_information"
	"time"
)

func (r *Registry) GetRegisteredInformation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := registered_information.NewDataStore(r.db)

	ri, err := ds.Get()

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.RegisteredInformationView(ctx, *ri)
	}
}

func (r *Registry) CreateRegisteredInformation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := registered_information.NewDataStore(r.db)

	var ipt inputRegisteredInformation
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	if err := registered_information.ValidateMail(ipt.Mail); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	ri, err := ds.Store(&registered_information.RegisteredInformation{
		ID:        models.GenerateUUID(),
		Mail:      ipt.Mail,
		CreatedAt: time.Time{},
	})

	if err != nil {
		presenters.ViewInternalServerError(ctx, err)
	}

	presenters.RegisteredInformationView(ctx, *ri)
}

type inputRegisteredInformation struct {
	Mail string `json:mail`
}
