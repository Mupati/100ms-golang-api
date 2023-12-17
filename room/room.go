package room

import (
	"bytes"
	"encoding/json"
	"os"
	"strconv"

	"net/http"
	"net/url"

	"api/helpers"

	"github.com/gin-gonic/gin"
)

type HMSRequestBody struct {
	Name               string          `json:"name,omitempty"`
	Description        string          `json:"description,omitempty"`
	TemplateId         string          `json:"template_id,omitempty"`
	RecordingInfo      *RECORDING_INFO `json:"recording_info,omitempty"`
	Region             string          `json:"region,omitempty"`
	LargeRoom          bool            `json:"large_room,omitempty"`
	Size               int             `json:"size,omitempty"`
	MaxDurationSeconds string          `json:"max_duration_seconds,omitempty"`
	Polls              []string        `json:"polls,omitempty"`
}

type RECORDING_INFO struct {
	Enabled    bool         `json:"enabled,omitempty"`
	Polls      []string     `json:"polls,omitempty"`
	UploadInfo *UPLOAD_INFO `json:"upload_info,omitempty"`
}

type UPLOAD_INFO struct {
	UploadType  string              `json:"type,omitempty"`
	Location    string              `json:"location,omitempty"`
	Prefix      string              `json:"prefix,omitempty"`
	Options     *UPLOAD_OPTIONS     `json:"options,omitempty"`
	Credentials *UPLOAD_CREDENTIALS `json:"credentials,omitempty"`
}

type UPLOAD_OPTIONS struct {
	Region string `json:"region,omitempty"`
}

type UPLOAD_CREDENTIALS struct {
	Key    string `json:"key,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type HMSRoomQueryParam struct {
	Name    string `form:"name,omitempty"`
	Enabled *bool  `form:"enabled,omitempty"`
	Before  string `form:"before,omitempty"`
	After   string `form:"after,omitempty"`
}

const MISSING_ROOM_ID_ERROR_MESSAGE = "provide a room ID"

var roomBaseUrl = os.Getenv("BASE_URL") + "rooms"

// Get the post request body
func getRequestBody(ctx *gin.Context) *bytes.Buffer {
	var rb HMSRequestBody
	var recordingInfo *RECORDING_INFO
	var uploadInfo *UPLOAD_INFO
	var uploadOptions *UPLOAD_OPTIONS
	var uploadCredentials *UPLOAD_CREDENTIALS

	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if rb.RecordingInfo != nil {

		if rb.RecordingInfo.UploadInfo != nil {

			if rb.RecordingInfo.UploadInfo.Options != nil {
				uploadOptions = &UPLOAD_OPTIONS{
					Region: rb.RecordingInfo.UploadInfo.Options.Region,
				}
			}

			if rb.RecordingInfo.UploadInfo.Credentials != nil {
				uploadCredentials = &UPLOAD_CREDENTIALS{
					Key:    rb.RecordingInfo.UploadInfo.Credentials.Key,
					Secret: rb.RecordingInfo.UploadInfo.Credentials.Secret,
				}
			}

			uploadInfo = &UPLOAD_INFO{
				UploadType:  rb.RecordingInfo.UploadInfo.UploadType,
				Location:    rb.RecordingInfo.UploadInfo.Location,
				Prefix:      rb.RecordingInfo.UploadInfo.Prefix,
				Options:     uploadOptions,
				Credentials: uploadCredentials,
			}
		}

		recordingInfo = &RECORDING_INFO{
			Enabled:    rb.RecordingInfo.Enabled,
			Polls:      rb.RecordingInfo.Polls,
			UploadInfo: uploadInfo,
		}
	}

	postBody, _ := json.Marshal(HMSRequestBody{
		Name:               rb.Name,
		Description:        rb.Description,
		TemplateId:         rb.TemplateId,
		RecordingInfo:      recordingInfo,
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
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
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
		qs.Add("enabled", strconv.FormatBool(*param.Enabled))
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
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}

	payload := getRequestBody(ctx)
	helpers.MakeApiRequest(ctx, roomBaseUrl+"/"+roomId, "POST", payload)
}

// Enable a room
func EnableRoom(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}
	postBody, _ := json.Marshal(map[string]bool{"enabled": true})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, roomBaseUrl+"/"+roomId, "POST", payload)
}

// Disable a room
func DisableRoom(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}
	postBody, _ := json.Marshal(map[string]bool{"enabled": false})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, roomBaseUrl+"/"+roomId, "POST", payload)
}
