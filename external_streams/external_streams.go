package external_streams

import (
	"api/helpers"
	"api/hms_errors"
	"bytes"
	"net/http"
	"net/url"
	"strconv"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

var externalStreamsBaseUrl = helpers.GetEndpointUrl("external-streams")

type VideoResolution struct {
	Height uint32 `json:"height,omitempty"`
	Width  uint32 `json:"width,omitempty"`
}

type HMSStartExternalStreamBody struct {
	MeetingUrl  string           `json:"meeting_url,omitempty"`
	RTMPUrls    []string         `json:"rtmp_urls"`
	Recording   bool             `json:"recording,omitempty"`
	Destination string           `json:"destination,omitempty"`
	Resolution  *VideoResolution `json:"resolution,omitempty"`
}

type HMSExternalStreamsQueryParam struct {
	RoomId    string `form:"room_id,omitempty"`
	SessionId string `form:"session_id,omitempty"`
	Status    string `form:"status,omitempty"`
	Start     string `form:"start,omitempty"`
	Limit     int32  `form:"limit,omitempty"`
}

// Start an external stream for a room
func StartExternalStream(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingRoomId})
	}

	var rb HMSStartExternalStreamBody
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var videoResolution *VideoResolution

	if rb.Resolution != nil {
		videoResolution = &VideoResolution{
			Height: rb.Resolution.Height,
			Width:  rb.Resolution.Width,
		}
	}

	postBody, _ := json.Marshal(HMSStartExternalStreamBody{
		MeetingUrl:  rb.MeetingUrl,
		RTMPUrls:    rb.RTMPUrls,
		Recording:   rb.Recording,
		Resolution:  videoResolution,
		Destination: rb.Destination,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, externalStreamsBaseUrl+"/room/"+roomId+"/start", "POST", payload)
}

// Stop all external streams in the given room
func StopExternalStreams(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingRoomId})
	}
	helpers.MakeApiRequest(ctx, externalStreamsBaseUrl+"/room/"+roomId+"/stop", "POST", nil)
}

// Stop an external stream given the stream ID
func StopExternalStream(ctx *gin.Context) {
	streamId, ok := ctx.Params.Get("streamId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingStreamId})
	}
	helpers.MakeApiRequest(ctx, externalStreamsBaseUrl+"/"+streamId+"/stop", "POST", nil)
}

// Get an external stream by its ID
func GetExternalStream(ctx *gin.Context) {
	streamId, ok := ctx.Params.Get("streamId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingStreamId})
	}
	helpers.MakeApiRequest(ctx, externalStreamsBaseUrl+"/"+streamId, "GET", nil)
}

// List all external streams
// Applicable filters: room_id string, session_id string, status string, start string, limit int32
func ListExternalStreams(ctx *gin.Context) {

	var param HMSExternalStreamsQueryParam

	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("room_id", param.RoomId)
		qs.Add("session_id", param.SessionId)
		qs.Add("status", param.Status)
		qs.Add("start", param.Start)
		qs.Add("limit", strconv.Itoa(int(param.Limit)))
	}
	helpers.MakeApiRequest(ctx, externalStreamsBaseUrl+"?"+qs.Encode(), "GET", nil)

}
