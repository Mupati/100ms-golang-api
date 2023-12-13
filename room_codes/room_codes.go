package room_codes

import (
	"api/helpers"
	"bytes"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

const MISSING_ROOM_ID_ERROR_MESSAGE = "provide a room ID"

type HMSRoomCodeUpdateRequestBody struct {
	Code    string `json:"code"`
	Enabled bool   `json:"enabled"`
}

var roomCodeBaseUrl = os.Getenv("BASE_URL") + "room-codes"
var authBaseUrl = os.Getenv("AUTH_BASE_URL")

// Get room codes
func GetRoomCode(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}

	helpers.MakeApiRequest(ctx, roomCodeBaseUrl+"/room/"+roomId, "GET", nil)
}

// Create room code for all roles
func CreateRoomCode(ctx *gin.Context) {

	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}

	helpers.MakeApiRequest(ctx, roomCodeBaseUrl+"/room/"+roomId, "POST", nil)
}

// Create room code for a given role
func CreateRoomCodeForRole(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	role, ok1 := ctx.Params.Get("role")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE + " and role"})
	}
	helpers.MakeApiRequest(ctx, roomCodeBaseUrl+"/room/"+roomId+"/role/"+role, "POST", nil)
}

// Create room code for a given role
func UpdateRoomCode(ctx *gin.Context) {
	var rb HMSRoomCodeUpdateRequestBody
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	requestBody := HMSRoomCodeUpdateRequestBody{
		Code:    rb.Code,
		Enabled: rb.Enabled,
	}

	postBody, _ := json.Marshal(requestBody)
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, roomCodeBaseUrl+"/code", "POST", payload)
}

func CreateShortCodeAuthToken(ctx *gin.Context) {
	code, ok := ctx.Params.Get("code")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "provide the code"})
	}

	postBody, _ := json.Marshal(map[string]string{
		"code": code,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, authBaseUrl+"token", "POST", payload)
}
