package room

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"io"
	"net/http"
	"os"

	"api/helpers"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Name *string `json:"name,omitempty"`
}

var ContentTypeHeader = map[string]string{"Content-Type": "application/json"}

// Create a   room with a given room name
func CreateRoom(ctx *gin.Context) {

	var rb RequestBody
	managementToken := helpers.GenerateManagementToken()

	if err := ctx.ShouldBind(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"name":   strings.ToLower(*rb.Name),
		"active": true,
	})

	fmt.Println("rb.Name: ", *rb.Name)
	payload := bytes.NewBuffer(postBody)
	// helpers.MakeApiRequest(ctx, "room", "POST", payload)

	baseUrl := os.Getenv("BASE_URL")
	method := "POST"
	roomUrl := baseUrl + "room"

	client := &http.Client{}
	req, err := http.NewRequest(method, roomUrl, payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// Add Authorization header
	req.Header.Add("Authorization", "Bearer "+managementToken)
	req.Header.Add("Content-Type", "application/json")

	// Send HTTP request
	res, err := client.Do(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	resp, err := io.ReadAll(res.Body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	defer res.Body.Close()

	ctx.Data(http.StatusOK, gin.MIMEJSON, resp)

}

// Get details of a given room
func GetRoomDetails(ctx *gin.Context) {

	roomId, ok := ctx.Params.Get("roomId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "roomId is missing"})
	}

	endpointPath := "room/" + roomId
	helpers.MakeApiRequest(ctx, endpointPath, "GET", nil)
}
