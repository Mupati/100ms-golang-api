package sessions

import (
	"api/helpers"
	"api/hms_errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sessionsBaseUrl = helpers.GetEndpointUrl("sessions")

type HMSSessionQueryParam struct {
	RoomId string `form:"room_id,omitempty"`
	Active *bool  `form:"active,omitempty"`
	Before string `form:"before,omitempty"`
	After  string `form:"after,omitempty"`
}

// Get a session's details
func GetSession(ctx *gin.Context) {
	sessionId, ok := ctx.Params.Get("sessionId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingSessionId})
	}

	helpers.MakeApiRequest(ctx, sessionsBaseUrl+"/"+sessionId, "GET", nil)
}

// List all sessions
// Applicable filters: room_id string, active *bool, after string, before string
func ListSessions(ctx *gin.Context) {
	var param HMSSessionQueryParam

	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("room_id", param.RoomId)
		if param.Active != nil {
			qs.Add("active", strconv.FormatBool(*param.Active))
		}
		qs.Add("before", param.Before)
		qs.Add("after", param.After)
	}

	helpers.MakeApiRequest(ctx, sessionsBaseUrl+"?"+qs.Encode(), "GET", nil)
}
