package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/plugins/hotpepper"
	"github.com/yuuis/PersonalDataRepository/models/small_category"
	"time"
)

func (r *Registry) GetIntake(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := hotpepper.NewDataStore(r.db)

	i, err := ds.All()

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.IntakesView(ctx, *i)
	}
}

func (r *Registry) CreateIntake(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := hotpepper.NewDataStore(r.db)

	var ipt inputIntake
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	_, err := small_category.NewDataStore(r.db).Find(ipt.SmallCategoryID)
	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewBadRequest(ctx, fmt.Errorf("small_category: %v", err))
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		i, err := ds.Store(&hotpepper.Intake{
			ID:              models.GenerateUUID(),
			Menu:            ipt.Menu,
			Calorie:         ipt.Calorie,
			SmallCategoryID: ipt.SmallCategoryID,
			CreatedAt:       time.Now(),
		})

		if err != nil {
			presenters.ViewInternalServerError(ctx, err)
		}

		presenters.IntakeView(ctx, *i)
	}
}

type inputIntake struct {
	Menu            string  `json:"menu"`
	Calorie         float64 `json:"calorie""`
	SmallCategoryID string  `json:"small_category_id"`
}
