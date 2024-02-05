package helpers

import (
	"api/hmserrors"
	"bytes"
	"time"

	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GetEnvironmentVariable(key string) (string, bool) {
	envValue, ok := os.LookupEnv(key)
	if ok {
		return envValue, ok
	}
	return "", ok
}

func GetEndpointUrl(path string) string {
	baseUrl, ok := GetEnvironmentVariable("BASE_URL")
	if !ok {
		panic(hmserrors.ErrMissingBaseUrl)
	}
	return baseUrl + path
}

func GenerateManagementToken() (string, error) {
	appAccessKey, ok := GetEnvironmentVariable("APP_ACCESS_KEY")

	if !ok {
		return "", hmserrors.ErrMissingAppAccessKey
	}
	appSecret, ok := GetEnvironmentVariable("APP_SECRET")
	if !ok {
		return "", hmserrors.ErrMissingAppSecretKey
	}

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
	return signedToken, nil
}

// Helper method to make all api calls to 100ms
func MakeApiRequest(ctx *gin.Context, url, method string, payload *bytes.Buffer) {

	var requestBody io.Reader
	managementToken, err := GenerateManagementToken()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

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

	ctx.Data(res.StatusCode, gin.MIMEJSON, resp)

}
