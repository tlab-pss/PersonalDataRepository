package presenters

import (
	"context"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models/registration_information"
	"net/http"
)

func RegistrationInformationView(ctx context.Context, ri registration_information.RegistrationInformation) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, ri)
}
