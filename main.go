package main

import (
	"net/http"

	"api/activeroom"
	"api/analytics"
	externalstreams "api/externalstreams"
	"api/livestreams"
	"api/policy"
	"api/polls"
	"api/recording"
	"api/recordingassets"
	"api/room"
	"api/roomcodes"
	"api/sessions"
	"api/streamkey"
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
		roomCodesEndpoints.GET("/:roomId", roomcodes.GetRoomCode)
		roomCodesEndpoints.POST("/code/:code", roomcodes.CreateShortCodeAuthToken)
		roomCodesEndpoints.POST("/:roomId", roomcodes.CreateRoomCode)
		roomCodesEndpoints.POST("/:roomId/role/:role", roomcodes.CreateRoomCodeForRole)
		roomCodesEndpoints.POST("/update", roomcodes.UpdateRoomCode)

	}

	activeRoomsEndpoints := router.Group("/active-rooms")
	{
		activeRoomsEndpoints.GET("/:roomId", activeroom.GetActiveRoom)
		activeRoomsEndpoints.GET("/:roomId/peers/:peerId", activeroom.GetPeer)
		activeRoomsEndpoints.GET("/:roomId/peers", activeroom.ListPeers)
		activeRoomsEndpoints.POST("/:roomId/peers/:peerId", activeroom.UpdatePeer)
		activeRoomsEndpoints.POST("/:roomId/send-message", activeroom.SendMessage)
		activeRoomsEndpoints.POST("/:roomId/remove-peers", activeroom.RemovePeer)
		activeRoomsEndpoints.POST("/:roomId/end-room", activeroom.EndRoom)
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
		recordingAssetsEndpoints.GET("", recordingassets.ListRecordingAssets)
		recordingAssetsEndpoints.GET("/:assetId", recordingassets.GetRecordingAsset)
		recordingAssetsEndpoints.GET("/:assetId/url", recordingassets.GetPresignedUrl)
	}

	externalStreamsEndpoints := router.Group("/external-streams")
	{
		externalStreamsEndpoints.POST("/room/:roomId/start", externalstreams.StartExternalStream)
		externalStreamsEndpoints.POST("/room/:roomId/stop", externalstreams.StopExternalStreams)
		externalStreamsEndpoints.POST("/:streamId/stop", externalstreams.StopExternalStream)
		externalStreamsEndpoints.GET("", externalstreams.ListExternalStreams)
		externalStreamsEndpoints.GET("/:streamId", externalstreams.GetExternalStream)
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

	liveStreamsEndpoints := router.Group("/live-streams")
	{
		liveStreamsEndpoints.POST("/room/:roomId/start", livestreams.StartLiveStream)
		liveStreamsEndpoints.POST("/room/:roomId/stop", livestreams.StopLiveStreams)
		liveStreamsEndpoints.POST("/:streamId/stop", livestreams.StopLiveStream)
		liveStreamsEndpoints.POST("/:streamId/metadata", livestreams.SendTimedMetada)
		liveStreamsEndpoints.POST("/:streamId/pause-recording", livestreams.PauseLiveStreamRecording)
		liveStreamsEndpoints.POST("/:streamId/resume-recording", livestreams.ResumeLiveStreamRecording)
		liveStreamsEndpoints.GET("", livestreams.ListLiveStreams)
		liveStreamsEndpoints.GET("/:streamId", livestreams.GetLiveStream)
	}

	policyEndpoints := router.Group("/templates")
	{

		policyEndpoints.GET("", policy.ListTemplates)
		policyEndpoints.GET("/:templateId", policy.GetTemplate)
		policyEndpoints.GET("/:templateId/roles/:roleName", policy.GetTemplateRole)
		policyEndpoints.GET("/:templateId/settings", policy.GetTemplateSettings)
		policyEndpoints.GET("/:templateId/destinations", policy.GetTemplateDestinations)

		policyEndpoints.POST("", policy.CreateTemplate)
		policyEndpoints.POST("/:templateId", policy.UpdateTemplate)
		policyEndpoints.POST("/:templateId/roles/:roleName", policy.ModifyTemplateRole)
		policyEndpoints.POST("/:templateId/settings", policy.UpdateTemplateSettings)
		policyEndpoints.POST("/:templateId/destinations", policy.UpdateTemplateDestinations)

		policyEndpoints.DELETE("/:templateId/roles/:roleName", policy.DeleteTemplateRole)
	}

	streamKeyEndpoints := router.Group("/stream-keys")
	{
		streamKeyEndpoints.GET("/:roomId", streamkey.GetStreamKey)
		streamKeyEndpoints.POST("/:roomId", streamkey.CreateStreamKey)
		streamKeyEndpoints.POST("/:roomId/disable", streamkey.DisableStreamKey)
	}

	// Analytics Events
	router.GET("/analytics", analytics.GetAnalyticsEvents)

	router.Run()

}
