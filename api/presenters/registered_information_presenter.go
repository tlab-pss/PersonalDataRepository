package presenters

import (
	"context"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models/registered_information"
	"net/http"
)

func RegisteredInformationView(ctx context.Context, ri registered_information.RegisteredInformation) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, ri)
}
