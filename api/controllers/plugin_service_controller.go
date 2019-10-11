package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/big_category"
	"github.com/yuuis/PersonalDataRepository/models/plugin_service"
	"time"
)

func (r *Registry) GetPluginService(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := plugin_service.NewDataStore(r.db)

	ps, err := ds.All()

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.PluginServicesView(ctx, ps)
	}
}

func (r *Registry) CreatePluginService(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := plugin_service.NewDataStore(r.db)

	var ipt inputPluginService
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	_, err := big_category.NewDataStore(r.db).Find(ipt.BigCategoryId)
	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewBadRequest(ctx, fmt.Errorf("big_category: %v", err))
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		ps, err := ds.Store(&plugin_service.PluginService{
			ID:            models.GenerateUUID(),
			Name:          ipt.Name,
			BigCategoryId: ipt.BigCategoryId,
			CreatedAt:     time.Now(),
		})

		if err != nil {
			presenters.ViewInternalServerError(ctx, err)
		}

		presenters.PluginServiceView(ctx, *ps)
	}
}

type inputPluginService struct {
	Name          string `json:"name"`
	BigCategoryId string `json:"big_category_id"`
}
