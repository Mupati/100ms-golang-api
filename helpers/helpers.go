package helpers

import (
	"bytes"
	"time"

	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateManagementToken() string {
	appAccessKey := os.Getenv("APP_ACCESS_KEY")
	appSecret := os.Getenv("APP_SECRET")

	mySigningKey := []byte(appSecret)
	expiresIn := uint32(24 * 3600)
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

// Helper method to make all api calls to 100ms
func MakeApiRequest(ctx *gin.Context, url, method string, payload *bytes.Buffer) {

	var requestBody io.Reader

	managementToken := GenerateManagementToken()

	client := &http.Client{}

	if payload == nil {
		requestBody = nil
	} else {
		requestBody = payload
	}

	req, err := http.NewRequest(method, url, requestBody)
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

	// TODO
	// Get the actual response to send the correct status code.
	// Currently when there's an error, we still send 200, but the error code is in the payload.
	ctx.Data(http.StatusOK, gin.MIMEJSON, resp)

}
