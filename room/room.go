package room

import (
	"bytes"
	"encoding/json"
	"strconv"

	"net/http"
	"net/url"

	"api/helpers"
	"api/hms_errors"

	"github.com/gin-gonic/gin"
)

type RecordingInfo struct {
	Enabled    bool        `json:"enabled,omitempty"`
	Polls      []string    `json:"polls,omitempty"`
	UploadInfo *UploadInfo `json:"upload_info,omitempty"`
}

type UploadInfo struct {
	Type        string             `json:"type"`
	Location    string             `json:"location"`
	Prefix      string             `json:"prefix,omitempty"`
	Options     *UploadOptions     `json:"options,omitempty"`
	Credentials *UploadCredentials `json:"credentials,omitempty"`
}

type UploadOptions struct {
	Region string `json:"region,omitempty"`
}

type UploadCredentials struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type HMSRoom struct {
	Name               string         `json:"name,omitempty"`
	Description        string         `json:"description,omitempty"`
	TemplateId         string         `json:"template_id,omitempty"`
	RecordingInfo      *RecordingInfo `json:"recording_info,omitempty"`
	Region             string         `json:"region,omitempty"`
	LargeRoom          bool           `json:"large_room,omitempty"`
	Size               int            `json:"size,omitempty"`
	MaxDurationSeconds string         `json:"max_duration_seconds,omitempty"`
	Polls              []string       `json:"polls,omitempty"`
}
type HMSRoomQueryParam struct {
	Name    string `form:"name,omitempty"`
	Enabled *bool  `form:"enabled,omitempty"`
	Before  string `form:"before,omitempty"`
	After   string `form:"after,omitempty"`
}

var roomBaseUrl = helpers.GetEndpointUrl("rooms")

// Get the post request body
func getRequestBody(ctx *gin.Context) *bytes.Buffer {
	var rb HMSRoom

	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSRoom{
		Name:               rb.Name,
		Description:        rb.Description,
		TemplateId:         rb.TemplateId,
		RecordingInfo:      rb.RecordingInfo,
		Region:             rb.Region,
		LargeRoom:          rb.LargeRoom,
		Size:               rb.Size,
		MaxDurationSeconds: rb.MaxDurationSeconds,
		Polls:              rb.Polls,
	})
	payload := bytes.NewBuffer(postBody)
	return payload
}

// Get details of a given room
func GetRoom(ctx *gin.Context) {

	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingRoomId})
	}

	helpers.MakeApiRequest(ctx, roomBaseUrl+"/"+roomId, "GET", nil)

}

// Get a list of all rooms
// Applicable filters: name string, enabled *bool, after string, before string
func ListRooms(ctx *gin.Context) {
	var param HMSRoomQueryParam
	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("name", param.Name)
		if param.Enabled != nil {
			qs.Add("enabled", strconv.FormatBool(*param.Enabled))
		}
		qs.Add("before", param.Before)
		qs.Add("after", param.After)
	}
	helpers.MakeApiRequest(ctx, roomBaseUrl+"?"+qs.Encode(), "GET", nil)
}

// Create a   room with a given room name
func CreateRoom(ctx *gin.Context) {
	payload := getRequestBody(ctx)
	helpers.MakeApiRequest(ctx, roomBaseUrl, "POST", payload)
}

// Update a Room
func UpdateRoom(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingRoomId})
	}

	payload := getRequestBody(ctx)
	helpers.MakeApiRequest(ctx, roomBaseUrl+"/"+roomId, "POST", payload)
}

// Enable a room
func EnableRoom(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingRoomId})
	}
	postBody, _ := json.Marshal(map[string]bool{"enabled": true})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, roomBaseUrl+"/"+roomId, "POST", payload)
}

// Disable a room
func DisableRoom(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingRoomId})
	}
	postBody, _ := json.Marshal(map[string]bool{"enabled": false})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, roomBaseUrl+"/"+roomId, "POST", payload)
}
