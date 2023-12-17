package main

import (
	"net/http"

	"api/active_room"
	"api/external_streams"
	"api/polls"
	"api/recording"
	"api/recording_assets"
	"api/room"
	"api/room_codes"
	"api/sessions"
	"api/token"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "100ms API works...",
	})

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Max-Age", "10000")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Authorization,Content-Type,Accept")
	}
}

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", ping)
	router.POST("/token", token.CreateToken)

	roomEndpoints := router.Group("/rooms")
	{

		roomEndpoints.GET("", room.ListRooms)
		roomEndpoints.GET("/:roomId", room.GetRoom)
		roomEndpoints.POST("", room.CreateRoom)
		roomEndpoints.POST("/:roomId", room.UpdateRoom)
		roomEndpoints.POST("/:roomId/enable", room.EnableRoom)
		roomEndpoints.POST("/:roomId/disable", room.DisableRoom)
	}

	roomCodesEndpoints := router.Group("/room-codes")
	{
		roomCodesEndpoints.GET("/:roomId", room_codes.GetRoomCode)
		roomCodesEndpoints.POST("/code/:code", room_codes.CreateShortCodeAuthToken)
		roomCodesEndpoints.POST("/:roomId", room_codes.CreateRoomCode)
		roomCodesEndpoints.POST("/:roomId/role/:role", room_codes.CreateRoomCodeForRole)
		roomCodesEndpoints.POST("/update", room_codes.UpdateRoomCode)

	}

	activeRoomsEndpoints := router.Group("/active-rooms")
	{
		activeRoomsEndpoints.GET("/:roomId", active_room.GetActiveRoom)
		activeRoomsEndpoints.GET("/:roomId/peers/:peerId", active_room.GetPeer)
		activeRoomsEndpoints.GET("/:roomId/peers", active_room.ListPeers)
		activeRoomsEndpoints.POST("/:roomId/peers/:peerId", active_room.UpdatePeer)
		activeRoomsEndpoints.POST("/:roomId/send-message", active_room.SendMessage)
		activeRoomsEndpoints.POST("/:roomId/remove-peers", active_room.RemovePeer)
		activeRoomsEndpoints.POST("/:roomId/end-room", active_room.EndRoom)
	}

	recordingsEndpoints := router.Group("/recordings")
	{
		recordingsEndpoints.POST("/room/:roomId/start", recording.StartRecording)
		recordingsEndpoints.POST("/room/:roomId/stop", recording.StopRecordings)
		recordingsEndpoints.POST("/:recordingId/stop", recording.StopRecording)
		recordingsEndpoints.GET("", recording.ListRecordings)
		recordingsEndpoints.GET("/:recordingId", recording.GetRecording)
		recordingsEndpoints.GET("/:recordingId/config", recording.GetRecordingConfig)
	}

	sessionsEndpoints := router.Group("/sessions")
	{
		sessionsEndpoints.GET("", sessions.ListSessions)
		sessionsEndpoints.GET("/:sessionId", sessions.GetSession)
	}

	recordingAssetsEndpoints := router.Group("/recording-assets")
	{
		recordingAssetsEndpoints.GET("", recording_assets.ListRecordingAssets)
		recordingAssetsEndpoints.GET("/:assetId", recording_assets.GetRecordingAsset)
		recordingAssetsEndpoints.GET("/:assetId/url", recording_assets.GetPresignedUrl)
	}

	externalStreamsEndpoints := router.Group("/external-streams")
	{
		externalStreamsEndpoints.POST("/room/:roomId/start", external_streams.StartExternalStream)
		externalStreamsEndpoints.POST("/room/:roomId/stop", external_streams.StopExternalStreams)
		externalStreamsEndpoints.POST("/:streamId/stop", external_streams.StopExternalStream)
		externalStreamsEndpoints.GET("", external_streams.ListExternalStreams)
		externalStreamsEndpoints.GET("/:streamId", external_streams.GetExternalStream)
	}

	pollsEndpoints := router.Group("/polls")
	{
		pollsEndpoints.GET("/:pollId", polls.GetPoll)
		pollsEndpoints.GET("/:pollId/sessions/:sessionId", polls.GetPollSessions)
		pollsEndpoints.GET("/:pollId/sessions/:sessionId/results", polls.ListPollResults)
		pollsEndpoints.GET("/:pollId/sessions/:sessionId/results/:resultId", polls.GetPollResult)
		pollsEndpoints.GET("/:pollId/sessions/:sessionId/responses", polls.ListPollResponses)
		pollsEndpoints.GET("/:pollId/sessions/:sessionId/responses/:responseId", polls.GetPollResponse)
		pollsEndpoints.POST("", polls.CreatePoll)
		pollsEndpoints.POST("/:pollId", polls.UpdatePoll)
		pollsEndpoints.POST("/:pollId/questions/:questionId", polls.UpdatePollQuestion)
		pollsEndpoints.POST("/:pollId/questions/:questionId/options/:optionId", polls.UpdatePollOption)
		pollsEndpoints.DELETE("/:pollId/questions/:questionId/options/:optionId", polls.DeletePollOption)
		pollsEndpoints.DELETE("/:pollId/questions/:questionId", polls.DeletePollQuestion)

	}

	router.Run()

}
