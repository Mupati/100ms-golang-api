package token

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type RequestBody struct {
	UserId *string `json:"userId"`
	RoomId *string `json:"roomId"`
	Role   *string `json:"role"`
}

func CreateToken(ctx *gin.Context) {

	var rb RequestBody

	if err := ctx.ShouldBind(&rb); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	appAccessKey := os.Getenv("APP_ACCESS_KEY")
	appSecret := os.Getenv("APP_SECRET")

	mySigningKey := []byte(appSecret)
	expiresIn := uint32(24 * 3600)
	now := uint32(time.Now().UTC().Unix())
	exp := now + expiresIn
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": appAccessKey,
		"type":       "app",
		"version":    2,
		"room_id":    *rb.RoomId,
		"user_id":    *rb.UserId,
		"role":       *rb.Role,
		"jti":        uuid.New().String(),
		"iat":        now,
		"exp":        exp,
		"nbf":        now,
	})

	// Sign and get the complete encoded token as a string using the secret
	signedToken, err := token.SignedString(mySigningKey)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": signedToken})
}
