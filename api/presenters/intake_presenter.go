package presenters

import (
	"context"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models/plugins/hotpepper"
	"net/http"
)

func IntakeView(ctx context.Context, i hotpepper.Intake) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, i)
}

func IntakesView(ctx context.Context, i []hotpepper.Intake) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, i)
}
