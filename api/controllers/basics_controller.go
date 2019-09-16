package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/models/basic"
	"net/http"
)

func (r *Registry) GetBasics(c *gin.Context) {
	ds := basic.NewDataStore(r.db)
	basic, err := ds.Get()

	if err != nil {
		presenters.ViewInternalServerError(c, err)
	}

	presenters.JSON(c, http.StatusOK, basic)
}
