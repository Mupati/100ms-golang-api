package streamkey

import (
	"api/helpers"
	"api/hmserrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var streamKeysBaseUrl = helpers.GetEndpointUrl("stream-keys")

// Get stream key
func GetStreamKey(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
	}
	helpers.MakeApiRequest(ctx, streamKeysBaseUrl+"/"+roomId, "GET", nil)
}

// Create stream key
func CreateStreamKey(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
	}
	helpers.MakeApiRequest(ctx, streamKeysBaseUrl+"/"+roomId, "POST", nil)
}

// Disable stream key
func DisableStreamKey(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
	}
	helpers.MakeApiRequest(ctx, streamKeysBaseUrl+"/"+roomId+"/disable", "POST", nil)
}
