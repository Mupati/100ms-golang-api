package activeroom

import (
	"api/helpers"
	"api/hmserrors"
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

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

type HMSActiveRoomQueryParam struct {
	UserId string `form:"user_id,omitempty"`
	Role   string `form:"role,omitempty"`
}

var activeRoomBaseUrl = helpers.GetEndpointUrl("active-rooms")

// Get active room details
func GetActiveRoom(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
	}

	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId, "GET", nil)
}

// Fetch a single peer's details
func GetPeer(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	peerId, ok1 := ctx.Params.Get("peerId")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomIdAndPeerId})
	}
	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId+"/peers/"+peerId, "GET", nil)
}

// List all peers details
// Filters: user_id and role
func ListPeers(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
	}

	var param HMSActiveRoomQueryParam
	qs := url.Values{}
	if ctx.BindQuery(&param) == nil {
		qs.Add("user_id", param.UserId)
		qs.Add("role", param.Role)
	}

	helpers.MakeApiRequest(ctx, activeRoomBaseUrl+"/"+roomId+"/peers"+"?"+qs.Encode(), "GET", nil)
}

// Update a single peer's details
func UpdatePeer(ctx *gin.Context) {
	roomId, ok := ctx.Params.Get("roomId")
	peerId, ok1 := ctx.Params.Get("peerId")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomIdAndPeerId})
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
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
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
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
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
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": hmserrors.ErrMissingRoomId})
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
