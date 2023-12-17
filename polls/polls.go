package polls

import (
	"api/helpers"
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type PollOption struct {
	Index  int    `json:"index,omitempty"`
	Text   string `json:"text,omitempty"`
	Weight int    `json:"weight,omitempty"`
}

type PollAnswer struct {
	Hidden  bool          `json:"hidden,omitempty"`
	Options *[]PollOption `json:"options,omitempty"`
	Text    string        `json:"text,omitempty"`
	Case    bool          `json:"case,omitempty"`
	Trim    string        `json:"trim,omitempty"`
}

type PollQuestion struct {
	Index        int           `json:"index,omitempty"`
	Text         string        `json:"text,omitempty"`
	Format       string        `json:"format,omitempty"`
	Attachment   []string      `json:"attachment,omitempty"`
	Skippable    bool          `json:"skippable,omitempty"`
	Duration     int           `json:"duration,omitempty"`
	Once         bool          `json:"once,omitempty"`
	Weight       int           `json:"weight,omitempty"`
	AnswerMinLen bool          `json:"answer_min_len,omitempty"`
	AnswerMaxLen bool          `json:"answer_max_len,omitempty"`
	Answer       *PollAnswer   `json:"answer,omitempty"`
	Options      *[]PollOption `json:"options,omitempty"`
}

type HMSPoll struct {
	Title     string          `json:"title,omitempty"`
	Duration  int             `json:"duration,omitempty"`
	Anonymous bool            `json:"anonymous,omitempty"`
	Mode      string          `json:"mode,omitempty"`
	Type      string          `json:"type,omitempty"`
	Start     string          `json:"start,omitempty"`
	Questions *[]PollQuestion `json:"questions,omitempty"`
}

const MISSING_POLL_ID_ERROR_MESSAGE = "provide a poll ID"

var pollBaseUrl = os.Getenv("BASE_URL") + "polls"

// Create a poll
func CreatePoll(ctx *gin.Context) {

	var rb HMSPoll
	postBody, _ := json.Marshal(HMSPoll{
		Title:     rb.Title,
		Duration:  rb.Duration,
		Anonymous: rb.Anonymous,
		Mode:      rb.Mode,
		Type:      rb.Type,
		Start:     rb.Start,
		Questions: rb.Questions,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, pollBaseUrl, "POST", payload)
}

// Get a Poll
func GetPoll(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE})
	}
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId, "GET", nil)
}

// Update a poll
func UpdatePoll(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE})
	}

	var rb HMSPoll
	postBody, _ := json.Marshal(HMSPoll{
		Title:     rb.Title,
		Duration:  rb.Duration,
		Anonymous: rb.Anonymous,
		Mode:      rb.Mode,
		Type:      rb.Type,
		Start:     rb.Start,
		Questions: rb.Questions, // TODO: look into this object and instances of nil
	})
	payload := bytes.NewBuffer(postBody)

	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId, "POST", payload)
}

// Update a poll question
func UpdatePollQuestion(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	questionId, ok1 := ctx.Params.Get("questionId")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE + " and question ID"})
	}

	var rb PollQuestion
	postBody, _ := json.Marshal(PollQuestion{
		Index:        rb.Index,
		Text:         rb.Text,
		Format:       rb.Format,
		Attachment:   rb.Attachment,
		Skippable:    rb.Skippable,
		Duration:     rb.Duration,
		Once:         rb.Once,
		Weight:       rb.Weight,
		AnswerMinLen: rb.AnswerMinLen,
		AnswerMaxLen: rb.AnswerMaxLen,
		Answer:       rb.Answer,  // TODO: look into this object and instances of nil
		Options:      rb.Options, // TODO: look into this object and instances of nil
	})

	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId+"/questions/"+questionId, "POST", payload)
}

// Delete a poll question
func DeletePollQuestion(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	questionId, ok1 := ctx.Params.Get("questionId")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE + " and question ID"})
	}
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId+"/questions/"+questionId, "DELETE", nil)
}

// Update a poll option
func UpdatePollOption(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	questionId, ok1 := ctx.Params.Get("questionId")
	optionId, ok2 := ctx.Params.Get("optionId")
	if !ok || !ok1 || !ok2 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE + " ,question ID and option ID"})
	}

	var rb PollOption
	postBody, _ := json.Marshal(PollOption{
		Index:  rb.Index,
		Text:   rb.Text,
		Weight: rb.Weight,
	})
	payload := bytes.NewBuffer(postBody)
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId+"/questions/"+questionId+"/options/"+optionId, "POST", payload)
}

// Delete a poll option
func DeletePollOption(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	questionId, ok1 := ctx.Params.Get("questionId")
	optionId, ok2 := ctx.Params.Get("optionId")
	if !ok || !ok1 || !ok2 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE + " ,question ID and option ID"})
	}
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId+"/questions/"+questionId+"/options/"+optionId, "DELETE", nil)
}

// Get a poll session
func GetPollSessions(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	sessionId, ok1 := ctx.Params.Get("sessionId")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE + " and session ID"})
	}
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId+"/sessions/"+sessionId, "GET", nil)
}

// Get a poll result
func GetPollResult(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	sessionId, ok1 := ctx.Params.Get("sessionId")
	resultId, ok2 := ctx.Params.Get("resultId")
	if !ok || !ok1 || !ok2 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE + " ,session ID and results ID"})
	}
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId+"/sessions/"+sessionId+"/results/"+resultId, "GET", nil)
}

// List  poll results
func ListPollResults(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	sessionId, ok1 := ctx.Params.Get("sessionId")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE + " and session ID"})
	}
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId+"/sessions/"+sessionId+"/results", "GET", nil)
}

// List  poll responses
func ListPollResponses(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	sessionId, ok1 := ctx.Params.Get("sessionId")
	if !ok || !ok1 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE + " and session ID"})
	}
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId+"/sessions/"+sessionId+"/responses", "GET", nil)
}

// Get a poll response
func GetPollResponse(ctx *gin.Context) {
	pollId, ok := ctx.Params.Get("pollId")
	sessionId, ok1 := ctx.Params.Get("sessionId")
	responseId, ok2 := ctx.Params.Get("resultId")
	if !ok || !ok1 || !ok2 {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": MISSING_POLL_ID_ERROR_MESSAGE + " ,session ID and results ID"})
	}
	helpers.MakeApiRequest(ctx, pollBaseUrl+"/"+pollId+"/sessions/"+sessionId+"/responses/"+responseId, "GET", nil)
}
