package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/registration_information"
	"time"
)

func (r *Registry) GetRegistrationInformation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := registration_information.NewDataStore(r.db)
	ri, err := ds.Get()

	if err != nil {
		presenters.ViewInternalServerError(ctx, err)
	}

	presenters.RegistrationInformationView(ctx, *ri)
}

func (r *Registry) CreateRegistrationInformation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := registration_information.NewDataStore(r.db)

	var ipt inputRegistrationInformation
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	if err := registration_information.ValidateMail(ipt.Mail); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	ri, err := ds.Store(&registration_information.RegistrationInformation{
		ID:        models.GenerateUUID(),
		Mail:      ipt.Mail,
		CreatedAt: time.Time{},
	})

	if err != nil {
		presenters.ViewInternalServerError(ctx, err)
	}

	presenters.RegistrationInformationView(ctx, *ri)
}

type inputRegistrationInformation struct {
	Mail string `json:mail`
}
