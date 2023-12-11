package room

import (
	"bytes"
	"encoding/json"

	"net/http"

	"api/helpers"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Name               string          `json:"name,omitempty"`
	Description        string          `json:"description,omitempty"`
	TemplateId         string          `json:"template_id,omitempty"`
	RecordingInfo      *RECORDING_INFO `json:"recording_info,omitempty"`
	Region             string          `json:"region,omitempty"`
	LargeRoom          bool            `json:"large_room,omitempty"`
	Size               int             `json:"size,omitempty"`
	MaxDurationSeconds string          `json:"max_duration_seconds,omitempty"`
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

// Create a   room with a given room name
func CreateRoom(ctx *gin.Context) {

	var rb RequestBody
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

	requestBody := RequestBody{
		Name:               rb.Name,
		Description:        rb.Description,
		TemplateId:         rb.TemplateId,
		RecordingInfo:      recordingInfo,
		Region:             rb.Region,
		LargeRoom:          rb.LargeRoom,
		Size:               rb.Size,
		MaxDurationSeconds: rb.MaxDurationSeconds,
	}

	postBody, _ := json.Marshal(requestBody)

	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, "rooms", "POST", payload)

}

// Get details of a given room
func GetRoom(ctx *gin.Context) {

	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "roomId is missing"})
	}

	endpointPath := "rooms/" + roomId
	helpers.MakeApiRequest(ctx, endpointPath, "GET", nil)

}
