package presenters

import (
	"context"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models/basic"
	"net/http"
)

func BasicView(ctx context.Context, b basic.Basic) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, b)
}
