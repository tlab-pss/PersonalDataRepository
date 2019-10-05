package presenters

import (
	"context"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models/plugin_service"
	"net/http"
)

func PluginServiceView(ctx context.Context, p plugin_service.PluginService) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, p)
}

func PluginServicesView(ctx context.Context, p *[]plugin_service.PluginService) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, p)
}
