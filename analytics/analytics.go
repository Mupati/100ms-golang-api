package analytics

import (
	"api/helpers"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

var streamKeysBaseUrl = os.Getenv("BASE_URL") + "analytics"

type HMSAnalyticsQueryParam struct {
	Type string `form:"type,omitempty"`
}

// Get analytics events
func GetAnalyticsEvents(ctx *gin.Context) {
	var param HMSAnalyticsQueryParam
	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("type", param.Type)
	}
	helpers.MakeApiRequest(ctx, streamKeysBaseUrl+"/events"+"?"+qs.Encode(), "GET", nil)
}
