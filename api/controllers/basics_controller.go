package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/basic"
	"net/http"
	"time"
)

func (r *Registry) GetBasic(c *gin.Context) {
	ds := basic.NewDataStore(r.db)
	b, err := ds.Get()

	if err != nil {
		presenters.ViewInternalServerError(c, err)
	}

	presenters.JSON(c, http.StatusOK, b)
}

func (r *Registry) CreateBasic(c *gin.Context) {
	ds := basic.NewDataStore(r.db)

	var ipt inputBasic
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(c, err)
	}

	// TODO: validationどっかにまとめたい
	if err := basic.ValidateName(ipt.Name); err != nil {
		presenters.ViewBadRequest(c, err)
	}

	if err := basic.ValidateMail(ipt.Mail); err != nil {
		presenters.ViewBadRequest(c, err)
	}

	if err := basic.ValidateGender(ipt.Gender); err != nil {
		presenters.ViewBadRequest(c, err)
	}

	b, err := ds.Store(&basic.Basic{
		ID:        models.GenerateUUID(),
		Name:      ipt.Name,
		Birthday:  ipt.Birthday,
		Gender:    ipt.Gender,
		Mail:      ipt.Mail,
		Weight:    ipt.Weight,
		Height:    ipt.Height,
		CreatedAt: time.Now(),
	})

	if err != nil {
		presenters.ViewInternalServerError(c, err)
	}

	presenters.JSON(c, http.StatusOK, b)
}

type inputBasic struct {
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender   int       `json:"gender"`
	Mail     string    `json:"mail"`
	Weight   float64   `json:"weight"`
	Height   float64   `json:"heigth"`
}
