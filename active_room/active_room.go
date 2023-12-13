package active_room

import (
	"api/helpers"
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type HMSPeerUpdateBody struct {
	Name     string                 `json:"name,omitempty"`
	Role     string                 `json:"role,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type HMSMessageBody struct {
	PeerId  string `json:"peer_id,omitempty"`
	Role    string `json:"role,omitempty"`
	Message string `json:"message"`
	Type    string `json:"type,omitempty"`
}

type HMSRemovePeerBody struct {
	PeerId string `json:"peer_id,omitempty"`
	Role   string `json:"role,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type HMSEndRoomBody struct {
	Lock   bool   `json:"lock,omitempty"`
	Reason string `json:"reason,omitempty"`
}

var activeRoomBaseUrl = os.Getenv("BASE_URL") + "active-rooms"

const MISSING_ROOM_ID_ERROR_MESSAGE = "provide a room ID"

// Get active room details
func GetActiveRoom(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}

	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId, "GET", nil)
}

// Fetch a single peer's details
func GetPeer(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	peerId, ok1 := ctx.Params.Get("peerId")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE + " and peer ID"})
	}
	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId+"/peers/"+peerId, "GET", nil)
}

// List all peers details
// TODO: add filters user_id and role
func ListPeers(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}
	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId+"/peers", "GET", nil)
}

// Update a single peer's details
func UpdatePeer(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	peerId, ok1 := ctx.Params.Get("peerId")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE + " and peer ID"})
	}

	var rb HMSPeerUpdateBody

	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSPeerUpdateBody{
		Name:     rb.Name,
		Role:     rb.Role,
		Metadata: rb.Metadata,
	})

	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId+"/peers/"+peerId, "POST", payload)
}

// Send a message
func SendMessage(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}

	var rb HMSMessageBody
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSMessageBody{
		PeerId:  rb.PeerId,
		Role:    rb.Role,
		Message: rb.Message,
		Type:    rb.Type,
	})

	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId+"/send-message", "POST", payload)
}

// Remove a peer
func RemovePeer(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}

	var rb HMSRemovePeerBody
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSRemovePeerBody{
		PeerId: rb.PeerId,
		Role:   rb.Role,
		Reason: rb.Reason,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId+"/remove-peers", "POST", payload)
}

// Remove a peer
func EndRoom(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_ROOM_ID_ERROR_MESSAGE})
	}

	var rb HMSEndRoomBody
	if err := ctx.ShouldBindJSON(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(HMSEndRoomBody{
		Lock:   rb.Lock,
		Reason: rb.Reason,
	})
	payload := bytes.NewBuffer(postBody)

	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId+"/end-room", "POST", payload)
}