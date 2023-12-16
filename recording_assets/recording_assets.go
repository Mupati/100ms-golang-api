package recording_assets

import (
	"api/helpers"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var recordingAssetsBaseUrl = os.Getenv("BASE_URL") + "recording-assets"

type HMSRecordingAssetsQueryParam struct {
	RoomId    string `form:"room_id,omitempty"`
	SessionId string `form:"session_id,omitempty"`
	Status    string `form:"status,omitempty"`
	Start     string `form:"start,omitempty"`
	Limit     int32  `form:"limit,omitempty"`
}

// Get asset id
func GetRecordingAsset(ctx *gin.Context) {
	assetId, ok := ctx.Params.Get("assetId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "provide the asset ID"})
	}
	helpers.MakeApiRequest(ctx, recordingAssetsBaseUrl+"/"+assetId, "GET", nil)
}

// List all recording assets
// Applicable filters: room_id string, session_id string, status string, start string, limit int32
func ListRecordingAssets(ctx *gin.Context) {

	var param HMSRecordingAssetsQueryParam

	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("room_id", param.RoomId)
		qs.Add("session_id", param.SessionId)
		qs.Add("status", param.Status)
		qs.Add("start", param.Start)
		qs.Add("limit", strconv.Itoa(int(param.Limit)))
	}

	helpers.MakeApiRequest(ctx, recordingAssetsBaseUrl+"?"+qs.Encode(), "GET", nil)

}

// Get the Presigned url
func GetPresignedUrl(ctx *gin.Context) {
	assetId, ok := ctx.Params.Get("assetId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "provide the asset ID"})
	}

	presignDuration := ctx.Query("presign_duration")
	qs := url.Values{}
	qs.Add("presign_duration", presignDuration)

	helpers.MakeApiRequest(ctx, recordingAssetsBaseUrl+"/"+assetId+"/presigned-url?"+qs.Encode(), "GET", nil)
}
