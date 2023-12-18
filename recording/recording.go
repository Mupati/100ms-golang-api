package recording

import (
	"api/helpers"
	"bytes"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

var recordingBaseUrl = os.Getenv("BASE_URL") + "recordings"

const missingRoomIdErrorMessage = "provide a room ID"
const missingRecordingIdErrorMessage = "provide the recording ID"

type RecordingResolution struct {
	Height uint32 `json:"height,omitempty"`
	Width  uint32 `json:"width,omitempty"`
}

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

type RecordingTranscription struct {
	Enabled          bool                  `json:"enabled,omitempty"`
	OutputModes      []string              `json:"output_modes,omitempty"`
	CustomVocabulary []string              `json:"custom_vocabulary,omitempty"`
	Summary          *TranscriptionSummary `json:"summary,omitempty"`
}

type HMSStartRecordingBody struct {
	MeetingUrl    string                  `json:"meeting_url,omitempty"`
	Destination   string                  `json:"destination,omitempty"`
	AudioOnly     bool                    `json:"audio_only,omitempty"`
	Resolution    *RecordingResolution    `json:"resolution,omitempty"`
	Transcription *RecordingTranscription `json:"transcription,omitempty"`
}

// Start a recording
func StartRecording(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": missingRoomIdErrorMessage})
	}

	var rb HMSStartRecordingBody
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSStartRecordingBody{
		MeetingUrl:    rb.MeetingUrl,
		Destination:   rb.Destination,
		AudioOnly:     rb.AudioOnly,
		Resolution:    rb.Resolution,
		Transcription: rb.Transcription,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, recordingBaseUrl+"/room/"+roomId+"/start", "POST", payload)
}

// Stop all recordings in the given room
func StopRecordings(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": missingRoomIdErrorMessage})
	}
	helpers.MakeApiRequest(ctx, recordingBaseUrl+"/room/"+roomId+"/stop", "POST", nil)
}

// Stop a recording given the recording ID
func StopRecording(ctx *gin.Context) {
	recordingId, ok := ctx.Params.Get("recordingId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "provide the recording ID"})
	}
	helpers.MakeApiRequest(ctx, recordingBaseUrl+"/"+recordingId+"/stop", "POST", nil)
}

// Get a recording by its ID
func GetRecording(ctx *gin.Context) {
	recordingId, ok := ctx.Params.Get("recordingId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": missingRecordingIdErrorMessage})
	}
	helpers.MakeApiRequest(ctx, recordingBaseUrl+"/"+recordingId, "GET", nil)
}

// List all recordings in the room.
func ListRecordings(ctx *gin.Context) {
	helpers.MakeApiRequest(ctx, recordingBaseUrl, "GET", nil)
}

// Get the configuration of a recording
func GetRecordingConfig(ctx *gin.Context) {
	recordingId, ok := ctx.Params.Get("recordingId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": missingRecordingIdErrorMessage})
	}
	helpers.MakeApiRequest(ctx, recordingBaseUrl+"/"+recordingId+"/config", "GET", nil)
}
