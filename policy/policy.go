package policy

import (
	"api/helpers"
	"api/hmserrors"
	"api/livestreams"
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

var policyBaseUrl = helpers.GetEndpointUrl("templates")

type HMSAudio struct {
	Bitrate uint16 `json:"bitRate,omitempty"`
	Codec   string `json:"codec,omitempty"`
}

type HMSVideo struct {
	Bitrate   uint   `json:"bitRate,omitempty"`
	Codec     string `json:"codec,omitempty"`
	Framerate uint8  `json:"frameRate,omitempty"`
	Height    uint16 `json:"height,omitempty"`
	Width     uint16 `json:"width,omitempty"`
}

type HMSScreen = HMSVideo

type HMSSimulcastLayer struct {
	MaxBitrate            uint    `json:"maxBitrate,omitempty"`
	MaxFramerate          uint8   `json:"maxFramerate"`
	ScaleResolutionDownBy float32 `json:"scaleResolutionDownBy"`
	Rid                   string  `json:"rid"`
}

type HMSSimulcast struct {
	Layers *[]HMSSimulcastLayer `json:"layers,omitempty"`
}

type HMSPublishParams struct {
	Allowed   []string                 `json:"allowed,omitempty"`
	Audio     *HMSAudio                `json:"audio,omitempty"`
	Video     *HMSVideo                `json:"video,omitempty"`
	Screen    *HMSScreen               `json:"screen,omitempty"`
	Simulcast map[string]*HMSSimulcast `json:"simulcast,omitempty"`
}

type HMSSubscribeDegradation struct {
	PacketLossThreshold       uint8 `json:"packetLossThreshold,omitempty"`
	DegradeGracePeriodSeconds uint8 `json:"degradeGracePeriodSeconds,omitempty"`
	RecoverGracePeriodSeconds uint8 `json:"recoverGracePeriodSeconds,omitempty"`
}

type HMSSubscribeParams struct {
	MaxSubsBitRate       int                      `json:"maxSubsBitRate,omitempty"`
	SubscribeToRoles     []string                 `json:"subscribeToRoles,omitempty"`
	SubscribeDegradation *HMSSubscribeDegradation `json:"subscribeDegradation,omitempty"`
}

type HMSPermissions struct {
	EndRoom          bool `json:"endRoom,omitempty"`
	RemoveOthers     bool `json:"removeOthers,omitempty"`
	Mute             bool `json:"mute,omitempty"`
	Unmute           bool `json:"unmute,omitempty"`
	ChangeRole       bool `json:"changeRole,omitempty"`
	SendRoomState    bool `json:"sendRoomState,omitempty"`
	PollRead         bool `json:"pollRead,omitempty"`
	PollWrite        bool `json:"pollWrite,omitempty"`
	BrowserRecording bool `json:"browserRecording,omitempty"`
	RtmpStreaming    bool `json:"rtmpStreaming,omitempty"`
	HlsStreaming     bool `json:"hlsStreaming,omitempty"`
}

type HMSRole struct {
	Name            string              `json:"name,omitempty"`
	PublishParams   *HMSPublishParams   `json:"publishParams,omitempty"`
	SubscribeParams *HMSSubscribeParams `json:"subscribeParams,omitempty"`
	Permissions     *HMSPermissions     `json:"permissions,omitempty"`
	Priority        uint8               `json:"priority,omitempty"`
	MaxPeerCount    int                 `json:"maxPeerCount,omitempty"`
}

type UploadOptions struct {
	Region string `json:"region,omitempty"`
}

type UploadCredentials struct {
	Key       string `json:"key"`
	SecretKey string `json:"secretKey"`
}

type HMSUpload struct {
	Type        string             `json:"type"`
	Location    string             `json:"location"`
	Prefix      string             `json:"prefix"`
	Options     *UploadOptions     `json:"options,omitempty"`
	Credentials *UploadCredentials `json:"credentials,omitempty"`
}

type HMSRecording struct {
	Enabled bool       `json:"enabled,omitempty"`
	Upload  *HMSUpload `json:"upload,omitempty"`
}

type HMSRoomState struct {
	MessageInterval     uint16 `json:"messageInterval,omitempty"`
	SendPeerList        bool   `json:"sendPeerList,omitempty"`
	Enabled             bool   `json:"enabled,omitempty"`
	StopRoomStateOnJoin bool   `json:"stopRoomStateOnJoin,omitempty"`
}

type HMSSetting struct {
	Region    string        `json:"region,omitempty"`
	Recording *HMSRecording `json:"recording,omitempty"`
	RoomState *HMSRoomState `json:"roomState,omitempty"`
}

type HMSBrowserRecordings struct {
	Name            string         `json:"name"`
	Width           uint16         `json:"width,omitempty"`
	Height          uint16         `json:"height,omitempty"`
	MaxDuration     uint16         `json:"maxDuration,omitempty"`
	Thumbnails      *HLSThumbnails `json:"thumbnails,omitempty"`
	PresignDuration uint           `json:"presignDuration,omitempty"`
	Role            string         `json:"role"`
	AutoStart       bool           `json:"autoStart,omitempty"`
	AutoStopTimeout uint           `json:"autoStopTimeout,omitempty"`
}

type HMSRtmpDestinations struct {
	Name             string   `json:"name"`
	Width            uint16   `json:"width,omitempty"`
	Height           uint16   `json:"height,omitempty"`
	MaxDuration      uint16   `json:"maxDuration,omitempty"`
	RtmpUrls         []string `json:"rtmpUrls,omitempty"`
	RecordingEnabled bool     `json:"recordingEnabled,omitempty"`
	AutoStopTimeout  uint     `json:"autoStopTimeout,omitempty"`
}

type HLSDestinationLayers struct {
	Width        int `json:"width,omitempty"`
	Height       int `json:"height,omitempty"`
	VideoBitrate int `json:"videoBitrate,omitempty"`
	AudioBitrate int `json:"audioBitrate,omitempty"`
}

type HLSThumbnails struct {
	Enabled bool  `json:"enabled,omitempty"`
	Width   int   `json:"width,omitempty"`
	Height  int   `json:"height,omitempty"`
	Fps     int   `json:"fps,omitempty"`
	Offsets []int `json:"offsets,omitempty"`
}

type HLSRecording struct {
	HlsVod             bool                    `json:"hlsVod,omitempty"`
	SingleFilePerLayer bool                    `json:"singleFilePerLayer,omitempty"`
	EnableZipUpload    bool                    `json:"enableZipUpload,omitempty"`
	Layers             *[]HLSDestinationLayers `json:"layers,omitempty"`
	Thumbnails         *HLSThumbnails          `json:"thumbnails,omitempty"`
	PresignDuration    int                     `json:"presignDuration,omitempty"`
}

type HMSHlsDestinations struct {
	Name                    string                  `json:"name"`
	MaxDuration             uint16                  `json:"maxDuration,omitempty"`
	Layers                  *[]HLSDestinationLayers `json:"layers,omitempty"`
	PlaylistType            string                  `json:"playlistType,omitempty"`
	NumPlaylistSegments     uint                    `json:"numPlaylistSegments,omitempty"`
	VideoFrameRate          uint16                  `json:"videoFrameRate,omitempty"`
	EnableMetadataInsertion bool                    `json:"enableMetadataInsertion,omitempty"`
	EnableStaticUrl         bool                    `json:"enableStaticUrl,omitempty"`
	Recording               *HLSRecording           `json:"recording,omitempty"`
	AutoStopTimeout         uint                    `json:"autoStopTimeout,omitempty"`
}
type HMSTranscriptions struct {
	Name             string                            `json:"name"`
	Modes            []string                          `json:"modes,omitempty"`
	Role             string                            `json:"role"`
	OutputModes      []string                          `json:"outputModes,omitempty"`
	CustomVocabulary []string                          `json:"customVocabulary,omitempty"`
	Summary          *livestreams.TranscriptionSummary `json:"summary,omitempty"`
}

type HMSDestination struct {
	BrowserRecordings map[string]*HMSBrowserRecordings `json:"browserRecordings,omitempty"`
	RtmpDestinations  map[string]*HMSRtmpDestinations  `json:"rtmpDestinations,omitempty"`
	HlsDestinations   map[string]*HMSHlsDestinations   `json:"hlsDestinations,omitempty"`
	Transcriptions    map[string]*HMSTranscriptions    `json:"transcriptions,omitempty"`
}

type HMSTemplate struct {
	Name         string              `json:"name,omitempty"`
	Roles        map[string]*HMSRole `json:"roles,omitempty"`
	Settings     *HMSSetting         `json:"settings,omitempty"`
	Destinations *HMSDestination     `json:"destinations,omitempty"`
}

type HMSTemplateQueryParam struct {
	Limit uint8  `form:"limit,omitempty"`
	Start string `form:"start,omitempty"`
}

// Get the post request body
func getTemplateRequestBody(ctx *gin.Context) *bytes.Buffer {
	var rb HMSTemplate
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSTemplate{
		Name:         rb.Name,
		Roles:        rb.Roles,
		Settings:     rb.Settings,
		Destinations: rb.Destinations,
	})

	payload := bytes.NewBuffer(postBody)
	return payload
}

// Create a template
func CreateTemplate(ctx *gin.Context) {
	payload := getTemplateRequestBody(ctx)
	helpers.MakeApiRequest(ctx, policyBaseUrl, "POST", payload)
}

// Update a template using the template ID
func UpdateTemplate(ctx *gin.Context) {
	templateId, ok := ctx.Params.Get("templateId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingTemplateId})
	}

	payload := getTemplateRequestBody(ctx)
	helpers.MakeApiRequest(ctx, policyBaseUrl+"/"+templateId, "POST", payload)
}

// Get a list of all rooms
// Applicable filters: start string, limit int
func ListTemplates(ctx *gin.Context) {
	var param HMSTemplateQueryParam
	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("start", param.Start)
		if param.Limit >= 10 {
			qs.Add("limit", strconv.Itoa(int(param.Limit)))
		}
	}
	helpers.MakeApiRequest(ctx, policyBaseUrl+"?"+qs.Encode(), "GET", nil)
}

// Get a template using the template ID
func GetTemplate(ctx *gin.Context) {
	templateId, ok := ctx.Params.Get("templateId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingTemplateId})
	}
	helpers.MakeApiRequest(ctx, policyBaseUrl+"/"+templateId, "GET", nil)
}

// Modify a role in a template
func ModifyTemplateRole(ctx *gin.Context) {
	templateId, ok := ctx.Params.Get("templateId")
	roleName, ok1 := ctx.Params.Get("roleName")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingTemplateIdAndRoleName})
	}

	var rb HMSRole
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSRole{
		Name:            rb.Name,
		PublishParams:   rb.PublishParams,
		SubscribeParams: rb.SubscribeParams,
		Permissions:     rb.Permissions,
		Priority:        rb.Priority,
		MaxPeerCount:    rb.MaxPeerCount,
	})

	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, policyBaseUrl+"/"+templateId+"/roles/"+roleName, "POST", payload)
}

// Retrieve a specific role
func GetTemplateRole(ctx *gin.Context) {
	templateId, ok := ctx.Params.Get("templateId")
	roleName, ok1 := ctx.Params.Get("roleName")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingTemplateIdAndRoleName})
	}
	helpers.MakeApiRequest(ctx, policyBaseUrl+"/"+templateId+"/roles/"+roleName, "GET", nil)
}

// Delete a specific role
func DeleteTemplateRole(ctx *gin.Context) {
	templateId, ok := ctx.Params.Get("templateId")
	roleName, ok1 := ctx.Params.Get("roleName")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingTemplateIdAndRoleName})
	}
	helpers.MakeApiRequest(ctx, policyBaseUrl+"/"+templateId+"/roles/"+roleName, "DELETE", nil)
}

// Retrieve template settings
func GetTemplateSettings(ctx *gin.Context) {
	templateId, ok := ctx.Params.Get("templateId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingTemplateId})
	}
	helpers.MakeApiRequest(ctx, policyBaseUrl+"/"+templateId+"/settings", "GET", nil)
}

// Update template settings
func UpdateTemplateSettings(ctx *gin.Context) {
	templateId, ok := ctx.Params.Get("templateId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingTemplateId})
	}

	var rb HMSSetting
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	postBody, _ := json.Marshal(HMSSetting{
		Region:    rb.Region,
		Recording: rb.Recording,
		RoomState: rb.RoomState,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, policyBaseUrl+"/"+templateId+"/settings", "POST", payload)
}

// Retrieve template destinations
func GetTemplateDestinations(ctx *gin.Context) {
	templateId, ok := ctx.Params.Get("templateId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingTemplateId})
	}
	helpers.MakeApiRequest(ctx, policyBaseUrl+"/"+templateId+"/destinations", "GET", nil)
}

// Update template destinations
func UpdateTemplateDestinations(ctx *gin.Context) {
	templateId, ok := ctx.Params.Get("templateId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingTemplateId})
	}

	var rb HMSDestination
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	postBody, _ := json.Marshal(HMSDestination{
		BrowserRecordings: rb.BrowserRecordings,
		RtmpDestinations:  rb.RtmpDestinations,
		HlsDestinations:   rb.HlsDestinations,
		Transcriptions:    rb.Transcriptions,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, policyBaseUrl+"/"+templateId+"/destinations", "POST", payload)
}
