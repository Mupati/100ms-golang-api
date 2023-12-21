package analytics

import (
	"api/helpers"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

var streamKeysBaseUrl = helpers.GetEndpointUrl("analytics")

type HMSAnalyticsQueryParam struct {
	Type      string `form:"type"`
	RoomId    string `form:"room_id"`
	SessionId string `form:"session_id"`
	PeerId    string `form:"peer_id"`
	UserId    string `form:"user_id"`
	Limit     int32  `form:"limit"`
	Start     string `form:"start"`
}

// Get analytics events
func GetAnalyticsEvents(ctx *gin.Context) {
	var param HMSAnalyticsQueryParam
	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("type", param.Type)
		qs.Add("room_id", param.RoomId)
		qs.Add("session_id", param.SessionId)
		qs.Add("peer_id", param.PeerId)
		qs.Add("user_id", param.UserId)
		qs.Add("start", param.Start)
		qs.Add("limit", strconv.Itoa(int(param.Limit)))
	}
	helpers.MakeApiRequest(ctx, streamKeysBaseUrl+"/events"+"?"+qs.Encode(), "GET", nil)
}
