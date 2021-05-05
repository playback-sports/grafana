package api

import (
	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/infra/metrics"
	"github.com/grafana/grafana/pkg/models"
)

func (hs *HTTPServer) GetDashboardSchema(c *models.ReqContext) response.Response {
	var dsSchema *simplejson.Json
	var err error

	dsSchema, err = hs.LoadSchemaService.GetDashboardSchema()
	if err != nil {
		return response.Error(500, "Error while trim default value from dashboard json", err)
	}

	dto := dtos.TrimDashboardFullWithMeta{
		Dashboard: dsSchema,
	}

	c.TimeRequest(metrics.MApiDashboardGet)
	return response.JSON(200, dto)
}
