package presenters

import (
	"context"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models/user_like"
	"net/http"
)

func UserLikeView(ctx context.Context, u user_like.UserLike) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, u)
}

func UserLikesView(ctx context.Context, u *[]user_like.UserLike) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, u)
}
