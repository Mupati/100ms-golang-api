package room_codes

import (
	"api/helpers"
	"api/hms_errors"
	"bytes"
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

type HMSRoomCodeUpdateRequestBody struct {
	Code    string `json:"code"`
	Enabled bool   `json:"enabled"`
}

var roomCodeBaseUrl = helpers.GetEndpointUrl("room-codes")
var authBaseUrl, _ = helpers.GetEnvironmentVariable("AUTH_BASE_URL")

// Get room codes
func GetRoomCode(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingRoomId})
	}

	helpers.MakeApiRequest(ctx, roomCodeBaseUrl+"/room/"+roomId, "GET", nil)
}

// Create room code for all roles
func CreateRoomCode(ctx *gin.Context) {

	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingRoomId})
	}

	helpers.MakeApiRequest(ctx, roomCodeBaseUrl+"/room/"+roomId, "POST", nil)
}

// Create room code for a given role
func CreateRoomCodeForRole(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	role, ok1 := ctx.Params.Get("role")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingRoomIdAndRole})
	}
	helpers.MakeApiRequest(ctx, roomCodeBaseUrl+"/room/"+roomId+"/role/"+role, "POST", nil)
}

// Create room code for a given role
func UpdateRoomCode(ctx *gin.Context) {
	var rb HMSRoomCodeUpdateRequestBody
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSRoomCodeUpdateRequestBody{
		Code:    rb.Code,
		Enabled: rb.Enabled,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, roomCodeBaseUrl+"/code", "POST", payload)
}

func CreateShortCodeAuthToken(ctx *gin.Context) {
	code, ok := ctx.Params.Get("code")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hms_errors.ErrMissingAuthCode})
	}

	postBody, _ := json.Marshal(map[string]string{
		"code": code,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, authBaseUrl+"token", "POST", payload)
}
