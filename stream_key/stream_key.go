package stream_key

import (
	"api/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var streamKeysBaseUrl = os.Getenv("BASE_URL") + "stream-keys"

const missingRoomIdErrorMessage = "provide a room ID"

// Get stream key
func GetStreamKey(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": missingRoomIdErrorMessage})
	}
	helpers.MakeApiRequest(ctx, streamKeysBaseUrl+"/"+roomId, "GET", nil)
}

// Create stream key
func CreateStreamKey(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": missingRoomIdErrorMessage})
	}
	helpers.MakeApiRequest(ctx, streamKeysBaseUrl+"/"+roomId, "POST", nil)
}

// Disable stream key
func DisableStreamKey(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": missingRoomIdErrorMessage})
	}
	helpers.MakeApiRequest(ctx, streamKeysBaseUrl+"/"+roomId+"/disable", "POST", nil)
}
