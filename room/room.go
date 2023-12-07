package room

import (
	"bytes"
	"encoding/json"

	"net/http"

	"api/helpers"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Name string `json:"name,omitempty"`
}

// Create a   room with a given room name
func CreateRoom(ctx *gin.Context) {

	var rb RequestBody

	if err := ctx.ShouldBind(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"name":    rb.Name,
		"enabled": true,
	})

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
