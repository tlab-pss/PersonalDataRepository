package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/small_category"
	"github.com/yuuis/PersonalDataRepository/models/user_like"
	"time"
)

func (r *Registry) GetUserLike(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := user_like.NewDataStore(r.db)

	ps, err := ds.All()

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.UserLikesView(ctx, ps)
	}
}

func (r *Registry) CreateUserLike(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := user_like.NewDataStore(r.db)

	var ipt inputUserLike
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	_, err := small_category.NewDataStore(r.db).Find(ipt.SmallCategoryId)
	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewBadRequest(ctx, fmt.Errorf("small_category: %v", err))
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		ul, err := ds.Store(&user_like.UserLike{
			ID:              models.GenerateUUID(),
			SmallCategoryId: ipt.SmallCategoryId,
			CreatedAt:       time.Now(),
		})

		if err != nil {
			presenters.ViewInternalServerError(ctx, err)
		}

		presenters.UserLikeView(ctx, *ul)
	}
}

type inputUserLike struct {
	SmallCategoryId string `json:"small_category_id"`
}
