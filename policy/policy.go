package policy

import (
	"api/helpers"
	"api/live_streams"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var policyBaseUrl = os.Getenv("BASE_URL") + "templates"

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

type HMSSimulcastLayers struct {
	MaxBitrate            uint    `json:"maxBitrate,omitempty"`
	MaxFramerate          uint8   `json:"maxFramerate"`
	ScaleResolutionDownBy float32 `json:"scaleResolutionDownBy"`
	Rid                   string  `json:"rid"`
}

type HMSSimulcast struct {
	Layers *HMSSimulcastLayers `json:"layers,omitempty"`
}

type HMSPublishParams struct {
	Allowed []string   `json:"allowed,omitempty"`
	Audio   *HMSAudio  `json:"audio,omitempty"`
	Video   *HMSVideo  `json:"video,omitempty"`
	Screen  *HMSScreen `json:"screen,omitempty"`
}

type HMSSubscribeDegradation struct {
	MaxFramerate              uint8 `json:"packetLossThreshold,omitempty"`
	DegradeGracePeriodSeconds uint8 `json:"degradeGracePeriodSeconds,omitempty"`
	RecoverGracePeriodSeconds uint8 `json:"recoverGracePeriodSeconds,omitempty"`
}

type HMSSubscribeParams struct {
	MaxBitrate           int                      `json:"maxSubsBitRate,omitempty"`
	SubscribeToRoles     *[]HMSRole               `json:"subscribeToRoles,omitempty"`
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
	BrowserRecording bool `json:"browserRecording"`
	RtmpStreaming    bool `json:"rtmpStreaming"`
	HlsStreaming     bool `json:"hlsStreaming"`
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
	MessageInterval uint16 `json:"messageInterval,omitempty"`
	SendPeerList    bool   `json:"sendPeerList,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
}

type HMSSetting struct {
	Region    string        `json:"region,omitempty"`
	Recording *HMSRecording `json:"recording,omitempty"`
	RoomState *HMSRoomState `json:"roomState,omitempty"`
}

type HMSBrowserRecordingsThumbnails struct {
}

type HMSBrowserRecordings struct {
	Name            string                          `json:"name"`
	Width           uint16                          `json:"width,omitempty"`
	Height          uint16                          `json:"height,omitempty"`
	MaxDuration     uint16                          `json:"maxDuration,omitempty"`
	Thumbnails      *HMSBrowserRecordingsThumbnails `json:"thumbnails,omitempty"`
	PresignDuration uint                            `json:"presignDuration,omitempty"`
	Role            string                          `json:"role"`
	AutoStart       bool                            `json:"autoStart,omitempty"`
	AutoStopTimeout uint                            `json:"autoStopTimeout,omitempty"`
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
}

type HLSRecordingLayers struct {
	Width        int `json:"width,omitempty"`
	Height       int `json:"height,omitempty"`
	VideoBitrate int `json:"videoBitrate,omitempty"`
	AudioBitrate int `json:"audioBitrate,omitempty"`
}

type HLSRecordingThumbnails struct {
}

type HLSRecording struct {
	HlsVod             bool                    `json:"hlsVod,omitempty"`
	SingleFilePerLayer bool                    `json:"singleFilePerLayer,omitempty"`
	EnableZipUpload    bool                    `json:"enableZipUpload,omitempty"`
	Layers             *HLSRecordingLayers     `json:"layers,omitempty"`
	Thumbnails         *HLSRecordingThumbnails `json:"thumbnails,omitempty"`
	PresignDuration    int                     `json:"presignDuration,omitempty"`
}

type HMSHlsDestinations struct {
	Name                    string                `json:"name"`
	MaxDuration             uint16                `json:"maxDuration,omitempty"`
	Layers                  *HLSDestinationLayers `json:"layers,omitempty"`
	PlaylistType            string                `json:"playlistType,omitempty"`
	MumPlaylistSegments     uint                  `json:"numPlaylistSegments,omitempty"`
	VideoFrameRate          uint16                `json:"videoFrameRate,omitempty"`
	EnableMetadataInsertion uint16                `json:"enableMetadataInsertion,omitempty"`
	EnableStaticUrl         uint16                `json:"enableStaticUrl,omitempty"`
	Recording               *HLSRecording         `json:"recording,omitempty"`
	AutoStopTimeout         uint                  `json:"autoStopTimeout,omitempty"`
}
type HMSTranscriptions struct {
	Name             string                             `json:"name"`
	Modes            []string                           `json:"modes,omitempty"`
	Role             string                             `json:"role"`
	OutputModes      []string                           `json:"outputModes,omitempty"`
	CustomVocabulary []string                           `json:"customVocabulary,omitempty"`
	Summary          *live_streams.TranscriptionSummary `json:"summary,omitempty"`
}

type HMSDestination struct {
	BrowserRecordings *HMSBrowserRecordings `json:"browserRecordings,omitempty"`
	RtmpDestinations  *HMSRtmpDestinations  `json:"rtmpDestinations,omitempty"`
	HlsDestinations   *HMSHlsDestinations   `json:"hlsDestinations,omitempty"`
	Transcriptions    *HMSTranscriptions    `json:"transcriptions,omitempty"`
}

type HMSTemplate struct {
	Name         string      `json:"name,omitempty"`
	Roles        *HMSRole    `json:"roles,omitempty"`
	Settings     *HMSSetting `json:"settings,omitempty"`
	Destinations *HMSRole    `json:"destinations,omitempty"`
}

type HMSTemplateQueryParam struct {
}

// Create a template
func CreateTemplate(ctx *gin.Context) {

	helpers.MakeApiRequest(ctx, policyBaseUrl, "POST", nil)
}

// Get a list of all rooms
// Applicable filters: name string, enabled *bool, after string, before string
func ListTemplates(ctx *gin.Context) {
	var param HMSTemplateQueryParam
	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("name", param.Name)
		if param.Enabled != nil {
			qs.Add("enabled", strconv.FormatBool(*param.Enabled))
		}
		qs.Add("before", param.Before)
		qs.Add("after", param.After)
	}
	helpers.MakeApiRequest(ctx, policyBaseUrl+"?"+qs.Encode(), "GET", nil)
}
