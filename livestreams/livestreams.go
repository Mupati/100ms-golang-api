package livestreams

import (
	"api/helpers"
	"api/hmserrors"
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

var liveStreamsBaseUrl = helpers.GetEndpointUrl("live-streams")

type TranscriptionSummarySection struct {
	Title  string `json:"title"`
	Format string `json:"format"`
}

type TranscriptionSummary struct {
	Enabled     bool                           `json:"enabled,omitempty"`
	Context     string                         `json:"context,omitempty"`
	Temperature float64                        `json:"temperature,omitempty"`
	Sections    *[]TranscriptionSummarySection `json:"sections,omitempty"`
}

type Transcription struct {
	Enabled          bool                  `json:"enabled,omitempty"`
	Modes            []string              `json:"modes,omitempty"`
	OutputModes      []string              `json:"output_modes,omitempty"`
	CustomVocabulary []string              `json:"custom_vocabulary,omitempty"`
	Summary          *TranscriptionSummary `json:"summary,omitempty"`
}

type Recording struct {
	HLSVod             bool `json:"hls_vod,omitempty"`
	SingleFilePerLayer bool `json:"single_file_per_layer,omitempty"`
}

type HMSLivestream struct {
	MeetingUrl    string         `json:"meeting_url,omitempty"`
	Destination   string         `json:"destination,omitempty"`
	Recording     *Recording     `json:"recording,omitempty"`
	Transcription *Transcription `json:"transcription,omitempty"`
}

type HMSLiveStreamsQueryParam struct {
	RoomId    string `form:"room_id,omitempty"`
	SessionId string `form:"session_id,omitempty"`
	Status    string `form:"status,omitempty"`
	Start     string `form:"start,omitempty"`
	Limit     int32  `form:"limit,omitempty"`
}

type TimedMetaDataBody struct {
	Payload  string `json:"payload"`
	Duration int32  `json:"duration"`
}

// Start a live stream for a room
func StartLiveStream(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
	}

	var rb HMSLivestream
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSLivestream{
		MeetingUrl:    rb.MeetingUrl,
		Recording:     rb.Recording,
		Destination:   rb.Destination,
		Transcription: rb.Transcription,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, liveStreamsBaseUrl+"/room/"+roomId+"/start", "POST", payload)
}

// Stop all live stream in the given room
func StopLiveStreams(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
	}
	helpers.MakeApiRequest(ctx, liveStreamsBaseUrl+"/room/"+roomId+"/stop", "POST", nil)
}

// Stop a livestream given the stream ID
func StopLiveStream(ctx *gin.Context) {
	streamId, ok := ctx.Params.Get("streamId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingStreamId})
	}
	helpers.MakeApiRequest(ctx, liveStreamsBaseUrl+"/"+streamId+"/stop", "POST", nil)
}

// Get a livestream by its ID
func GetLiveStream(ctx *gin.Context) {
	streamId, ok := ctx.Params.Get("streamId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingStreamId})
	}
	helpers.MakeApiRequest(ctx, liveStreamsBaseUrl+"/"+streamId, "GET", nil)
}

// List all livestreams
// Applicable filters: room_id string, session_id string, status string, start string, limit int32
func ListLiveStreams(ctx *gin.Context) {
	var param HMSLiveStreamsQueryParam
	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("room_id", param.RoomId)
		qs.Add("session_id", param.SessionId)
		qs.Add("status", param.Status)
		qs.Add("start", param.Start)
		qs.Add("limit", strconv.Itoa(int(param.Limit)))
	}
	helpers.MakeApiRequest(ctx, liveStreamsBaseUrl+"?"+qs.Encode(), "GET", nil)

}

// Send timed-metadata during a livestream
func SendTimedMetada(ctx *gin.Context) {
	streamId, ok := ctx.Params.Get("streamId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingStreamId})
	}

	var rb TimedMetaDataBody
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(TimedMetaDataBody{
		Payload:  rb.Payload,
		Duration: rb.Duration,
	})

	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, liveStreamsBaseUrl+"/"+streamId+"/timed-metadata", "POST", payload)
}

// Pause a livestream recording
func PauseLiveStreamRecording(ctx *gin.Context) {
	streamId, ok := ctx.Params.Get("streamId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingStreamId})
	}
	helpers.MakeApiRequest(ctx, liveStreamsBaseUrl+"/"+streamId+"/pause-recording", "POST", nil)
}

// Resuming a livestream recording
func ResumeLiveStreamRecording(ctx *gin.Context) {
	streamId, ok := ctx.Params.Get("streamId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingStreamId})
	}
	helpers.MakeApiRequest(ctx, liveStreamsBaseUrl+"/"+streamId+"/resume-recording", "POST", nil)
}
