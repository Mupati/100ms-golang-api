package room

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type RequestBody struct {
	Room     *string `json:"room"`
	Duration *int    `json:"duration,omitempty"`
}

var ContentTypeHeader = map[string]string{"Content-Type": "application/json"}

func generateManagementToken(durationInHours int) string {
	appAccessKey := os.Getenv("APP_ACCESS_KEY")
	appSecret := os.Getenv("APP_SECRET")

	mySigningKey := []byte(appSecret)
	expiresIn := uint32(durationInHours * 3600)
	now := uint32(time.Now().UTC().Unix())
	exp := now + expiresIn
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": appAccessKey,
		"type":       "management",
		"version":    2,
		"jti":        uuid.New().String(),
		"iat":        now,
		"exp":        exp,
		"nbf":        now,
	})

	// Sign and get the complete encoded token as a string using the secret
	signedToken, _ := token.SignedString(mySigningKey)
	return signedToken
}

// Create a  call room with a given room name
func CreateRoom(ctx *gin.Context) {

	var managementToken string
	var rb RequestBody

	if err := ctx.ShouldBind(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if rb.Duration != nil {
		managementToken = generateManagementToken(*rb.Duration)
	} else {
		managementToken = generateManagementToken(24)
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"name":   strings.ToLower(*rb.Room),
		"active": true,
	})
	payload := bytes.NewBuffer(postBody)

	roomUrl := os.Getenv("ROOM_URL")
	method := "POST"

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
